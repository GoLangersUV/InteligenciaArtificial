package searchAlgorithms

import (
	"fmt"
	"github.com/Krud3/InteligenciaArtificial/src/datatypes"
	"math"
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

type Predicate[T any] func(T) bool

// Generic Find function that takes a slice of any type and a predicate
func Find[T any](slice []T, predicate Predicate[T]) (T, bool) {
	for _, value := range slice {
		if predicate(value) {
			return value, true // Return the index if the predicate is true
		}
	}
	var zeroValue T
	return zeroValue, false // Return -1 if no element satisfies the predicate
}

func (a *BreadthFirstSearch) LookForGoal(e *enviroment) SearchResult {
	// For build the complete path at the end.
	var paths [][]datatypes.AgentStep
	// Keep the parent nodes saved then these were removed from the queue
	parentNodes := []datatypes.AgentStep{}
	// Keep the expanded nodes
	expandenNodes := 0
	// Get the initial position from the pass in environment
	initialPosition := e.agent.position
	// Data structure that implement the FIFO queue
	queue := datatypes.Queue[datatypes.AgentStep]{}
	queue.Enqueue(initialPosition)

	// Start the computation time
	start := time.Now()
	// Check if there are more children vto process
	for !queue.IsEmpty() {

		currentStep, empty := queue.Dequeue()
		if empty {
			return SearchResult{
				[]datatypes.BoardCoordinate{},
				false,
				expandenNodes,
				0,
				0,
				time.Now().Sub(start),
			}
		}
		parentNodes = append(parentNodes, currentStep)
		fmt.Println("***************************************************")
		fmt.Println("")
		fmt.Printf("current: \n", currentStep)
		fmt.Println("")
		fmt.Println("***************************************************")
		fmt.Println("")
		time.Sleep(1000 * time.Millisecond)

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
			paths = append(paths, parentNodes)
			parentNodes = append(parentNodes, datatypes.AgentStep{
				Action:          currentStep.Action,
				Depth:           currentStep.Depth + 1,
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

		if e.agent.passenger && e.board[currentStep.CurrentPosition.X][currentStep.CurrentPosition.Y] == 6 {
			end := time.Now()
			traveledPath := []datatypes.BoardCoordinate{}
			traveledStep := currentStep

			// Start backtracking from the goal to the initial position
			for traveledStep.CurrentPosition != e.agent.position.CurrentPosition {
				// Store the current position in the path
				traveledPath = append([]datatypes.BoardCoordinate{traveledStep.CurrentPosition}, traveledPath...)

				// Find the previous step by matching the PreviousPosition
				predecessor, found := Find(parentNodes, func(agent datatypes.AgentStep) bool {
					return agent.CurrentPosition == traveledStep.PreviousPosition
				})

				if found {
					traveledStep = predecessor // Move to the previous step
				} else {
					break // Exit if no predecessor is found
				}
			}

			// Add the initial position to the path
			traveledPath = append([]datatypes.BoardCoordinate{e.agent.position.CurrentPosition}, traveledPath...)

			return SearchResult{
				traveledPath,
				true,
				expandenNodes,
				len(traveledPath),
				1.333,
				end.Sub(start),
			}
		}

		fmt.Println("\nPosition: ", currentStep)
		e.agent.position = currentStep
		agentPerception := Percept(e.agent, e.board)
		expandenNodes++
		for _, perception := range agentPerception {
			fmt.Println("perception ", perception)
			fmt.Println("Previous step ", currentStep.PreviousPosition)
			fmt.Println("parentNodes ", parentNodes)
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
		fmt.Println("Queu ", queue)
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
