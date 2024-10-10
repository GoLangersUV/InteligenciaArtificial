package entities

import "github.com/hajimehoshi/ebiten/v2"

type Passenger struct {
	PosX, PosY               int
	InitialPosX, InitialPosY int
	Image                    *ebiten.Image
}

func NewPassenger(x, y int) *Passenger {
	return &Passenger{
		PosX: x,
		PosY: y,
	}
}
