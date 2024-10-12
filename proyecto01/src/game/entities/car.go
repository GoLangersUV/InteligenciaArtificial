package entities

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Car struct {
	PosX, PosY               int
	InitialPosX, InitialPosY int
	Path                     [][]int
	Index                    int
	Image                    *ebiten.Image
	frameCount               int
	Delay                    int
}

func (c *Car) Update() {
	if c.Index >= len(c.Path) {
		return // No more movement if the car reached the end of the path
	}

	c.frameCount++

	if c.frameCount >= c.Delay {
		// Move the car to the next position in the path
		nextPos := c.Path[c.Index]
		c.PosX = nextPos[1]
		c.PosY = nextPos[0]

		c.Index++        // Move to the next path index
		c.frameCount = 0 // Reset the frame counter
	}
}

func (c *Car) Reset() {
	c.PosX = c.InitialPosX
	c.PosY = c.InitialPosY
	c.Index = 0
}

func (c *Car) SetPath(path [][]int) {
	c.Path = path
}

func NewCar(x, y int) *Car {
	// Load the car image
	carImage, _, err := ebitenutil.NewImageFromFile("./game/assets/images/moto-1-narvaez.png")
	if err != nil {
		log.Fatal(err)
	}
	return &Car{
		PosX:        x,
		PosY:        y,
		InitialPosX: x,
		InitialPosY: y,
		Path:        [][]int{},
		Index:       0,
		Image:       carImage,
		Delay:       30,
	}
}

func (c *Car) SetImageWithPassenger() {
	carImage, _, err := ebitenutil.NewImageFromFile("./game/assets/images/moto-1-narvaez-girl.png")
	if err != nil {
		log.Fatal(err)
	}
	c.Image = carImage
}
