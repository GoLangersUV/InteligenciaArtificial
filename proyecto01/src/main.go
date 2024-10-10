package main

import (
	"log"
	"syscall/js"

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

	// game.SetScene(matrixFileName)

	// This will expose the SetPath function to JavaScript
	js.Global().Set("setCarPath", js.FuncOf(setCarPath))
	// This will expose the SetPath function to JavaScript
	js.Global().Set("setCarPath", js.FuncOf(setScene))

	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Searching Algorithms")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}

	select {}
}

func setCarPath(this js.Value, p []js.Value) interface{} {
	path := p[0].String()
	g.SetCarPath(path)
	return nil
}

func setScene(this js.Value, p []js.Value) interface{} {
	path := p[0].String()
	g.SetScene(path)
	return nil
}
