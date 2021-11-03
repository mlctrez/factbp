package blueprint

type Tile struct {
	Name     string    `json:"name"`
	Position *Position `json:"position"`
}

func TileWithPosition(name string, x float64, y float64) (tile *Tile) {
	return &Tile{
		Name:     name,
		Position: &Position{X: x, Y: y},
	}
}
