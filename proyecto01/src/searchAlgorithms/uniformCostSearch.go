package searchAlgorithms

import (
	"math"
	"time"

	"github.com/Krud3/InteligenciaArtificial/src/datatypes"
)

type UniformCostSearch struct{}

func (a *UniformCostSearch) LookForGoal(e *enviroment) datatypes.SearchResult {
	// Step 1: Find the path to the passenger
	start := time.Now()
	pathToPassenger, passengerExpandedNodes, passengerCost := a.findPath(e, 5)
	// If the path to the passenger is not found, return empty result
	if len(pathToPassenger) == 0 {
		return datatypes.SearchResult{
			PathFound:     []datatypes.BoardCoordinate{},
			SolutionFound: false,
			ExpandenNodes: passengerExpandedNodes,
			TreeDepth:     0,
			Cost:          float32(passengerCost),
			TimeExe:       time.Since(start),
		}
	}
	// Step 2: Find the path from the passenger to the goal
	pathToGoal, goalExpandedNodes, goalCost := a.findPath(e, 6)

	// If the path to the goal is not found, return the path to the passenger
	if len(pathToGoal) == 0 {
		return datatypes.SearchResult{
			PathFound:     pathToPassenger,
			SolutionFound: false,
			ExpandenNodes: goalExpandedNodes, // You might want to count these nodes as well
			TreeDepth:     len(pathToPassenger),
			Cost:          float32(goalCost),
			TimeExe:       time.Since(start),
		}
	}

	// Step 3: Combine the two paths
	combinedPath := append(pathToPassenger, pathToGoal[1:]...) // Avoid duplicate passenger position

	return datatypes.SearchResult{
		PathFound:     combinedPath,
		SolutionFound: true,
		ExpandenNodes: goalExpandedNodes + passengerExpandedNodes, // You might want to count these nodes as well
		TreeDepth:     len(combinedPath),
		Cost:          float32(goalCost + passengerCost),
		TimeExe:       time.Since(start),
	}
}

// Helper function to perform Uniform Cost Search from start to goal
func (a *UniformCostSearch) findPath(e *enviroment, goal int) ([]datatypes.BoardCoordinate, int, float32) {
	var parentNodes []datatypes.AgentStep // Holds parent nodes
	expandenNodes := 0
	priorityQueue := datatypes.PriorityQueue[datatypes.AgentStep]{}

	// Start with the initial position and cost 0
	initialCost := float32(0)
	priorityQueue.Push(datatypes.Element[datatypes.AgentStep]{
		Value: datatypes.AgentStep{
			Depth:  e.agent.position.Depth,
			Action: e.agent.position.Action,
			Cost:   int(initialCost),
			PreviousPosition: datatypes.BoardCoordinate{
				X: math.MaxInt,
				Y: math.MaxInt,
			},
			CurrentPosition: e.agent.position.CurrentPosition,
		},
		Priority: 0}, // Priority can be based on cost from start
	)

	visited := make(map[datatypes.BoardCoordinate]bool)

	for !priorityQueue.IsEmpty() {
		currentStep, _ := priorityQueue.Pop()

		// Check if the current position is the goal
		if e.board[currentStep.CurrentPosition.X][currentStep.CurrentPosition.Y] == goal {
			e.agent.position = currentStep
			return reconstructPath(parentNodes, currentStep), expandenNodes, float32(currentStep.Cost)
		}

		if visited[currentStep.CurrentPosition] {
			continue
		}

		visited[currentStep.CurrentPosition] = true
		parentNodes = append(parentNodes, currentStep)
		expandenNodes++

		e.agent.position = currentStep
		agentPerception := Percept(e.agent, e.board)

		for _, perception := range agentPerception {
			if !visited[perception.Coordinate] {
				movementCost := float32(currentStep.Cost) + getCellCost(e.board[perception.Coordinate.X][perception.Coordinate.Y])
				priorityQueue.Push(
					datatypes.Element[datatypes.AgentStep]{
						Value: datatypes.AgentStep{
							Depth:            currentStep.Depth + 1,
							Action:           perception.Direction,
							Cost:             int(movementCost), // Use the calculated movement cost
							CurrentPosition:  perception.Coordinate,
							PreviousPosition: currentStep.CurrentPosition,
						},
						Priority: int(movementCost), // Priority is based on the total cost to reach this node
					},
				)
			}
		}
	}

	// If no path found, return empty result
	return []datatypes.BoardCoordinate{}, expandenNodes, 0
}
