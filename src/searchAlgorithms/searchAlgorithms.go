package searchAlgorithms

import (
  "time"
)

const (
  WALL = iota + 1
  INITPOSITION
  MIDCOST
  HEAVYCOST
  DOG
  GOAL
)

// enviroment represents the enviroment where the agent is going to move
type enviroment struct {
  matrix [][]int
  initPosition, dogPosition, goalPosition [2]int
}

// NewEnviroment creates a new enviroment
func NewEnviroment(matrix [][]int) *enviroment {
  initPosition := [2]int{0, 0}
  dogPosition := [2]int{0, 0}
  goalPosition := [2]int{0, 0}

  for i, row := range matrix {
    for j, cell := range row {
      switch cell {
      case INITPOSITION:
        initPosition[0], initPosition[1] = i, j
      case DOG:
        dogPosition[0], dogPosition[1] = i, j
      case GOAL:
        goalPosition[0], goalPosition[1] = i, j
      }
    }
  }

  return &enviroment{
    matrix,
    initPosition,
    dogPosition,
    goalPosition,
  }
}

// GetAgentPosition returns the initial position of the agent
func (e *enviroment) GetAgentPosition() *[2]int {
  return &e.initPosition
}

// agent represents the agent that is going to move in the enviroment
type agent struct {
  x, y int
  SearchAlgorithm SearchAgorithm
  ambientPerception [4]byte
  dog bool
}

// NewAgent creates a new agent
func NewAgent(x, y int, searchAlgorithm SearchAgorithm) *agent {
  return &agent{x, y, searchAlgorithm, [4]byte{}, false}
}

// generatePerception generates the perception of the agent based on its position
func (a *agent) generatePerception(env *enviroment) {

  const (
    Up = iota
    Right
    Down
    Left
  )

  // simulate the agent senses
  moveUp := a.x - 1
  moveRight := a.y + 1
  moveDown := a.x + 1
  moveLeft := a.y - 1

  // generates the perception of the agent
  if moveUp >= 0 && env.matrix[moveUp][a.y] != wall {
    a.ambientPerception[Up] = 1 
  } else {
    a.ambientPerception[Up] = 0
  }
  if moveRight < len(env.matrix[0]) && env.matrix[a.x][moveRight] != wall {
    a.ambientPerception[Right] = 1
  } else {
    a.ambientPerception[Right] = 0
  }
  if moveDown < len(env.matrix) && env.matrix[moveDown][a.y] != wall {
    a.ambientPerception[Down] = 1 
  } else {
    a.ambientPerception[Down] = 0
  }
  if moveLeft >= 0 && env.matrix[a.x][moveLeft] != wall {
    a.ambientPerception[Left] = 1 
  } else {
    a.ambientPerception[Left] = 0
  }
}

// SearchAlgorithm is the interface that the search algorithms must implement
type SearchAgorithm interface {
  LookForGoal(*enviroment) (solutionFound bool, expandenNodes, treeDepth, cost float32, timeExe time.Duration) 
}
