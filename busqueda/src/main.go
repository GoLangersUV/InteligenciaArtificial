package main

import (
	"fmt"

	"github.com/Krud3/InteligenciaArtificial/busqueda/src/search"
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
