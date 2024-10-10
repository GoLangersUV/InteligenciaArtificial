package main

import (
	"log"

	"github.com/Krud3/InteligenciaArtificial/src/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	matrixFileName := "Prueba1.txt"

	game, err := game.NewGame(matrixFileName)
	if err != nil {
		log.Fatal(err)
	}

	game.SetCarPath("callDummy")

	// game.SetScene(matrixFileName)

	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Searching Algorithms")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
