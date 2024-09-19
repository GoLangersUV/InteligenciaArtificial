package searchAlgorithms

import (
  "fmt"
	"time"
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
	Matrix        *[10][10]int
	InitPosition  Position
	DogPosition   Position
	GoalPosition  Position
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
