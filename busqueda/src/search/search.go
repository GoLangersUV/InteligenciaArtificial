package search

type State struct {
	i int
	j int
}

// Function to generate possible next states (up, down, left, right)
func operands(s State) []State {
	return []State{
		{s.i - 1, s.j}, // up
		{s.i + 1, s.j}, // down
		{s.i, s.j - 1}, // left
		{s.i, s.j + 1}, // right
	}
}

// Function to get possible actions based on the current state
func get_actions(matrix *[][]int, s State) []State {
	var actions []State

	// Get all possible next states from the current state
	states := operands(s)

	// Iterate over each possible state and check if it's valid (matrix value is 0)
	for _, each := range states {
		if each.i >= 0 && each.i < len(*matrix) && each.j >= 0 && each.j < len((*matrix)[0]) { // Check boundaries
			if (*matrix)[each.i][each.j] == 0 { // Only consider positions where the matrix value is 0
				actions = append(actions, each)
			}
		}
	}

	return actions
}
