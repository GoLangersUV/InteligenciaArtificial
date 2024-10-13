package main

import (
	"image"
	"log"

	"github.com/Krud3/InteligenciaArtificial/src/game"
	"github.com/Krud3/InteligenciaArtificial/src/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

var g *game.Game

func main() {

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
