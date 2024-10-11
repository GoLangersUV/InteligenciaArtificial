package entities

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Passenger struct {
	PosX, PosY               int
	InitialPosX, InitialPosY int
	Image                    *ebiten.Image
}

func NewPassenger(x, y int) *Passenger {

	// Load the passenger image
	passengerImage, _, err := ebitenutil.NewImageFromFile("./game/assets/images/moto-2-narvaez.png")
	if err != nil {
		log.Fatal(err)
	}

	return &Passenger{
		PosX:  x,
		PosY:  y,
		Image: passengerImage,
	}
}
