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
        log.Fatalf("Error creating environment: %v", err)
    }

    aStarSearch := new(search.AStarSearch)
    agent := search.NewAgent(env.InitPosition, aStarSearch)

    result := agent.SearchAlgorithm.LookForGoal(env)

    if result.SolutionFound {
        log.Printf("Solution found!")
        log.Printf("Path: %v", result.Path)
        log.Printf("Cost: %v", result.Cost)
        log.Printf("Expanded Nodes: %d", result.ExpandedNodes)
        log.Printf("Tree Depth: %d", result.TreeDepth)
        log.Printf("Time Executed: %v", result.TimeExecuted)
    } else {
        log.Printf("No solution found.")
    }
}
