package main

import (
	"fmt"

	"github.com/puyojulian/p1ia/busqueda/src/search"
)

func main() {
	matrix, err := search.GetMatrix()
	if err != nil {
		fmt.Println("Error getting matrix:", err)
		return
	}

	// Print the matrix
	for _, row := range matrix {
		fmt.Println(row)
	}
}
