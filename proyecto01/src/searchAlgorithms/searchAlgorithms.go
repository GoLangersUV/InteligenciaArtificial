package searchAlgorithms

import (
	"fmt"
	"github.com/Krud3/InteligenciaArtificial/src/datatypes"
	"time"
)

// enviroment represents the enviroment where the agent is going to move
type enviroment struct {
	agent     agent
	board     [][]int
	totalGoal int
}

// A point in the game board
type BoardCoordinate struct {
	x, y int
}

// SearchAlgorithm is the interface that the search algorithms must implement
type SearchAgorithm interface {
	LookForGoal(*enviroment) SearchResult
}

// agent represents the agent that is going to move in the enviroment
type agent struct {
	position          BoardCoordinate
	searchAlgorithm   SearchAgorithm
	ambientPerception [4]int
}

// MoveUp, MoveRight, MoveDown y MoveLeft son los actuadores del agente
func (a *agent) MoveUp() {
	a.position.x--
}
func (a *agent) MoveRight() {
	a.position.y++
}
func (a *agent) MoveDown() {
	a.position.x++
}
func (a *agent) MoveLeft() {
	a.position.y--
}

var contiguousMovements = [4]BoardCoordinate{BoardCoordinate{0, -1}, BoardCoordinate{1, 0}, BoardCoordinate{0, 1}, BoardCoordinate{-1, 0}}

func coordinateAdd(firstCoordinate BoardCoordinate, secondCoordinates BoardCoordinate) BoardCoordinate {
	return BoardCoordinate{firstCoordinate.x + secondCoordinates.x, firstCoordinate.y + secondCoordinates.y}
}

func Percept(a agent, board [][]int) []BoardCoordinate {
	canMove := []BoardCoordinate{}
	for _, contiguous := range contiguousMovements {
		tryCoordinate := coordinateAdd(a.position, contiguous)
		if board[tryCoordinate.x][tryCoordinate.y] == 0 {
			canMove = append(canMove, tryCoordinate)
		}
	}
	return canMove
}

type BreadthFirstSearch struct {
}

type SearchResult struct {
	solutionFound            bool
	expandenNodes, treeDepth int
	cost                     float32
	timeExe                  time.Duration
}

func (a *BreadthFirstSearch) LookForGoal(e *enviroment) SearchResult {
	parentNodes := make(datatypes.Set)
	expandenNodes := 0
	treeDepth := 0
	initialPosition := e.agent.position
	queue := datatypes.Queue[BoardCoordinate]{}
	queue.Enqueue(BoardCoordinate{initialPosition.x, initialPosition.y})
	start := time.Now()
	for !queue.IsEmpty() {
		currentPosition, error := queue.Dequeue()
		if error {
			return SearchResult{}
		}
		parentNodes.Add(currentPosition)
		if e.board[currentPosition.x][currentPosition.x] == 6 {
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

func StartGame() {
	board, error := GetMatrix()
	if error != nil {
		fmt.Println("Error al cargar el tablero")
	}
	searchStrategy := BreadthFirstSearch{}
	agent := agent{
		BoardCoordinate{0, 0},
		&searchStrategy,
		[4]int{},
	}
	agent.searchAlgorithm.LookForGoal(&enviroment{agent, board, 0})
}
