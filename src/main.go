package main

import (
  algorithms "github.com/Krud3/InteligenciaArtificial/src/searchAlgorithms"
)

func main() {
// 1 es un obstáculo y 2 es el objetivo.
  envMatrix := [][]int{
    {0,0,0,0},
    {0,1,1,0},
    {0,1,0,0},
    {2,0,2,1},
  }
  enviroment := algorithms.NewEnviroment(envMatrix)
  
  amplitudeSearch := new(algorithms.AmplitudeSearch)

  agent := algorithms.NewAgent(enviroment.GetAgentPosition()[0], enviroment.GetAgentPosition()[1], amplitudeSearch)

  agent.SearchAlgorithm.LookForGoal(enviroment)
}
