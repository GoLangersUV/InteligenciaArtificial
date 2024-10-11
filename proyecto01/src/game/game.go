package game

import (
	"image"
	"image/color"
	"log"

	"github.com/Krud3/InteligenciaArtificial/src/searchAlgorithms"
	"github.com/Krud3/InteligenciaArtificial/src/utils"

	"github.com/Krud3/InteligenciaArtificial/src/game/entities"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/ncruces/zenity"
)

type GameState int

type Game struct {
	state     GameState
	scene     *Scene
	car       *entities.Car
	passenger *entities.Passenger // Optional if the passenger needs independent logic
}

const (
	MenuState GameState = iota
	PlayingState
	EndState
)

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
		state:     MenuState,
		scene:     scene,
		car:       car,
		passenger: passenger,
	}, nil
}

func (g *Game) Update() error {

	switch g.state {
	case MenuState:
		g.UpdateMenu()
	case PlayingState:
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
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case MenuState:
		g.DrawMenu(screen)
	case PlayingState:
		// Draw the game
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
	case EndState:
		// Draw the end screen
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return MaxSize * TileSize, MaxSize * TileSize
}

func (g *Game) SetCarPath(algorithmKey string) {
	var newPath [][]int // Declare newPath outside the conditional block

	switch algorithmKey {
	case "callDummy":
		newPath = searchAlgorithms.DummyAlgorithm() // Call the dummy algorithm
	default:
		newPath = [][]int{} // Initialize with an empty slice for other cases
	}

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

func (g *Game) DrawMenu(screen *ebiten.Image) {
	// Draw the start button
	startButtonRect := image.Rect(100, 100, 300, 150)
	ebitenutil.DrawRect(screen, float64(startButtonRect.Min.X), float64(startButtonRect.Min.Y), float64(startButtonRect.Dx()), float64(startButtonRect.Dy()), color.RGBA{0, 255, 0, 255})

	// Add text to the button
	ebitenutil.DebugPrintAt(screen, "Start Game", 130, 120)

	// Add algorithm selection (as simple buttons or a dropdown)
	// Example: Draw buttons for different algorithms
	ebitenutil.DebugPrintAt(screen, "Algorithm: Dummy", 100, 200)
	// ... Add more options
}

func (g *Game) UpdateMenu() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		// Check if the start button was clicked
		if x >= 100 && x <= 300 && y >= 100 && y <= 150 {
			g.state = PlayingState
			g.SetCarPath("callDummy")
		}
	}
}

func (g *Game) UploadMatrix() {
	fileName, err := zenity.SelectFile()
	if err != nil {
		log.Println("No file selected or error:", err)
		return
	}
	g.SetScene(fileName) // Set new scene from the selected file
}
