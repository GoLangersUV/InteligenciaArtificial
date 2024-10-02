package game

import (
	"github.com/Krud3/InteligenciaArtificial/src/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	scene *Scene
}

func NewGame() (*Game, error) {
	matrix, err := utils.GetMatrix()
	if err != nil {
		return nil, err
	}

	return &Game{
		scene: NewScene(matrix),
	}, nil
}

func (g *Game) Update() error {
	// Implement game update logic here
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.scene.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return MaxSize * TileSize, MaxSize * TileSize
}
