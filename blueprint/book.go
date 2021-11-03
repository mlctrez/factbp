package blueprint

type Book struct {
	Item        string              `json:"item"`
	Label       string              `json:"label,omitempty"`
	LabelColor  string              `json:"label_color,omitempty"`
	ActiveIndex uint64              `json:"active_index"`
	Blueprints  []*IndexedBlueprint `json:"blueprints"`
	Version     *uint64             `json:"version,omitempty"`
}

type IndexedBlueprint struct {
	Index     uint64     `json:"index"`
	Blueprint *Blueprint `json:"blueprint"`
}

func (b *Book) AddBlueprint(bp *Blueprint) {
	if b.Blueprints == nil {
		b.Blueprints = make([]*IndexedBlueprint, 0)
	}
	index := uint64(len(b.Blueprints))
	b.Blueprints = append(b.Blueprints, &IndexedBlueprint{
		Index:     index,
		Blueprint: bp,
	})

}
