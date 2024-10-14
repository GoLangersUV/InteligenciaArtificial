package searchAlgorithms

import (
	"fmt"
	"math"
	"time"

	"github.com/Krud3/InteligenciaArtificial/src/datatypes"
)

const (
	WALL = iota + 1
	INIT_POSITION
	MIDCOST
	HEAVYCOST
	DOG
	GOAL
)

// Position representa una posición en la matriz.
type Position struct {
	X, Y int
}

// Environment representa el entorno donde el agente se moverá.
type Environment struct {
	Matrix       *[10][10]int
	InitPosition Position
	DogPosition  Position
	GoalPosition Position
}

// NewEnvironment crea un nuevo entorno a partir de una matriz.
func NewEnvironment(matrix [10][10]int) (*Environment, error) {
	var initPos, dogPos, goalPos Position
	foundInit, foundDog, foundGoal := false, false, false

	for i, row := range matrix {
		for j, cell := range row {
			switch cell {
			case INIT_POSITION:
				initPos = Position{X: i, Y: j}
				foundInit = true
			case DOG:
				dogPos = Position{X: i, Y: j}
				foundDog = true
			case GOAL:
				goalPos = Position{X: i, Y: j}
				foundGoal = true
			}
		}
	}

	if !foundInit || !foundDog || !foundGoal {
		return nil, fmt.Errorf("environment must have init, dog, and goal positions")
	}

	return &Environment{
		Matrix:       &matrix,
		InitPosition: initPos,
		DogPosition:  dogPos,
		GoalPosition: goalPos,
	}, nil
}

// Perception representa la percepción del agente en las cuatro direcciones.
type Perception struct {
	Up, Right, Down, Left bool
}

// Agent representa al agente que se moverá en el entorno.
type Agent struct {
	Position        Position
	SearchAlgorithm SearchAlgorithm
	Perception      Perception
	Dog             bool
}

// enviroment represents the enviroment where the agent is going to move
type enviroment struct {
	agent             agent
	passengerPosition datatypes.BoardCoordinate
	board             [][]int
	totalGoal         int
}

// SearchAlgorithm is the interface that the search algorithms must implement
type SearchAgorithm interface {
	LookForGoal(*enviroment) datatypes.SearchResult
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

func (a *BreadthFirstSearch) LookForGoal(e *enviroment) datatypes.SearchResult {
	var parentNodes []datatypes.AgentStep
	// var paths [][]datatypes.AgentStep // To keep track of both the path to the passenger and from the passenger to the goal
	expandenNodes := 0

	initialPosition := e.agent.position
	queue := datatypes.Queue[datatypes.AgentStep]{}
	queue.Enqueue(initialPosition)

	start := time.Now()

	pathToPassenger := []datatypes.BoardCoordinate{} // Path from initial position to the passenger
	pathToGoal := []datatypes.BoardCoordinate{}      // Path from passenger to the goal

	for !queue.IsEmpty() {
		currentStep, empty := queue.Dequeue()
		if empty {
			return datatypes.SearchResult{
				PathFound:     []datatypes.BoardCoordinate{},
				SolutionFound: false,
				ExpandenNodes: expandenNodes,
				TreeDepth:     0,
				Cost:          0,
				TimeExe:       time.Since(start),
			}
		}

		parentNodes = append(parentNodes, currentStep)

		// Phase 1: Find the passenger
		if !e.agent.passenger && e.board[currentStep.CurrentPosition.X][currentStep.CurrentPosition.Y] == 5 {
			e.passengerPosition = datatypes.BoardCoordinate{
				X: currentStep.CurrentPosition.X,
				Y: currentStep.CurrentPosition.Y,
			}
			e.agent.passenger = true // Mark the agent as having picked up the passenger

			// Reconstruct path from the initial position to the passenger
			pathToPassenger = reconstructPath(parentNodes, currentStep)

			fmt.Println("Path to passenger: %s", pathToPassenger)
			fmt.Println("Nodes a Passenger: %s", parentNodes)
			// Clear the queue and start BFS again from the passenger's position
			queue.Clear()
			parentNodes = nil // Clear parent nodes for the next phase
			queue.Enqueue(
				datatypes.AgentStep{
					PreviousPosition: datatypes.BoardCoordinate{
						X: math.MaxInt,
						Y: math.MaxInt,
					},
					CurrentPosition: currentStep.CurrentPosition,
					Depth:           currentStep.Depth + 1,
					Action:          math.MaxInt,
				},
			) // Start from the passenger's position
			continue
		}

		// Phase 2: Search for the goal (once the passenger is found)
		if e.agent.passenger && e.board[currentStep.CurrentPosition.X][currentStep.CurrentPosition.Y] == 6 {
			end := time.Now()
			// Reconstruct path from the passenger to the goal
			pathToGoal = reconstructPath(parentNodes, currentStep)

			// Combine the two paths: initial -> passenger + passenger -> goal
			combinedPath := append(pathToPassenger, pathToGoal...)

			return datatypes.SearchResult{
				PathFound:     combinedPath,
				SolutionFound: true,
				ExpandenNodes: expandenNodes,
				TreeDepth:     len(combinedPath),
				Cost:          1.333, // Example cost, you can modify it as needed
				TimeExe:       end.Sub(start),
			}
		}

		// Continue exploring neighbors
		e.agent.position = currentStep
		agentPerception := Percept(e.agent, e.board)
		expandenNodes++
		for _, perception := range agentPerception {
			if (perception.Coordinate.X != currentStep.PreviousPosition.X) || (perception.Coordinate.Y != currentStep.PreviousPosition.Y) {
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
	return datatypes.SearchResult{}
}

// Helper function to reconstruct the path from a given node back to the start
func reconstructPath(parentNodes []datatypes.AgentStep, endStep datatypes.AgentStep) []datatypes.BoardCoordinate {
	path := []datatypes.BoardCoordinate{}
	currentStep := endStep

	visited := make(map[datatypes.BoardCoordinate]bool) // Track visited steps to avoid infinite loop

	for {
		path = append([]datatypes.BoardCoordinate{currentStep.CurrentPosition}, path...)
		if visited[currentStep.CurrentPosition] {
			// Stop if we have visited the same step again
			break
		}
		visited[currentStep.CurrentPosition] = true

		if currentStep.PreviousPosition.X == math.MaxInt && currentStep.PreviousPosition.Y == math.MaxInt {
			break // We've reached the starting point
		}

		// Find the parent step
		previousStep, found := Find(parentNodes, func(step datatypes.AgentStep) bool {
			return step.CurrentPosition == currentStep.PreviousPosition
		})
		fmt.Println("Found: ", previousStep, found)
		if found {
			currentStep = previousStep
		} else {
			break
		}
	}
	return path
}

func (a *UniformCostSearch) LookForGoal(e *enviroment) datatypes.SearchResult {
	// Step 1: Find the path to the passenger
	pathToPassenger := a.findPath(e, 5)
	fmt.Println("Path to passenger: %s", pathToPassenger)
	// If the path to the passenger is not found, return empty result
	if len(pathToPassenger) == 0 {
		return datatypes.SearchResult{
			PathFound:     []datatypes.BoardCoordinate{},
			SolutionFound: false,
			ExpandenNodes: 0,
			TreeDepth:     0,
			Cost:          0.0,
			TimeExe:       0,
		}
	}

	// Step 2: Find the path from the passenger to the goal
	pathToGoal := a.findPath(e, 6)

	// If the path to the goal is not found, return the path to the passenger
	if len(pathToGoal) == 0 {
		return datatypes.SearchResult{
			PathFound:     pathToPassenger,
			SolutionFound: true,
			ExpandenNodes: 0, // You might want to count these nodes as well
			TreeDepth:     0,
			Cost:          0.0,
			TimeExe:       0,
		}
	}

	// Step 3: Combine the two paths
	combinedPath := append(pathToPassenger, pathToGoal[1:]...) // Avoid duplicate passenger position

	return datatypes.SearchResult{
		PathFound:     combinedPath,
		SolutionFound: true,
		ExpandenNodes: 0, // You might want to count these nodes as well
		TreeDepth:     0,
		Cost:          0.0,
		TimeExe:       0,
	}
}

// Helper function to perform Uniform Cost Search from start to goal
func (a *UniformCostSearch) findPath(e *enviroment, goal int) []datatypes.BoardCoordinate {
	var parentNodes []datatypes.AgentStep // Holds parent nodes
	expandenNodes := 0
	priorityQueue := datatypes.PriorityQueue[datatypes.AgentStep]{}
	fmt.Println("Finding path from start to goal", e.agent.position.CurrentPosition, e.agent.position.PreviousPosition)
	// The expanded nodes at any given time
	priorityQueue.Push(datatypes.Element[datatypes.AgentStep]{

		Value: datatypes.AgentStep{
			Depth:  e.agent.position.Depth,
			Action: e.agent.position.Action,
			Cost:   e.agent.position.Cost,
			PreviousPosition: datatypes.BoardCoordinate{
				X: math.MaxInt,
				Y: math.MaxInt,
			},
			CurrentPosition: e.agent.position.CurrentPosition,
		},
		Priority: 0}, // Priority can be based on cost from start
	)

	for !priorityQueue.IsEmpty() {
		currentStep, _ := priorityQueue.Pop()
		// Check if the current position is the goal
		if e.board[currentStep.CurrentPosition.X][currentStep.CurrentPosition.Y] == goal {
			fmt.Println("Found final goal!")
			e.agent.position = currentStep
			return reconstructPath(parentNodes, currentStep) // Implement this function to trace back the path
		}

		parentNodes = append(parentNodes, currentStep)
		expandenNodes++

		e.agent.position = currentStep
		agentPerception := Percept(e.agent, e.board)
		fmt.Println("Parent Nodes: ", parentNodes)

		for _, perception := range agentPerception {
			_, isContained := Find(parentNodes, func(agent datatypes.AgentStep) bool {
				return agent.CurrentPosition == perception.Coordinate
			})
			if !isContained {
				fmt.Println("Pushing: ", perception)
				movementCost := currentStep.Cost + checkCellWeight(e.board[perception.Coordinate.X][perception.Coordinate.Y], goal)
				priorityQueue.Push(
					datatypes.Element[datatypes.AgentStep]{
						Value: datatypes.AgentStep{
							Depth:            currentStep.Depth + 1,
							Action:           perception.Direction,
							Cost:             movementCost,
							CurrentPosition:  perception.Coordinate,
							PreviousPosition: currentStep.CurrentPosition,
						},
						Priority: movementCost,
					},
				)
			}
		}
		fmt.Println("Priority Queue: ", priorityQueue)
		//time.Sleep(10 * time.Millisecond)
	}
	return []datatypes.BoardCoordinate{} // No path found to goal
}

func (a *DepthSearch) LookForGoal(e *enviroment) datatypes.SearchResult {
	return datatypes.SearchResult{}
}

func StartSearch(strategy int, scannedMatrix datatypes.ScannedMatrix) datatypes.SearchResult {

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
			Cost:            0,
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
		return result
	} else {

	}
	return datatypes.SearchResult{}
}

func checkCellWeight(value int, goalValue int) int {
	if value != goalValue && value != 2 && value != 0 {
		return value
	}
	return 1
}

// SearchResult encapsula los resultados de una búsqueda.
type SearchResult struct {
	SolutionFound bool
	ExpandedNodes int
	TreeDepth     int
	Cost          float32
	TimeExecuted  time.Duration
	Path          []Position
}

// SearchAlgorithm es la interfaz que deben implementar los algoritmos de búsqueda.
type SearchAlgorithm interface {
	LookForGoal(env *Environment) SearchResult
}

// NewAgent crea un nuevo agente.
func NewAgent(pos Position, algo SearchAlgorithm) *Agent {
	return &Agent{
		Position:        pos,
		SearchAlgorithm: algo,
		Perception:      Perception{},
		Dog:             false,
	}
}

// GeneratePerception actualiza la percepción del agente basado en su posición.
func (a *Agent) GeneratePerception(env *Environment) {
	matrix := env.Matrix
	x, y := a.Position.X, a.Position.Y
	a.Perception.Up = x > 0 && matrix[x-1][y] != WALL
	a.Perception.Right = y < len(matrix[0])-1 && matrix[x][y+1] != WALL
	a.Perception.Down = x < len(matrix)-1 && matrix[x+1][y] != WALL
	a.Perception.Left = y > 0 && matrix[x][y-1] != WALL
}
