package datatypes

// Matris representation of the board
type Matrix = [][]int

type ScannedMatrix struct {
	Matrix          Matrix
	MainCoordinates map[string]BoardCoordinate
}
