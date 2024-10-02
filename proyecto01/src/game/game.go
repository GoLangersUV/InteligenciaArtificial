package game

import (
	"github.com/Krud3/InteligenciaArtificial/src/utils"

	"github.com/hajimehoshi/ebiten/v2"
	// Import the entities package
)

type Game struct {
	scene     *Scene
	car       *Car
	passenger *Passenger // Optional if the passenger needs independent logic
}

func NewGame() (*Game, error) {
	matrix, err := utils.GetMatrix() // Assuming this loads the matrix
	if err != nil {
		return nil, err
	}

	// Create the scene
	scene := NewScene(matrix)

	// Create the car at its initial position using the scene's car start position
	carPath := [][]int{
		// Sample path or generate it using a search algorithm
	}
	car := NewCar(scene.CarPosX, scene.CarPosY, carPath)

	// Optionally create a passenger if it has independent logic
	passenger := NewPassenger(scene.PassengerPosX, scene.PassengerPosY)

	return &Game{
		scene:     scene,
		car:       car,
		passenger: passenger,
	}, nil
}

func (g *Game) Update() error {
	// Move the car along its path
	g.car.Move()

	// Check if the car reaches the passenger
	if g.car.PosX == g.scene.PassengerPosX && g.car.PosY == g.scene.PassengerPosY {
		// Remove the passenger or handle pickup logic
		g.passenger = nil // Passenger disappears after being picked up
	}

	// Check if the car reaches the goal
	if g.car.PosX == g.scene.GoalPosX && g.car.PosY == g.scene.GoalPosY {
		// Handle reaching the goal (e.g., end the game or display success)
	}

	return nil
}

func (s *Scene) Draw(screen *ebiten.Image) {
	// Draw the grid
	for y := 0; y < MaxSize; y++ {
		for x := 0; x < MaxSize; x++ {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x*TileSize), float64(y*TileSize))
			screen.DrawImage(s.Images[s.Grid[y][x]], op)
		}
	}

	// Draw the car at its current position
	carOp := &ebiten.DrawImageOptions{}
	carOp.GeoM.Translate(float64(s.Car.PosX*TileSize), float64(s.Car.PosY*TileSize))
	screen.DrawImage(s.Images[Car], carOp)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return MaxSize * TileSize, MaxSize * TileSize
}
