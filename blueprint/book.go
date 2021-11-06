package blueprint

type Book struct {
	Base
	ActiveIndex uint64       `json:"active_index"`
	Blueprints  []*Container `json:"blueprints"`
}

func (b *Book) AddBlueprint(blueprint *Blueprint) {
	if b.Blueprints == nil {
		b.Blueprints = make([]*Container, 0)
	}
	index := uint64(len(b.Blueprints))
	c := &Container{Blueprint: blueprint, Index: &index}
	b.Blueprints = append(b.Blueprints, c)
}

func (b *Book) AddBook(book *Book) {
	if b.Blueprints == nil {
		b.Blueprints = make([]*Container, 0)
	}
	index := uint64(len(b.Blueprints))
	c := &Container{Book: book, Index: &index}
	b.Blueprints = append(b.Blueprints, c)
}
