package searchAlgorithms

import (
	"fmt"
	"time"

	"github.com/Krud3/InteligenciaArtificial/src/datatypes"
)

// enviroment represents the enviroment where the agent is going to move
type enviroment struct {
	agent     agent
	board     [][]int
	totalGoal int
}

// SearchAlgorithm is the interface that the search algorithms must implement
type SearchAgorithm interface {
	LookForGoal(*enviroment) SearchResult
}

// agent represents the agent that is going to move in the enviroment
type agent struct {
	position          datatypes.BoardCoordinate
	searchAlgorithm   SearchAgorithm
	ambientPerception [4]int
}

// MoveUp, MoveRight, MoveDown y MoveLeft son los actuadores del agente
func (a *agent) MoveUp() {
	a.position.X--
}
func (a *agent) MoveRight() {
	a.position.Y++
}
func (a *agent) MoveDown() {
	a.position.X++
}
func (a *agent) MoveLeft() {
	a.position.Y--
}

var contiguousMovements = [4]datatypes.BoardCoordinate{{X: 0, Y: -1}, {X: 1, Y: 0}, {X: 0, Y: 1}, {X: -1, Y: 0}}

func coordinateAdd(firstCoordinate datatypes.BoardCoordinate, secondCoordinates datatypes.BoardCoordinate) datatypes.BoardCoordinate {
	return datatypes.BoardCoordinate{X: firstCoordinate.X + secondCoordinates.X, Y: firstCoordinate.Y + secondCoordinates.Y}
}

func Percept(a agent, board [][]int) []datatypes.BoardCoordinate {
	canMove := []datatypes.BoardCoordinate{}
	for _, contiguous := range contiguousMovements {
		tryCoordinate := coordinateAdd(a.position, contiguous)
		if board[tryCoordinate.X][tryCoordinate.Y] == 0 {
			canMove = append(canMove, tryCoordinate)
		}
	}
	return canMove
}

type BreadthFirstSearch struct{}

type DepthSearch struct{}

type SearchResult struct {
	solutionFound            bool
	expandenNodes, treeDepth int
	cost                     float32
	timeExe                  time.Duration
}

func (a *BreadthFirstSearch) LookForGoal(e *enviroment) SearchResult {
	parentNodes := make(datatypes.Set) // For Hold the parents node removed from the queue
	expandenNodes := 0
	treeDepth := 0
	initialPosition := e.agent.position
	queue := datatypes.Queue[datatypes.BoardCoordinate]{}
	queue.Enqueue(datatypes.BoardCoordinate{X: initialPosition.X, Y: initialPosition.Y})
	start := time.Now()
	for !queue.IsEmpty() {
		currentPosition, error := queue.Dequeue()
		if error {
			return SearchResult{}
		}
		parentNodes.Add(currentPosition)
		if e.board[currentPosition.X][currentPosition.X] == 6 {
			end := time.Now()
			return SearchResult{
				true,
				expandenNodes,
				treeDepth,
				1.0,
				end.Sub(start),
			}
		}
		e.agent.position = currentPosition
		agentPerception := Percept(e.agent, e.board)
		expandenNodes++
		for _, perception := range agentPerception {
			if !parentNodes.Contains(perception) {
				queue.Enqueue(perception)
			}
		}
	}
	return SearchResult{}
}

func (a *DepthSearch) LookForGoal(e *enviroment) SearchResult {
	return SearchResult{}
}

func StartGame(strategy int) {
	scannedMatrix, error := GetMatrix()
	if error != nil {
		fmt.Println("Error al cargar el tablero")
	}

	var searchStrategy SearchAgorithm

	switch strategy {
	case 1:
		searchStrategy = &BreadthFirstSearch{}
	case 2:
		searchStrategy = &DepthSearch{}
	case 4:
		searchStrategy = &BreadthFirstSearch{}
	default:
		fmt.Println("Unknown strategy")
	}

	agent := agent{
		datatypes.BoardCoordinate{X: 0, Y: 0},
		searchStrategy,
		[4]int{},
	}
	result := agent.searchAlgorithm.LookForGoal(&enviroment{agent, scannedMatrix.Matrix, 0})
	fmt.Println(result)
}
