package searchAlgorithms

import (
	"fmt"
	"github.com/Krud3/InteligenciaArtificial/src/datatypes"
	"math"
	"sort"
	"time"
)

// enviroment represents the enviroment where the agent is going to move
type enviroment struct {
	agent             agent
	passengerPosition datatypes.BoardCoordinate
	board             [][]int
	totalGoal         int
}

// SearchAlgorithm is the interface that the search algorithms must implement
type SearchAgorithm interface {
	LookForGoal(*enviroment) SearchResult
}

// agent represents the agent that is going to move in the enviroment
type agent struct {
	position          datatypes.AgentStep
	passenger         bool
	searchAlgorithm   SearchAgorithm
	ambientPerception [4]int
}

var contiguousMovements = [4]datatypes.CoordinateMovement{
	datatypes.CoordinateMovement{
		Direction:  datatypes.UP,
		Coordinate: datatypes.BoardCoordinate{X: 0, Y: -1}},
	datatypes.CoordinateMovement{
		Direction:  datatypes.RIGHT,
		Coordinate: datatypes.BoardCoordinate{X: 1, Y: 0}},
	datatypes.CoordinateMovement{
		Direction:  datatypes.DOWN,
		Coordinate: datatypes.BoardCoordinate{X: 0, Y: 1}},
	datatypes.CoordinateMovement{
		Direction:  datatypes.LEFT,
		Coordinate: datatypes.BoardCoordinate{X: -1, Y: 0}},
}

func coordinateAdd(firstCoordinate datatypes.BoardCoordinate, secondCoordinates datatypes.BoardCoordinate) datatypes.BoardCoordinate {
	return datatypes.BoardCoordinate{X: firstCoordinate.X + secondCoordinates.X, Y: firstCoordinate.Y + secondCoordinates.Y}
}

func Percept(a agent, board [][]int) []datatypes.CoordinateMovement {
	canMove := []datatypes.CoordinateMovement{}
	for _, contiguous := range contiguousMovements {
		testingCoordinate := coordinateAdd(a.position.CurrentPosition, contiguous.Coordinate)
		if testingCoordinate.X >= 0 &&
			testingCoordinate.Y >= 0 &&
			testingCoordinate.X < len(board) &&
			testingCoordinate.Y < len(board) &&
			board[testingCoordinate.X][testingCoordinate.Y] != 1 {
			canMove = append(canMove, datatypes.CoordinateMovement{
				Direction:  contiguous.Direction,
				Coordinate: testingCoordinate,
			})
		}
	}
	fmt.Println("Can move: %s", canMove)
	return canMove
}

type BreadthFirstSearch struct{}

type UniformCostSearch struct{}

type DepthSearch struct{}

type SearchResult struct {
	pathFound                []datatypes.BoardCoordinate
	solutionFound            bool
	expandenNodes, treeDepth int
	cost                     float32
	timeExe                  time.Duration
}

func removeDuplicates(slice []datatypes.BoardCoordinate) []datatypes.BoardCoordinate {
	// Create a map to track seen items
	seen := make(map[datatypes.BoardCoordinate]bool)
	result := []datatypes.BoardCoordinate{}

	for _, item := range slice {
		// If the item has not been seen, add it to the result
		if !seen[item] {
			seen[item] = true
			result = append(result, item)
		}
	}

	return result
}

func (a *BreadthFirstSearch) LookForGoal(e *enviroment) SearchResult {
	parentNodes := make(datatypes.Set) // For Hold the parents node removed from the queue
	expandenNodes := 0
	treeDepth := 0
	initialPosition := e.agent.position
	queue := datatypes.Queue[datatypes.AgentStep]{}
	queue.Enqueue(initialPosition)
	start := time.Now()
	for !queue.IsEmpty() {
		currentStep, empty := queue.Dequeue()
		if empty {
			return SearchResult{
				[]datatypes.BoardCoordinate{},
				false,
				expandenNodes,
				treeDepth,
				0.5,
				time.Now().Sub(start),
			}
		}
		parentNodes.Add(currentStep)
		fmt.Printf("current: %s", currentStep)

		if e.board[currentStep.CurrentPosition.X][currentStep.CurrentPosition.Y] == 5 {
			e.passengerPosition = datatypes.BoardCoordinate{
				X: currentStep.CurrentPosition.X,
				Y: currentStep.CurrentPosition.Y,
			}
			e.agent = agent{
				initialPosition,
				true,
				e.agent.searchAlgorithm,
				e.agent.ambientPerception,
			}
			parentNodes = make(datatypes.Set)
			parentNodes.Add(datatypes.AgentStep{
				Action:          datatypes.UP,
				Depth:           0,
				CurrentPosition: currentStep.CurrentPosition,
				PreviousPosition: datatypes.BoardCoordinate{
					X: math.MaxInt,
					Y: math.MaxInt,
				},
			})
			// queue.Clear()
			// queue.Enqueue(datatypes.AgentStep{
			// 	Action:          datatypes.UP,
			// 	Depth:           0,
			// 	CurrentPosition: currentStep.CurrentPosition,
			// 	PreviousPosition: datatypes.BoardCoordinate{
			// 		X: math.MaxInt,
			// 		Y: math.MaxInt,
			// 	},
			// })

		}

		if false && e.agent.passenger && e.board[currentStep.CurrentPosition.X][currentStep.CurrentPosition.Y] == 6 {
			end := time.Now()
			traveledPath := []datatypes.BoardCoordinate{}
			traveledStep := currentStep.CurrentPosition
			fmt.Println("Passenger ", e.passengerPosition)

			for traveledStep.X != e.passengerPosition.X && e.passengerPosition.Y != traveledStep.Y { //
				fmt.Println("Traveled path: ", traveledPath)
				fmt.Println("Traveled step: ", traveledStep)
				fmt.Println("passenger ", e.passengerPosition)
				fmt.Println("Parents ", parentNodes)
				passes := []datatypes.AgentStep{}
				for key, _ := range parentNodes {
					//fmt.Println("Element: ", key)
					if step, ok := key.(datatypes.AgentStep); ok {
						if step.CurrentPosition.X == traveledStep.X && step.CurrentPosition.Y == traveledStep.Y {
							passes = append(passes, datatypes.AgentStep{
								Action:           step.Action,
								Depth:            step.Depth,
								CurrentPosition:  step.CurrentPosition,
								PreviousPosition: step.PreviousPosition,
							})
						}

						// if step.CurrentPosition.X == step.PreviousPosition.X && step.CurrentPosition.Y == step.PreviousPosition.Y {
						// 	traveledStep = traveledPath[len(traveledPath)-1]
						// }
						// if traveledStep.X == e.passengerPosition.X && traveledStep.Y == e.passengerPosition.Y {
						// 	// Set the final variable before exiting the loop
						// 	traveledPath = append(traveledPath, datatypes.BoardCoordinate{
						// 		X: step.PreviousPosition.X,
						// 		Y: step.PreviousPosition.Y,
						// 	})
						// }
					} else { //
						fmt.Println("element is not of type AgentStep")
					}
				}
				// Sort by Direction Weight
				sort.Sort(datatypes.ByAction(passes))
				fmt.Println("Same coordinate", passes)
				if len(passes) > 0 {
					passed := passes[0]
					priory := parentNodes.GetAndRemove(passed)
					fmt.Println("Passed ", priory)
					traveledPath = append(traveledPath, datatypes.BoardCoordinate{
						X: passed.CurrentPosition.X,
						Y: passed.CurrentPosition.Y,
					})
					fmt.Println("Traveled", traveledPath)
					time.Sleep(1000 * time.Millisecond)
					traveledStep = passed.PreviousPosition
				}
			}

			return SearchResult{
				traveledPath,
				true,
				expandenNodes,
				treeDepth,
				1.333,
				end.Sub(start),
			}
		}
		fmt.Println("\nPosition: %s", currentStep)
		e.agent.position = currentStep
		agentPerception := Percept(e.agent, e.board)
		expandenNodes++
		for _, perception := range agentPerception {
			fmt.Println("perception ", perception)
			fmt.Println("step ", currentStep.PreviousPosition)
			fmt.Println("parent ", parentNodes)
			fmt.Println("Queeu ", queue)
			if (perception.Coordinate.X != currentStep.PreviousPosition.X) || (perception.Coordinate.Y != currentStep.PreviousPosition.Y) {
				fmt.Println("entra")
				queue.Enqueue(
					datatypes.AgentStep{
						Action:           perception.Direction,
						Depth:            currentStep.Depth + 1,
						CurrentPosition:  perception.Coordinate,
						PreviousPosition: currentStep.CurrentPosition,
					},
				)
			}
		}
	}
	fmt.Println("No solution found")
	return SearchResult{}
}

func (a *UniformCostSearch) LookForGoal(e *enviroment) SearchResult {
	parentNodes := make(datatypes.Set) // For Hold the parents node removed from the queue
	expandenNodes := 0
	treeDepth := 0
	initialPosition := e.agent.position
	priorityQueue := datatypes.PriorityQueue[datatypes.AgentStep]{}
	priorityQueue.Push(datatypes.Element[datatypes.AgentStep]{
		Value:    initialPosition,
		Priority: 1},
	)
	for !priorityQueue.IsEmpty() {
		currentStep, empty := priorityQueue.Pop()
		if empty {
			return SearchResult{
				[]datatypes.BoardCoordinate{},
				false,
				expandenNodes,
				treeDepth,
				0.0,
				0,
			}
		}
		parentNodes.Add(currentStep)
		e.agent.position = currentStep
		agentPerception := Percept(e.agent, e.board)
		expandenNodes++
		for _, perception := range agentPerception {
			if !parentNodes.Contains(perception) {
				priorityQueue.Push(
					datatypes.Element[datatypes.AgentStep]{
						Value:    datatypes.AgentStep{CurrentPosition: perception.Coordinate, PreviousPosition: currentStep.CurrentPosition},
						Priority: 2,
					},
				)
			}
		}
		return SearchResult{}
	}
	return SearchResult{}
}

func (a *DepthSearch) LookForGoal(e *enviroment) SearchResult {
	return SearchResult{}
}

func StartGame(strategy int) {
	scannedMatrix, error := GetMatrix()
	if error != nil {
		fmt.Println("Error to load the matrix")
	}

	var searchStrategy SearchAgorithm

	switch strategy {
	case 1:
		searchStrategy = &BreadthFirstSearch{}
	case 2:
		searchStrategy = &UniformCostSearch{}
	case 4:
		searchStrategy = &BreadthFirstSearch{}
	default:
		fmt.Println("Unknown strategy")
	}

	initialPosition, exists := scannedMatrix.MainCoordinates["init"]
	if exists {
		initialStep := datatypes.AgentStep{
			Action:          math.MaxInt,
			Depth:           0,
			CurrentPosition: initialPosition,
			PreviousPosition: datatypes.BoardCoordinate{
				X: math.MaxInt,
				Y: math.MaxInt,
			},
		}
		agent := agent{
			initialStep,
			false,
			searchStrategy,
			[4]int{},
		}
		result := agent.searchAlgorithm.LookForGoal(&enviroment{
			agent,
			datatypes.BoardCoordinate{X: math.MaxInt, Y: math.MaxInt},
			scannedMatrix.Matrix,
			0},
		)
		fmt.Println(result)
	} else {

	}

}
