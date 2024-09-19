package main

import (
	"log"
	search "github.com/Krud3/InteligenciaArtificial/src/searchAlgorithms"
)

func main() {
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

	env, err := search.NewEnvironment(envMatrix)
	if err != nil {
		log.Fatalf("Error al crear el entorno: %v", err)
	}

	amplitudeSearch := new(search.AmplitudeSearch)
	agent := search.NewAgent(env.InitPosition, amplitudeSearch)

	agent.SearchAlgorithm.LookForGoal(env)
}

