package main

import (
	"log"

	"github.com/Krud3/InteligenciaArtificial/src/game"
	"github.com/hajimehoshi/ebiten/v2"
)

var g *game.Game

func main() {

	matrixFileName := "Prueba1.txt"

	var err error
	g, err = game.NewGame(matrixFileName)
	if err != nil {
		log.Fatal(err)
	}

	g.SetCarPath("callDummy")

	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Searching Algorithms")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
