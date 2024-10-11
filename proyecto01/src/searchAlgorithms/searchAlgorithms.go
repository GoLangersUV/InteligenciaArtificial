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
	passenger         bool
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
		if tryCoordinate.X >= 0 &&
			tryCoordinate.Y >= 0 &&
			tryCoordinate.X < len(board) &&
			tryCoordinate.Y < len(board) &&
			board[tryCoordinate.X][tryCoordinate.Y] != 1 {
			canMove = append(canMove, tryCoordinate)
		}
	}
	fmt.Println("Can move: %s", canMove)
	return canMove
}

type BreadthFirstSearch struct{}

type UniformCostSearch struct{}

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
		currentPosition, empty := queue.Dequeue()
		if empty {
			return SearchResult{
				false,
				expandenNodes,
				treeDepth,
				0.0,
				time.Now().Sub(start),
			}
		}
		parentNodes.Add(currentPosition)
		fmt.Printf("current: %s", currentPosition)

		if e.board[currentPosition.X][currentPosition.Y] == 5 {
			e.agent = agent{
				initialPosition,
				true,
				e.agent.searchAlgorithm,
				e.agent.ambientPerception,
			}
			parentNodes = make(datatypes.Set)
		}

		if e.agent.passenger && e.board[currentPosition.X][currentPosition.Y] == 6 {
			end := time.Now()
			return SearchResult{
				true,
				expandenNodes,
				treeDepth,
				1.0,
				end.Sub(start),
			}
		}
		fmt.Println("Position: %s", currentPosition)
		e.agent.position = currentPosition
		agentPerception := Percept(e.agent, e.board)
		expandenNodes++
		for _, perception := range agentPerception {
			if !parentNodes.Contains(perception) {
				queue.Enqueue(perception)
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
	priorityQueue := datatypes.PriorityQueue[datatypes.BoardCoordinate]{}
	priorityQueue.Push(datatypes.Element[datatypes.BoardCoordinate]{
		Value:    initialPosition,
		Priority: 1},
	)
	for !priorityQueue.IsEmpty() {
		currentPosition, empty := priorityQueue.Pop()
		if empty {
			return SearchResult{
				false,
				expandenNodes,
				treeDepth,
				0.0,
				0,
			}
		}
		parentNodes.Add(currentPosition)
		e.agent.position = currentPosition
		agentPerception := Percept(e.agent, e.board)
		expandenNodes++
		for _, perception := range agentPerception {
			if !parentNodes.Contains(perception) {
				priorityQueue.Push(
					datatypes.Element[datatypes.BoardCoordinate]{
						Value:    perception,
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
		agent := agent{
			initialPosition,
			false,
			searchStrategy,
			[4]int{},
		}
		result := agent.searchAlgorithm.LookForGoal(&enviroment{agent, scannedMatrix.Matrix, 0})
		fmt.Println(result)
	} else {

	}

}
