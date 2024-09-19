package searchAlgorithms

import (
  "time"
)

const (
  wall = iota + 1
  goal
)

// enviroment represents the enviroment where the agent is going to move
type enviroment struct {
  matrix [][]int
  agentPosition, dogPosition, goalPosition [2]int
}

func NewEnviroment(matrix [][]int, totalGoal int) *enviroment {
  // TODO: Extract positions of agent, dog, and goal
  return &enviroment{
    matrix: matrix,
    agentPosition: [2]int{0, 0},
    dogPosition: [2]int{0, 0},
    goalPosition: [2]int{0, 0},
  }
}

// agent represents the agent that is going to move in the enviroment
type agent struct {
  x, y int
  searchAlgorithm SearchAgorithm
  ambientPerception [4]byte
  dog bool
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

// generatePerception genera una percepción basada en el entorno
func (a *agent) generatePerception(env enviroment) {
  // Simulan los sensores
  moveUp := a.x - 1
  moveRight := a.y + 1
  moveDown := a.x + 1
  moveLeft := a.y - 1

  const (
    Up = iota
    Right
    Down
    Left
  )

  // Se genera una percepción
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

func NewAgent(x, y int, searchAlgorithm SearchAgorithm) *agent {
  return &agent{x, y, searchAlgorithm, [4]byte{}, false}
}

// SearchAlgorithm is the interface that the search algorithms must implement
type SearchAgorithm interface {
  LookForGoal(*enviroment) (solutionFound bool, expandenNodes, treeDepth, cost float32, timeExe time.Duration) 
}
