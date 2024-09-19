package searchAlgorithms

import (
	"time"
)

type AmplitudeSearch struct {
	SolutionPath []Position
}

func (b *AmplitudeSearch) LookForGoal(env *Environment) SearchResult {
	startTime := time.Now()
	
	// Implementación de BFS
	queue := []Position{env.InitPosition}
	visited := make(map[Position]bool)
	visited[env.InitPosition] = true
	parent := make(map[Position]Position)
	var solutionFound bool
	var expandedNodes int
	var cost float32

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		expandedNodes++

		if current == env.GoalPosition {
			solutionFound = true
			break
		}

		neighbors := getNeighbors(current, env)
		for _, neighbor := range neighbors {
			if !visited[neighbor] {
				visited[neighbor] = true
				parent[neighbor] = current
				queue = append(queue, neighbor)
			}
		}
	}

	path := []Position{}
	if solutionFound {
		pos := env.GoalPosition
		for pos != env.InitPosition {
			path = append([]Position{pos}, path...)
			pos = parent[pos]
		}
		path = append([]Position{env.InitPosition}, path...)
		b.SolutionPath = path
		cost = float32(len(path))
	}

	timeExe := time.Since(startTime)

	return SearchResult{
		SolutionFound: solutionFound,
		ExpandedNodes: expandedNodes,
		TreeDepth:     len(path),
		Cost:          cost,
		TimeExecuted:  timeExe,
		Path:          path,
	}
}

func getNeighbors(pos Position, env *Environment) []Position {
	directions := []Position{
		{X: pos.X - 1, Y: pos.Y}, // Up
		{X: pos.X, Y: pos.Y + 1}, // Right
		{X: pos.X + 1, Y: pos.Y}, // Down
		{X: pos.X, Y: pos.Y - 1}, // Left
	}
	validNeighbors := []Position{}
	for _, neighbor := range directions {
		if neighbor.X >= 0 && neighbor.X < len(env.Matrix) &&
			neighbor.Y >= 0 && neighbor.Y < len(env.Matrix[0]) &&
			env.Matrix[neighbor.X][neighbor.Y] != WALL {
			validNeighbors = append(validNeighbors, neighbor)
		}
	}
	return validNeighbors
}

