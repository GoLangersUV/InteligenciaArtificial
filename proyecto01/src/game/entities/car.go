package entities

type Car struct {
	PosX, PosY int     // Current position
	Path       [][]int // List of positions to move through
	Index      int     // Current index in the path
}

func NewCar(startX, startY int, path [][]int) *Car {
	return &Car{
		PosX:  startX,
		PosY:  startY,
		Path:  path,
		Index: 0, // Start from the beginning of the path
	}
}

func (c *Car) Move() {
	if c.Index < len(c.Path) {
		// Update the car’s position to the next position in the path
		c.PosX = c.Path[c.Index][0]
		c.PosY = c.Path[c.Index][1]
		c.Index++
	}
}
