package blueprint

type Blueprint struct {
	Item       *string     `json:"item,omitempty"`
	Label      *string     `json:"label,omitempty"`
	LabelColor *string     `json:"label_color,omitempty"`
	Entities   []*Entity   `json:"entities,omitempty"`
	Tiles      []*Tile     `json:"tiles,omitempty"`
	Icons      []*Icon     `json:"icons,omitempty"`
	Schedules  []*Schedule `json:"schedules,omitempty"`
	Version    *uint64     `json:"version,omitempty"`
}

func (b *Blueprint) AddEntity(e *Entity) {
	if b.Entities == nil {
		b.Entities = []*Entity{}
	}
	e.Number = uint64(len(b.Entities) + 1)
	b.Entities = append(b.Entities, e)
}

func (b *Blueprint) AddTile(t *Tile) {
	if b.Tiles == nil {
		b.Tiles = []*Tile{}
	}
	b.Tiles = append(b.Tiles, t)
}

