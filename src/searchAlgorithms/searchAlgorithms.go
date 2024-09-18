package searchAlgorithms

import (
  "time"
)

// enviroment represents the enviroment where the agent is going to move
type enviroment struct {
  matrix [][]int
  totalGoal int
}

// SearchAlgorithm is the interface that the search algorithms must implement
type SearchAgorithm interface {
  LookForGoal(*enviroment) (solutionFound bool, expandenNodes, treeDepth, cost float32, timeExe time.Duration) 
}

// agent represents the agent that is going to move in the enviroment
type agent struct {
  x, y int
  searchAlgorithm SearchAgorithm
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
