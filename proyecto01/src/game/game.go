package game

import (
	"fmt"
	"image"
	"image/color"
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Krud3/InteligenciaArtificial/src/datatypes"
	"github.com/Krud3/InteligenciaArtificial/src/searchAlgorithms"
	"github.com/Krud3/InteligenciaArtificial/src/utils"

	"github.com/Krud3/InteligenciaArtificial/src/game/entities"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/ncruces/zenity"
)

type GameState int

type AlgorithmType int

type AreaOfKeyEvents int

type Game struct {
	state                  GameState
	scene                  *Scene
	car                    *entities.Car
	passenger              *entities.Passenger
	selectedFileIndex      int
	files                  []string
	frameCount             int
	algorithms             []string
	algorithmType          AlgorithmType
	selectedAlgorithmIndex int
	selectedBox            AreaOfKeyEvents
	nodesExpanded          int
	treeDepth              int
	computationTime        float64
	solutionCost           float64
	titleImage             *ebiten.Image
}

const (
	MenuState GameState = iota
	PlayingState
	EndState
)

const (
	InformedAlgorithm AlgorithmType = iota
	UninformedAlgorithm
)

const (
	LeftBox AreaOfKeyEvents = iota
	RightBox
)

const (
	verticalUploadPhase   int = 0
	horizontalUploadPhase int = -375
)

const (
	verticalSelectAlPhase   int = 0
	horizontalSelectAlPhase int = 380
)

var (
	informedAlgorithms   []string = []string{"dummyAlgorithm", "AStar"}
	uninformedAlgorithms []string = []string{"Breadth First Algorithm"}
)

var Matrix datatypes.ScannedMatrix

func NewGame(matrixFileName string) (*Game, error) {
  Matrix, _ = utils.GetMatrix(matrixFileName)

	// Create the scene
	scene := NewScene(Matrix.Matrix)

	car := entities.NewCar(scene.CarPosX, scene.CarPosY) // Create the car

	passenger := entities.NewPassenger(scene.PassengerPosX, scene.PassengerPosY) // Create the passenger

	titleImage, _, err := ebitenutil.NewImageFromFile("./game/assets/images/title.png")
	if err != nil {
		log.Fatal(err)
	}

	game := &Game{
		state:                  MenuState,
		scene:                  scene,
		car:                    car,
		passenger:              passenger,
		selectedFileIndex:      -1,
		algorithms:             informedAlgorithms,
		algorithmType:          InformedAlgorithm,
		selectedAlgorithmIndex: -1,
		selectedBox:            LeftBox,
		titleImage:             titleImage,
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
		g.DrawEndScreen(screen)
	}
}

func (g *Game) DrawMenu(screen *ebiten.Image) {
	// Draw the title image at the top-center of the screen
	screenWidth, _ := ebiten.WindowSize()
	titleImageWidth, _ := g.titleImage.Size()

	// Calculate the position to center the title image
	x := (screenWidth - titleImageWidth) / 2
	y := 50 // You can adjust this for vertical positioning

	// Draw the title image
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(g.titleImage, op)

	// Dibujar el botón de inicio
	startButtonRect := image.Rect(((MaxSize*TileSize)/2 - 300 + horizontalSelectAlPhase), 550+verticalSelectAlPhase, ((MaxSize*TileSize)/2 - 100 + horizontalSelectAlPhase), 650+verticalSelectAlPhase)
	ebitenutil.DrawRect(screen, float64(startButtonRect.Min.X), float64(startButtonRect.Min.Y), float64(startButtonRect.Dx()), float64(startButtonRect.Dy()), color.RGBA{0, 160, 0, 255})
	ebitenutil.DrawRect(screen, float64(startButtonRect.Min.X)+5, float64(startButtonRect.Min.Y)+5, float64(startButtonRect.Dx())-10, float64(startButtonRect.Dy())-10, color.RGBA{0, 140, 0, 255})
	ebitenutil.DebugPrintAt(screen, "Start Game", ((MaxSize*TileSize)/2-300)+65+horizontalSelectAlPhase, 592+verticalSelectAlPhase)

	// Dibujar el botón para subir la matriz
	uploadButtonRect := image.Rect((MaxSize*TileSize)/2+90+horizontalUploadPhase, (MaxSize*TileSize)/2-70+verticalUploadPhase, (MaxSize*TileSize)/2+310+horizontalUploadPhase, ((MaxSize*TileSize)/4*2)-20+verticalUploadPhase)
	ebitenutil.DrawRect(screen, float64(uploadButtonRect.Min.X), float64(uploadButtonRect.Min.Y), float64(uploadButtonRect.Dx()), float64(uploadButtonRect.Dy()), color.RGBA{0, 50, 255, 255})
	ebitenutil.DrawRect(screen, float64(uploadButtonRect.Min.X)+5, float64(uploadButtonRect.Min.Y)+5, float64(uploadButtonRect.Dx())-10, float64(uploadButtonRect.Dy())-10, color.RGBA{0, 35, 240, 255})
	ebitenutil.DebugPrintAt(screen, "Upload Matrix", ((MaxSize*TileSize)/2+90)+65+horizontalUploadPhase, ((MaxSize*TileSize)/4*2)-52+verticalUploadPhase)
	g.DrawFiles(screen)

	// Dibujar el botón de búsqueda informada
	informedSearchButtonRect := image.Rect((MaxSize*TileSize)/2-300+horizontalSelectAlPhase, 250+verticalSelectAlPhase, (MaxSize*TileSize)/2-100+horizontalSelectAlPhase, 300+verticalSelectAlPhase)
	ebitenutil.DrawRect(screen, float64(informedSearchButtonRect.Min.X), float64(informedSearchButtonRect.Min.Y), float64(informedSearchButtonRect.Dx()), float64(informedSearchButtonRect.Dy()), color.RGBA{235, 130, 0, 255})
	ebitenutil.DrawRect(screen, float64(informedSearchButtonRect.Min.X)+5, float64(informedSearchButtonRect.Min.Y)+5, float64(informedSearchButtonRect.Dx())-10, float64(informedSearchButtonRect.Dy())-10, color.RGBA{220, 115, 0, 255})
	ebitenutil.DebugPrintAt(screen, "Informed Search", ((MaxSize*TileSize)/2-300)+50+horizontalSelectAlPhase, 267+verticalSelectAlPhase)

	// Dibujar el botón de búsqueda no informada
	unInformedSearchButtonRect := image.Rect(((MaxSize*TileSize)/2 - 300 + horizontalSelectAlPhase), 300+verticalSelectAlPhase, ((MaxSize*TileSize)/2 - 100 + horizontalSelectAlPhase), 350+verticalSelectAlPhase)
	ebitenutil.DrawRect(screen, float64(unInformedSearchButtonRect.Min.X), float64(unInformedSearchButtonRect.Min.Y), float64(unInformedSearchButtonRect.Dx()), float64(unInformedSearchButtonRect.Dy()), color.RGBA{235, 130, 0, 255})
	ebitenutil.DrawRect(screen, float64(unInformedSearchButtonRect.Min.X)+5, float64(unInformedSearchButtonRect.Min.Y)+5, float64(unInformedSearchButtonRect.Dx())-10, float64(unInformedSearchButtonRect.Dy())-10, color.RGBA{220, 115, 0, 255})
	ebitenutil.DebugPrintAt(screen, "Uninformed Search", ((MaxSize*TileSize)/2-300)+45+horizontalSelectAlPhase, 317+verticalSelectAlPhase)

	g.DrawAlgorithms(screen)
}

func (g *Game) DrawFiles(screen *ebiten.Image) {
	// Dibujar los archivos disponibles en la carpeta 'battery'
	y := ((MaxSize * TileSize) / 2) + verticalUploadPhase
	ebitenutil.DrawRect(screen, float64((MaxSize*TileSize)/2+100-10+horizontalUploadPhase), float64(y-20), 220, 20, color.RGBA{100, 100, 100, 255})
	for i, file := range g.files {
		text := file
		if g.selectedFileIndex == i {
			text = "> " + file // Añadir un indicador para mostrar el archivo seleccionado
		}
		ebitenutil.DrawRect(screen, float64((MaxSize*TileSize)/2+100-10+horizontalUploadPhase), float64(y), 220, 20, color.RGBA{100, 100, 100, 255})
		ebitenutil.DebugPrintAt(screen, text, ((MaxSize*TileSize)/2 + 100 + horizontalUploadPhase), y)
		y += 20
	}
	ebitenutil.DrawRect(screen, float64((MaxSize*TileSize)/2+100-10+horizontalUploadPhase), float64(y), 220, 20, color.RGBA{100, 100, 100, 255})
}

func (g *Game) DrawAlgorithms(screen *ebiten.Image) {
	// Dibujar los archivos disponibles en la carpeta 'battery'
	y := 350 + verticalSelectAlPhase + 20
	ebitenutil.DrawRect(screen, float64((MaxSize*TileSize)/2-300+horizontalSelectAlPhase), float64(y-20), 200, 20*6, color.RGBA{100, 100, 100, 255})
	for i, algorithm := range g.algorithms {
		text := algorithm
		if g.selectedAlgorithmIndex == i {
			text = "> " + algorithm // Añadir un indicador para mostrar el archivo seleccionado
		}
		ebitenutil.DebugPrintAt(screen, text, ((MaxSize*TileSize)/2 - 300 + 10 + horizontalSelectAlPhase), y)
		y += 20
	}
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

func (g *Game) DrawEndScreen(screen *ebiten.Image) {
	ebitenutil.DebugPrintAt(screen, "Algorithm Execution Report", 50, 50)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Nodes Expanded: %d", g.nodesExpanded), 50, 80)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Tree Depth: %d", g.treeDepth), 50, 110)
	ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Computation Time: %.2f seconds", g.computationTime), 50, 140)

	// Only display the solution cost if applicable
	if g.solutionCost > 0 {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Solution Cost: %.2f", g.solutionCost), 50, 170)
	}

	// Add a button to return to the menu
	backButtonRect := image.Rect(50, 550, 200, 600)
	ebitenutil.DrawRect(screen, float64(backButtonRect.Min.X), float64(backButtonRect.Min.Y), float64(backButtonRect.Dx()), float64(backButtonRect.Dy()), color.RGBA{0, 160, 0, 255})
	ebitenutil.DebugPrintAt(screen, "Return to Menu", 70, 570)

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if x >= 50 && x <= 200 && y >= 550 && y <= 600 {
			g.state = MenuState // Return to the menu
		}
	}
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
			if g.selectedFileIndex < len(g.files)-1 && g.selectedBox == LeftBox {
				g.selectedFileIndex++
			}
			if g.selectedAlgorithmIndex < len(g.algorithms)-1 && g.selectedBox == RightBox {
				g.selectedAlgorithmIndex++
			}
		}

		// Navigate up the file list
		if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
			if g.selectedFileIndex == -1 && len(g.files) > 0 && g.selectedBox == LeftBox {
				g.selectedFileIndex = 1
			}
			if g.selectedFileIndex > 0 && g.selectedBox == LeftBox {
				g.selectedFileIndex--
			}

			if g.selectedAlgorithmIndex == -1 && len(g.algorithms) > 0 && g.selectedBox == RightBox {
				g.selectedAlgorithmIndex = 1
			}
			if g.selectedAlgorithmIndex > 0 && g.selectedBox == RightBox {
				g.selectedAlgorithmIndex--
			}
		}

		if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
			if g.selectedFileIndex == -1 && len(g.files) > 0 {
				g.selectedFileIndex = 0
			}
			if g.selectedBox != LeftBox {
				g.selectedBox = LeftBox
			}
			print("LeftBox")

		}

		if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
			if g.selectedAlgorithmIndex == -1 && len(g.algorithms) > 0 {
				g.selectedAlgorithmIndex = 0
			}
			if g.selectedBox != RightBox {
				g.selectedBox = RightBox
			}
			print("RightBox")

		}
	}

	// Confirmar la selección del archivo con Enter
	if ebiten.IsKeyPressed(ebiten.KeyEnter) && g.selectedFileIndex >= 0 {
		selectedFile := g.files[g.selectedFileIndex]
		g.SetScene("../battery/" + selectedFile) // Cargar la nueva escena con el archivo seleccionado
		g.state = PlayingState                   // Cambiar el estado al de juego
		g.SetCarPath(g.algorithms[g.selectedAlgorithmIndex])
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()

		// Check if any file was clicked
		fileStartY := ((MaxSize * TileSize) / 4 * 2) + verticalUploadPhase // Starting Y position of the first file in the list
		fileHeight := 20                                                   // Height of each file item
		for i, file := range g.files {
			fileY := fileStartY + i*fileHeight
			if x >= (MaxSize*TileSize)/2+100+horizontalUploadPhase && x <= (MaxSize*TileSize)/2+300+horizontalUploadPhase && y >= fileY && y <= fileY+fileHeight {
				g.selectedFileIndex = i // Set the selected file based on the click
				print("File selected:", file)
			}
		}

		// Check if any algorithm was clicked
		algorithmStartY := 350 + verticalSelectAlPhase + 20 // Starting Y position of the first file in the list
		algorithmHeight := 20                               // Height of each file item
		for i, algorithm := range g.algorithms {
			algorithmY := algorithmStartY + i*algorithmHeight
			if x >= ((MaxSize*TileSize)/2-300+10+horizontalSelectAlPhase) && x <= ((MaxSize*TileSize)/2-100+10+horizontalSelectAlPhase) && y >= algorithmY && y <= algorithmY+algorithmHeight {
				g.selectedAlgorithmIndex = i // Set the selected file based on the click
				print("File selected:", algorithm)
			}
		}

		// Verificar si se presionó el botón "Start Game"
		if x >= ((MaxSize*TileSize)/2-300+horizontalSelectAlPhase) && x <= ((MaxSize*TileSize)/2-100+horizontalSelectAlPhase) && y >= 550+verticalSelectAlPhase && y <= 650+verticalSelectAlPhase {
			if g.selectedFileIndex >= 0 {
				selectedFile := g.files[g.selectedFileIndex]
				g.SetScene("../battery/" + selectedFile)
				g.state = PlayingState
				g.SetCarPath(g.algorithms[g.selectedAlgorithmIndex])
				print("Start Game")
			}
		}

		// Verificar si se presionó el botón "Upload Matrix"
		if x >= (MaxSize*TileSize)/2+90+horizontalUploadPhase && x <= (MaxSize*TileSize)/2+310+horizontalUploadPhase && y >= ((MaxSize*TileSize)/4*2)-70+verticalUploadPhase && y <= ((MaxSize*TileSize)/4*2)-20+verticalUploadPhase {
			g.UploadMatrix() // Llamar al método para subir la matriz
		}

		// Verificar si se presionó el botón "Informed Search"
		if x >= (MaxSize*TileSize)/2-300+horizontalSelectAlPhase && x <= (MaxSize*TileSize)/2-100+horizontalSelectAlPhase && y >= 250+verticalSelectAlPhase && y <= 300+verticalSelectAlPhase {
			g.algorithms = informedAlgorithms
		}

		// Verificar si se presionó el botón "Unnformed Search"
		if x >= (MaxSize*TileSize)/2-300+horizontalSelectAlPhase && x <= (MaxSize*TileSize)/2-100+horizontalSelectAlPhase && y >= 300+verticalSelectAlPhase && y <= 350+verticalSelectAlPhase {
			g.algorithms = uninformedAlgorithms
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
			g.state = EndState
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
	files, err := os.ReadDir(batteryDir)
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
	startTime := time.Now()
	switch algorithmKey {
	case "Breadth First Algorithm":
		result := searchAlgorithms.StartSearch(1, Matrix) // Call the dummy algorithm
		var mappedCoordinates [][]int
		for _, coord := range result.PathFound {
			mappedCoordinates = append(mappedCoordinates, []int{coord.X, coord.Y})
		}
		newPath = mappedCoordinates
		g.nodesExpanded = result.ExpandenNodes
		g.treeDepth = result.TreeDepth
		g.solutionCost = 3
  case "AStar":

    envMatrix,_ := searchAlgorithms.ValidateMatrix(Matrix.Matrix)
    env, err := searchAlgorithms.NewEnvironment(envMatrix)
    if err != nil {
        log.Fatalf("Error creating environment: %v", err)
    }
    aStarSearch := new(searchAlgorithms.AStarSearch)
    agent := searchAlgorithms.NewAgent(env.InitPosition, aStarSearch)

    result := agent.SearchAlgorithm.LookForGoal(env)
    newPath = searchAlgorithms.FromPosToPath(result.Path)
    g.nodesExpanded = result.ExpandedNodes
    g.treeDepth = result.TreeDepth
    g.solutionCost = float64(result.Cost)

	default:
		newPath = [][]int{} // Initialize with an empty slice for other cases
	}
	g.computationTime = time.Since(startTime).Seconds()

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

	// Create the scene
	g.scene = NewScene(Matrix.Matrix)

	scene := NewScene(Matrix.Matrix)

	g.car = entities.NewCar(scene.CarPosX, scene.CarPosY)

	g.passenger = entities.NewPassenger(scene.PassengerPosX, scene.PassengerPosY)
}
