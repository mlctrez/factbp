package blueprint

type Blueprint struct {
	Base
	Entities               []*Entity    `json:"entities,omitempty"`
	Tiles                  []*Tile      `json:"tiles,omitempty"`
	Schedules              []*Schedule  `json:"schedules,omitempty"`
	// SnapToGrid will be non-nil when snap to grid is enabled.
	//
	// Minimum value is 1,1 and for blueprints containing rails x and y must be multiples of 2
	SnapToGrid             *PositionInt `json:"snap-to-grid,omitempty"`
	// AbsoluteSnapping must be present when SnapToGrid is set and indicates the type of snapping.
	AbsoluteSnapping       *bool        `json:"absolute-snapping,omitempty"`         // true when absolute snapping selected
	// PositionRelativeToGrid will be present only if absolute snapping is true.
	PositionRelativeToGrid *PositionInt `json:"position-relative-to-grid,omitempty"` // present when abs snapping on and not 0,0
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
