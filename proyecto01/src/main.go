package main

import (
	"flag"
	"fmt"
	"github.com/Krud3/InteligenciaArtificial/src/searchAlgorithms"
)

func main() {
	fmt.Println("Selecciona un numero de strategy")

	strategy := flag.Int("strategy", 1, "Search Strategy")
	// Additional logic based on the command-line arguments
	if *strategy >= 1 && *strategy < 5 {
		// Define flags for command-line arguments
		// Parse the command-line arguments
		flag.Parse()
		searchAlgorithms.StartGame(*strategy)
	} else {
		fmt.Println("Enter a valid strategy value")
	}
}
