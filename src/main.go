package main

import (
	algorithms "github.com/Krud3/InteligenciaArtificial/src/searchAlgorithms"
)

func main() {
// 1 es un obstáculo y 2 es el objetivo.
  envMatrix := [10][10]int{
    {0, 1, 1, 1, 1, 1, 1, 1, 1, 1},
    {0, 1, 1, 0, 0, 0, 4, 0, 0, 0},
    {2, 1, 1, 0, 1, 0, 1, 0, 1, 0},
    {0, 3, 3, 0, 4, 0, 0, 0, 4, 0},
    {0, 1, 1, 0, 1, 1, 1, 1, 1, 0},
    {0, 0, 0, 0, 1, 1, 0, 0, 0, 6},
    {5, 1, 1, 1, 1, 1, 0, 1, 1, 1},
    {0, 1, 0, 0, 0, 1, 0, 0, 0, 1},
    {0, 1, 0, 1, 0, 1, 1, 1, 0, 1},
    {0, 0, 0, 1, 0, 0, 0, 0, 0, 1},
  }
  enviroment := algorithms.NewEnviroment(envMatrix)

  amplitudeSearch := new(algorithms.AmplitudeSearch)

  agent := algorithms.NewAgent(enviroment.GetInitPosition()[0], enviroment.GetInitPosition()[1], amplitudeSearch)

  agent.SearchAlgorithm.LookForGoal(enviroment)
}
