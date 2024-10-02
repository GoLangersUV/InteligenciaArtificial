package entities

type Passenger struct {
	PosX, PosY int
}

func NewPassenger(x, y int) *Passenger {
	return &Passenger{
		PosX: x,
		PosY: y,
	}
}
