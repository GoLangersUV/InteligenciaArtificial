package game

import (
	"image"
	"image/color"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/Krud3/InteligenciaArtificial/src/searchAlgorithms"
	"github.com/Krud3/InteligenciaArtificial/src/utils"

	"github.com/Krud3/InteligenciaArtificial/src/game/entities"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/ncruces/zenity"
)

type GameState int

type Game struct {
	state             GameState
	scene             *Scene
	car               *entities.Car
	passenger         *entities.Passenger // Optional if the passenger needs independent logic
	selectedFileIndex int
	files             []string
	frameCount        int
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

	car := entities.NewCar(scene.CarPosX, scene.CarPosY) // Create the car

	passenger := entities.NewPassenger(scene.PassengerPosX, scene.PassengerPosY) // Create the passenger

	game := &Game{
		state:             MenuState,
		scene:             scene,
		car:               car,
		passenger:         passenger,
		selectedFileIndex: -1,
	}

	game.files = game.ListMatrixFiles() // List the files in the 'battery' folder

	return game, nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return MaxSize * TileSize, (MaxSize * TileSize) + TileSize
}

func (g *Game) Draw(screen *ebiten.Image) {
	switch g.state {
	case MenuState:
		g.DrawMenu(screen)
	case PlayingState:
		g.DrawGame(screen)
	case EndState:
		// Draw the end screen
	}
}

func (g *Game) DrawMenu(screen *ebiten.Image) {
	// Dibujar el botón de inicio
	startButtonRect := image.Rect(((MaxSize*TileSize)/2 - 300), 550, ((MaxSize*TileSize)/2 - 100), 650)
	ebitenutil.DrawRect(screen, float64(startButtonRect.Min.X), float64(startButtonRect.Min.Y), float64(startButtonRect.Dx()), float64(startButtonRect.Dy()), color.RGBA{0, 150, 0, 255})
	ebitenutil.DebugPrintAt(screen, "Start Game", ((MaxSize*TileSize)/2-300)+65, 592)

	// Dibujar el botón para subir la matriz
	verticalPhase := 80
	uploadButtonRect := image.Rect(((MaxSize*TileSize)/2 + 90), ((MaxSize*TileSize)/4*2)-70+verticalPhase, ((MaxSize*TileSize)/2 + 310), ((MaxSize*TileSize)/4*2)-20+verticalPhase)
	ebitenutil.DrawRect(screen, float64(uploadButtonRect.Min.X), float64(uploadButtonRect.Min.Y), float64(uploadButtonRect.Dx()), float64(uploadButtonRect.Dy()), color.RGBA{0, 0, 255, 255})
	ebitenutil.DebugPrintAt(screen, "Upload Matrix", ((MaxSize*TileSize)/2+90)+40, ((MaxSize*TileSize)/4*2)-50+verticalPhase)
	g.DrawFiles(screen, verticalPhase)

	// Dibujar el botón de inicio
	informedSearchButtonRect := image.Rect(((MaxSize*TileSize)/2 - 300), 190, ((MaxSize*TileSize)/2 - 100), 240)
	ebitenutil.DrawRect(screen, float64(informedSearchButtonRect.Min.X), float64(informedSearchButtonRect.Min.Y), float64(informedSearchButtonRect.Dx()), float64(informedSearchButtonRect.Dy()), color.RGBA{235, 130, 0, 255})
	ebitenutil.DebugPrintAt(screen, "Informed Search", ((MaxSize*TileSize)/2-300)+45, 210)

	// Dibujar el botón de inicio
	unInformedSearchButtonRect := image.Rect(((MaxSize*TileSize)/2 - 300), 250, ((MaxSize*TileSize)/2 - 100), 300)
	ebitenutil.DrawRect(screen, float64(unInformedSearchButtonRect.Min.X), float64(unInformedSearchButtonRect.Min.Y), float64(unInformedSearchButtonRect.Dx()), float64(unInformedSearchButtonRect.Dy()), color.RGBA{235, 130, 0, 255})
	ebitenutil.DebugPrintAt(screen, "Uninformed Search", ((MaxSize*TileSize)/2-300)+45, 270)
}

func (g *Game) DrawFiles(screen *ebiten.Image, verticalPhase int) {
	// Dibujar los archivos disponibles en la carpeta 'battery'
	y := ((MaxSize * TileSize) / 4 * 2) + verticalPhase
	ebitenutil.DrawRect(screen, float64((MaxSize*TileSize)/2+100)-10, float64(y-20), 220, 20, color.RGBA{100, 100, 100, 255})
	for i, file := range g.files {
		text := file
		if g.selectedFileIndex == i {
			text = "> " + file // Añadir un indicador para mostrar el archivo seleccionado
		}
		ebitenutil.DrawRect(screen, float64((MaxSize*TileSize)/2+100)-10, float64(y), 220, 20, color.RGBA{100, 100, 100, 255})
		ebitenutil.DebugPrintAt(screen, text, ((MaxSize*TileSize)/2 + 100), y)
		y += 20
	}
	ebitenutil.DrawRect(screen, float64((MaxSize*TileSize)/2+100)-10, float64(y), 220, 20, color.RGBA{100, 100, 100, 255})
}

func (g *Game) DrawGame(screen *ebiten.Image) {
	g.scene.Draw(screen)

	// Draw the car on top of the scene
	carOp := &ebiten.DrawImageOptions{}
	carOp.GeoM.Translate(float64(g.car.PosX*TileSize), float64(g.car.PosY*TileSize))
	screen.DrawImage(g.car.Image, carOp)

	// Optionally draw the passenger if it's still present
	if g.passenger != nil {
		passengerOp := &ebiten.DrawImageOptions{}
		passengerOp.GeoM.Translate(float64(g.passenger.PosX*TileSize), float64(g.passenger.PosY*TileSize))
		screen.DrawImage(g.passenger.Image, passengerOp)
	}

	// Render the "Back to Menu" button in the upper-right corner
	backButtonWidth := 120
	backButtonHeight := 40
	screenWidth, _ := ebiten.WindowSize()
	// Coordinates for the back button
	backButtonX := screenWidth - backButtonWidth
	backButtonY := MaxSize*TileSize + 10
	// Draw the button (a rectangle)
	ebitenutil.DrawRect(screen, float64(backButtonX), float64(backButtonY), float64(backButtonWidth), float64(backButtonHeight), color.RGBA{255, 0, 0, 255})
	// Add text to the button
	ebitenutil.DebugPrintAt(screen, "Back to Menu", backButtonX+20, backButtonY+10)

	// Render the "Back to Menu" button in the upper-right corner
	statsButtonWidth := 120
	statsButtonHeight := 40
	// Coordinates for the back button
	statsButtonX := 0
	statsButtonY := MaxSize*TileSize + 10
	// Draw the button (a rectangle)
	ebitenutil.DrawRect(screen, float64(statsButtonX), float64(statsButtonY), float64(statsButtonWidth), float64(statsButtonHeight), color.RGBA{125, 125, 125, 255})
	// Add text to the button
	ebitenutil.DebugPrintAt(screen, "Game Stats", statsButtonX+20, statsButtonY+10)
}

func (g *Game) Update() error {

	switch g.state {
	case MenuState:
		g.UpdateMenu()
	case PlayingState:
		g.UpdateGame()
	}
	return nil
}

func (g *Game) UpdateMenu() {
	const keyPressDelay = 8
	g.frameCount++

	if g.frameCount >= keyPressDelay {
		g.frameCount = 0

		// Navigate down the file list
		if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
			if g.selectedFileIndex < len(g.files)-1 {
				g.selectedFileIndex++
			}
		}

		// Navigate up the file list
		if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
			if g.selectedFileIndex > 0 {
				g.selectedFileIndex--
			}
		}

		// Confirm the file selection with Enter
		if ebiten.IsKeyPressed(ebiten.KeyEnter) && g.selectedFileIndex >= 0 {
			selectedFile := g.files[g.selectedFileIndex]
			g.SetScene("../battery/" + selectedFile) // Load the new scene with the selected file
			g.state = PlayingState                   // Change the state to playing
			g.SetCarPath("callDummy")                // Set the car's path
		}
	}

	// Confirmar la selección del archivo con Enter
	if ebiten.IsKeyPressed(ebiten.KeyEnter) && g.selectedFileIndex >= 0 {
		selectedFile := g.files[g.selectedFileIndex]
		g.SetScene("../battery/" + selectedFile) // Cargar la nueva escena con el archivo seleccionado
		g.state = PlayingState                   // Cambiar el estado al de juego
		g.SetCarPath("callDummy")                // Establecer la ruta del coche)
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		// Check if any file was clicked
		fileStartY := ((MaxSize * TileSize) / 4 * 2) // Starting Y position of the first file in the list
		fileHeight := 20                             // Height of each file item
		for i, file := range g.files {
			fileY := fileStartY + i*fileHeight
			if x >= ((MaxSize*TileSize)/2-100) && x <= ((MaxSize*TileSize)/2+100) && y >= fileY && y <= fileY+fileHeight {
				g.selectedFileIndex = i // Set the selected file based on the click
				print("File selected:", file)
			}
		}

		// Verificar si se presionó el botón "Start Game"
		if x >= ((MaxSize*TileSize)/2-100) && x <= ((MaxSize*TileSize)/2+100) && y >= 100 && y <= 150 {
			if g.selectedFileIndex >= 0 {
				selectedFile := g.files[g.selectedFileIndex]
				g.SetScene("../battery/" + selectedFile)
				g.state = PlayingState
				g.SetCarPath("callDummy")
				print("Start Game")
			}
		}

		// Verificar si se presionó el botón "Upload Matrix"
		if x >= ((MaxSize*TileSize)/2-100) && x <= ((MaxSize*TileSize)/2+100) && y >= 200 && y <= 250 {
			g.UploadMatrix() // Llamar al método para subir la matriz
		}
	}
}

func (g *Game) UpdateGame() {
	// Move the car along its path
	g.car.Update()

	// Check if the car reaches the passenger
	if g.car.PosX == g.scene.PassengerPosX && g.car.PosY == g.scene.PassengerPosY {
		// Remove the passenger or handle pickup logic
		g.passenger = nil // Passenger disappears after being picked up
		g.car.SetImageWithPassenger()
	}

	// Check if the car reaches the goal
	if g.car.PosX == g.scene.GoalPosX && g.car.PosY == g.scene.GoalPosY {
		// Handle reaching the goal (e.g., end the game or display success)
	}

	// Handle the mouse click for the "Back to Menu" button
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		backButtonWidth := 120
		backButtonHeight := 40
		screenWidth, _ := ebiten.WindowSize()

		// Coordinates for the back button
		backButtonX := screenWidth - backButtonWidth
		backButtonY := MaxSize*TileSize + 10

		// Check if the mouse click is within the button's area
		if x >= backButtonX && x <= backButtonX+backButtonWidth && y >= backButtonY && y <= backButtonY+backButtonHeight {
			g.state = MenuState // Go back to the menu state
			print("Going back to the menu...")
		}

		// Handle the mouse click for the "Game Stats" button
		statsButtonWidth := 120
		statsButtonHeight := 40
		statsButtonX := 0
		statsButtonY := MaxSize*TileSize + 10

		// Check if the mouse click is within the button's area
		if x >= statsButtonX && x <= statsButtonX+statsButtonWidth && y >= statsButtonY && y <= statsButtonY+statsButtonHeight {
			// Handle the click for the "Game Stats" button
			// g.state = EndState
			print("Game Stats")
		}
	}
}

func (g *Game) UploadMatrix() {
	fileName, err := zenity.SelectFile(
		zenity.Title("Select a Matrix File"),
		zenity.FileFilter{
			Name:     "Text Files",
			Patterns: []string{"*.txt"},
		},
	)
	if err != nil {
		if err == zenity.ErrCanceled {
			log.Println("No file selected or operation was canceled")
			return
		}
		log.Printf("Error selecting file: %v", err)
		return
	}

	targetDir := "../battery"
	if _, err := os.Stat(targetDir); os.IsNotExist(err) {
		err := os.MkdirAll(targetDir, os.ModePerm)
		if err != nil {
			log.Fatalf("Error creating battery directory: %v", err)
		}
	}

	targetPath := filepath.Join(targetDir, filepath.Base(fileName))

	// Copiar el archivo seleccionado a la carpeta 'battery'
	err = copyFile(fileName, targetPath)
	if err != nil {
		log.Fatalf("Error copying file: %v", err)
	}

	log.Printf("File copied to %s successfully.", targetPath)

	// Establecer la nueva escena utilizando el archivo copiado
	g.SetScene(targetPath)
	g.files = g.ListMatrixFiles()
}

// copyFile copia un archivo de origen a un destino
func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

func (g *Game) ListMatrixFiles() []string {
	batteryDir := "../battery"
	files, err := ioutil.ReadDir(batteryDir)
	if err != nil {
		log.Fatalf("Error reading battery directory: %v", err)
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileNames = append(fileNames, file.Name())
		}
	}

	return fileNames
}

func (g *Game) SetCarPath(algorithmKey string) {
	var newPath [][]int // Declare newPath outside the conditional block

	switch algorithmKey {
	case "callDummy":
		newPath = searchAlgorithms.DummyAlgorithm() // Call the dummy algorithm
	default:
		newPath = [][]int{} // Initialize with an empty slice for other cases
	}

	if g.car.PosX != g.car.InitialPosX && g.car.PosY != g.car.InitialPosY {
		g.car.Reset() // Reset the car position if it's not at the initial position
	}
	g.car.SetPath(newPath)

	// Check if the passenger is nil (i.e., removed)
	if g.passenger == nil {
		// Create a new passenger and reset it to the initial position
		g.passenger = entities.NewPassenger(g.scene.PassengerPosX, g.scene.PassengerPosY)
	}
}

func (g *Game) SetScene(fileName string) {
	matrix, err := utils.GetMatrix(fileName) // Load the matrix
	if err != nil {
		log.Fatal(err) // Handle the error appropriately
	}

	// Create the scene
	g.scene = NewScene(matrix)

	scene := NewScene(matrix)

	g.car = entities.NewCar(scene.CarPosX, scene.CarPosY)

	g.passenger = entities.NewPassenger(scene.PassengerPosX, scene.PassengerPosY)
}
