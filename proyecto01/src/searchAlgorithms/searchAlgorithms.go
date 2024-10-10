package searchAlgorithms

import (
	"time"
)

// enviroment represents the enviroment where the agent is going to move
type enviroment struct {
	matrix    [][]int
	totalGoal int
}

// SearchAlgorithm is the interface that the search algorithms must implement
type SearchAgorithm interface {
	LookForGoal(*enviroment) (solutionFound bool, expandenNodes, treeDepth, cost float32, timeExe time.Duration)
}

// agent represents the agent that is going to move in the enviroment
type agent struct {
	x, y              int
	searchAlgorithm   SearchAgorithm
	ambientPerception [4]int
}

// MoveUp, MoveRight, MoveDown y MoveLeft son los actuadores del agente
func (a *agent) MoveUp() {
	a.x--
}
func (a *agent) MoveRight() {
	a.y++
}
func (a *agent) MoveDown() {
	a.x++
}
func (a *agent) MoveLeft() {
	a.y--
}

type AmplitudeSearch struct {
}

func (a *AmplitudeSearch) LookForGoal(e *enviroment) (solutionFound bool, expandenNodes, treeDepth, cost float32, timeExe time.Duration) {
	return
}

func DummyAlgorithm() [][]int {
	carPath := [][]int{
		{2, 0},
		{3, 0},
		{4, 0},
		{5, 0},
		{6, 0},
		{5, 0},
		{4, 0},
		{3, 0},
		{3, 1},
		{3, 2},
		{3, 3},
		{2, 3},
		{1, 3},
		{1, 4},
		{1, 5},
		{2, 5},
		{3, 5},
		{3, 6},
		{3, 7},
		{2, 7},
		{1, 7},
		{1, 8},
		{1, 9},
		{2, 9},
		{3, 9},
		{4, 9},
		{5, 9},
	}
	return carPath
}
