package main

import (
	"log"

	"github.com/Krud3/InteligenciaArtificial/src/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game, err := game.NewGame()
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Your Game Title")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
