package entities

import "github.com/hajimehoshi/ebiten/v2"

type Car struct {
	PosX, PosY               int
	InitialPosX, InitialPosY int
	Path                     [][]int
	Index                    int
	Image                    *ebiten.Image
	frameCount               int
	Delay                    int
}

// func NewCar(startX, startY int, path [][]int) *Car {
// 	return &Car{
// 		PosX:  startX,
// 		PosY:  startY,
// 		Path:  path,
// 		Index: 0, // Start from the beginning of the path
// 	}
// }

// func (c *Car) Move() {
// 	if c.Index < len(c.Path) {
// 		// Update the car’s position to the next position in the path
// 		c.PosX = c.Path[c.Index][0]
// 		c.PosY = c.Path[c.Index][1]
// 		c.Index++
// 	}
// }

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
