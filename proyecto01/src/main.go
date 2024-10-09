package main

import (
	"log"

	"github.com/Krud3/InteligenciaArtificial/src/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	carPath := [][]int{
		{2, 0},
		{3, 0},
		{4, 0},
		{5, 0},
		{6, 0},
		{5, 0},
		{4, 0},
		{3, 0},
		{3, 1},
		{3, 2},
		{3, 3},
		{2, 3},
		{1, 3},
		{1, 4},
		{1, 5},
		{2, 5},
		{3, 5},
		{3, 6},
		{3, 7},
		{2, 7},
		{1, 7},
		{1, 8},
		{1, 9},
		{2, 9},
		{3, 9},
		{4, 9},
		{5, 9},
	}
	game, err := game.NewGame(carPath)
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(640, 640)
	ebiten.SetWindowTitle("Your Game Title")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
