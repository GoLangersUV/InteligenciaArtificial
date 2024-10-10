package game

import (
	"log"

	"github.com/Krud3/InteligenciaArtificial/src/utils"

	"github.com/Krud3/InteligenciaArtificial/src/game/entities"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	scene     *Scene
	car       *entities.Car
	passenger *entities.Passenger // Optional if the passenger needs independent logic
}

func NewGame(matrixFileName string) (*Game, error) {
	matrix, err := utils.GetMatrix(matrixFileName) // Load the matrix
	if err != nil {
		return nil, err
	}

	// Create the scene
	scene := NewScene(matrix)

	// Load the car image
	carImage, _, err := ebitenutil.NewImageFromFile("./game/assets/images/moto-1-narvaez.png")
	if err != nil {
		return nil, err
	}

	// Load the passenger image
	passengerImage, _, err := ebitenutil.NewImageFromFile("./game/assets/images/moto-2-narvaez.png")
	if err != nil {
		return nil, err
	}

	carPath := [][]int{}
	car := &entities.Car{
		PosX:        scene.CarPosX,
		PosY:        scene.CarPosY,
		InitialPosX: scene.CarPosX,
		InitialPosY: scene.CarPosY,
		Path:        carPath,
		Index:       0,
		Image:       carImage,
		Delay:       30,
	}

	// Create the passenger
	passenger := &entities.Passenger{
		PosX:  scene.PassengerPosX,
		PosY:  scene.PassengerPosY,
		Image: passengerImage, // Assign the passenger image
	}

	return &Game{
		scene:     scene,
		car:       car,
		passenger: passenger,
	}, nil
}

func (g *Game) Update() error {
	// Move the car along its path
	g.car.Update()

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

func (g *Game) Draw(screen *ebiten.Image) {
	// Draw the scene first
	g.scene.Draw(screen)

	// Draw the car on top of the scene
	carOp := &ebiten.DrawImageOptions{}
	carOp.GeoM.Translate(float64(g.car.PosX*TileSize), float64(g.car.PosY*TileSize))
	screen.DrawImage(g.scene.Images[Car], carOp)

	// Optionally draw the passenger if it's still present
	if g.passenger != nil {
		passengerOp := &ebiten.DrawImageOptions{}
		passengerOp.GeoM.Translate(float64(g.passenger.PosX*TileSize), float64(g.passenger.PosY*TileSize))
		screen.DrawImage(g.scene.Images[Passenger], passengerOp)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return MaxSize * TileSize, MaxSize * TileSize
}

func (g *Game) SetCarPath(newPath [][]int) {
	g.car.Path = newPath // Set the new path
	if g.car.PosX != g.car.InitialPosX && g.car.PosY != g.car.InitialPosY {
		g.car.Index = 0                // Reset the index so the car follows the new path from the beginning
		g.car.PosX = g.car.InitialPosX // Reset the car's position to the initial X
		g.car.PosY = g.car.InitialPosY // Reset the car's position to the initial Y
	}

	// Check if the passenger is nil (i.e., removed)
	if g.passenger == nil {
		// Recreate the passenger
		// Load the passenger image
		passengerImage, _, err := ebitenutil.NewImageFromFile("./game/assets/images/moto-2-narvaez.png")
		if err != nil {
			log.Fatal(err) // Handle the error appropriately
		}

		// Create a new passenger and reset it to the initial position
		g.passenger = &entities.Passenger{
			PosX:  g.scene.PassengerPosX,
			PosY:  g.scene.PassengerPosY,
			Image: passengerImage, // Assign the passenger image
		}
	}
}

func (g *Game) SetScene(fileName string) {
	matrix, err := utils.GetMatrix(fileName) // Load the matrix
	if err != nil {
		log.Fatal(err) // Handle the error appropriately
	}

	// Create the scene
	g.scene = NewScene(matrix)

	// Load the car image
	carImage, _, err := ebitenutil.NewImageFromFile("./game/assets/images/moto-1-narvaez.png")
	if err != nil {
		log.Fatal(err)
	}

	// Load the passenger image
	passengerImage, _, err := ebitenutil.NewImageFromFile("./game/assets/images/moto-2-narvaez.png")
	if err != nil {
		log.Fatal(err)
	}

	carPath := [][]int{}
	g.car = &entities.Car{
		PosX:        g.scene.CarPosX,
		PosY:        g.scene.CarPosY,
		InitialPosX: g.scene.CarPosX,
		InitialPosY: g.scene.CarPosY,
		Path:        carPath,
		Index:       0,
		Image:       carImage,
		Delay:       30,
	}

	// Create the passenger
	g.passenger = &entities.Passenger{
		PosX:  g.scene.PassengerPosX,
		PosY:  g.scene.PassengerPosY,
		Image: passengerImage, // Assign the passenger image
	}
}
