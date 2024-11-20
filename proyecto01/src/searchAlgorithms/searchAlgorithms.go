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

type Node struct {
	Position             Position
	Parent               *Node
	G                    float32 // Costo desde el inicio hasta el nodo actual
	H                    float32 // Costo heurístico al objetivo
	F                    float32 // Costo total (F = G + H)
	Depth                int
	HasPickedUpPassenger bool
}

type PriorityQueue struct {
	nodes   []*Node
	compare func(a, b *Node) bool
}

func (pq PriorityQueue) Len() int { return len(pq.nodes) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq.compare(pq.nodes[i], pq.nodes[j])
}

type State struct {
	Position             Position
	HasPickedUpPassenger bool
}

func (pq PriorityQueue) Swap(i, j int) {
	pq.nodes[i], pq.nodes[j] = pq.nodes[j], pq.nodes[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*Node)
	pq.nodes = append(pq.nodes, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := pq.nodes
	n := len(old)
	node := old[n-1]
	pq.nodes = old[0 : n-1]
	return node
}

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
	position        datatypes.AgentStep
	passenger       bool
	searchAlgorithm SearchAgorithm
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
	return canMove
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
		if found {
			currentStep = previousStep
		} else {
			break
		}
	}
	return path
}

func StartSearch(strategy int, scannedMatrix datatypes.ScannedMatrix) datatypes.SearchResult {

	var searchStrategy SearchAgorithm

	switch strategy {
	case 1:
		searchStrategy = &BreadthFirstSearch{}
	case 2:
		searchStrategy = &UniformCostSearch{}
	case 3:
		searchStrategy = &DepthSearch{}
	case 4:
		searchStrategy = &BreadthFirstSearch{}
	default:
		fmt.Println("Unknown strategy")
		return datatypes.SearchResult{}
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
		}
		result := agent.searchAlgorithm.LookForGoal(&enviroment{
			agent,
			datatypes.BoardCoordinate{X: math.MaxInt, Y: math.MaxInt},
			scannedMatrix.Matrix,
			0},
		)
		return result
	} else {
		fmt.Println("Initial position not found")
		return datatypes.SearchResult{}
	}
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
