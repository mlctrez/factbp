package blueprint

// Base is the common fields for a book or blueprint
type Base struct {
	Item        *string `json:"item,omitempty"` // blueprint or blueprint-book
	Version     *uint64 `json:"version,omitempty"`
	Label       *string `json:"label,omitempty"`
	Icons       []*Icon `json:"icons,omitempty"`
	LabelColor  *string `json:"label_color,omitempty"`
	Description *string `json:"description,omitempty"`
}
