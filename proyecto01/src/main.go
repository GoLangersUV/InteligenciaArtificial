package main

import (
	"flag"
	"fmt"
	"image"
	"log"

	"github.com/Krud3/InteligenciaArtificial/src/game"
	"github.com/Krud3/InteligenciaArtificial/src/searchAlgorithms"
	"github.com/Krud3/InteligenciaArtificial/src/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

var g *game.Game

func main() {

	cmd := flag.Bool("cmd", false, "interface option")

	// Parse the flags
	flag.Parse()

	// Check if verbose mode is enabled
	if *cmd {
		matrix, err := utils.GetMatrix("../battery/Prueba1.txt") // Load the matrix
		if err != nil {
			log.Fatal(err)
		}
		result := searchAlgorithms.StartSearch(2, matrix)
		fmt.Println(result)
	} else {
		icon := utils.LoadIcon("./game/assets/images/cantidad-nodos.png")

		matrixFileName := "Prueba1.txt"

		var err error
		g, err = game.NewGame(matrixFileName)
		if err != nil {
			log.Fatal(err)
		}
		g.SetCarPath("callDummy")
		ebiten.SetWindowSize(640, 640)
		ebiten.SetWindowTitle("DidIA Game")
		ebiten.SetWindowIcon([]image.Image{icon})
		if err := ebiten.RunGame(g); err != nil {
			log.Fatal(err)
		}
	}
}
