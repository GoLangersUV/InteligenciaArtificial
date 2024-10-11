package game

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	TileSize = 64
	MaxSize  = 10
)

type Tile int

const (
	Road Tile = iota
	Wall
	Car
	MediumTraffic
	HighTraffic
	Passenger
	Goal
)

type Scene struct {
	Grid          [MaxSize][MaxSize]Tile
	Images        map[Tile]*ebiten.Image
	CarPosX       int
	CarPosY       int
	PassengerPosX int
	PassengerPosY int
	GoalPosX      int
	GoalPosY      int
}

func NewScene(matrix [][]int) *Scene {
	scene := &Scene{
		Images: make(map[Tile]*ebiten.Image),
	}

	// Load images
	imageFiles := map[Tile]string{
		Road:          "./game/assets/images/calle-uldr.png",
		Wall:          "./game/assets/images/muro-1.png",
		Car:           "./game/assets/images/moto-1-narvaez.png",
		MediumTraffic: "./game/assets/images/trafico-medio.png",
		HighTraffic:   "./game/assets/images/trafico-pesado.png",
		Passenger:     "./game/assets/images/girl.png",
		Goal:          "./game/assets/images/destino.png",
	}

	for tile, file := range imageFiles {
		img, _, err := ebitenutil.NewImageFromFile(file)
		if err != nil {
			log.Fatal(err)
		}
		scene.Images[tile] = img
	}

	// Populate grid
	for y, row := range matrix {
		for x, val := range row {
			scene.Grid[y][x] = Tile(val)
			if Tile(val) == Car {
				scene.Grid[y][x] = Tile(0)
				scene.CarPosX = x
				scene.CarPosY = y
			}
			if Tile(val) == Passenger {
				scene.Grid[y][x] = Tile(0)
				scene.PassengerPosX = x
				scene.PassengerPosY = y
			}
			if Tile(val) == Goal {
				scene.GoalPosX = x
				scene.GoalPosY = y
			}
		}
	}

	return scene
}

func (s *Scene) Draw(screen *ebiten.Image) {
	for y := 0; y < MaxSize; y++ {
		for x := 0; x < MaxSize; x++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*TileSize), float64(y*TileSize))
			screen.DrawImage(s.Images[s.Grid[y][x]], op)
		}
	}
}
