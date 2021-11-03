package blueprint

type Entity struct {
	Number             uint64                              `json:"entity_number"`
	Name               string                              `json:"name"`
	Position           *Position                           `json:"position"`
	Direction          *uint64                             `json:"direction,omitempty"`
	Connections        map[string]map[string][]*Connection `json:"connections,omitempty"`
	ControlBehaviour   interface{}                         `json:"control_behavior,omitempty"`
	Recipe             *string                             `json:"recipe,omitempty"`
	Bar                *uint64                             `json:"bar,omitempty"`
	RequestFilters     []*LogisticFilter                   `json:"request_filters,omitempty"`
	RequestFromBuffers *bool                               `json:"request_from_buffers,omitempty"`
	Variation          *uint64                             `json:"variation,omitempty"`
}

func (e *Entity) AddConnection(key string, color string, connection *Connection) {
	if e.Connections == nil {
		e.Connections = make(map[string]map[string][]*Connection)
	}
	if _, ok := e.Connections[key]; !ok {
		e.Connections[key] = make(map[string][]*Connection)
	}
	if _, ok := e.Connections[key][color]; !ok {
		e.Connections[key][color] = make([]*Connection, 0)
	}
	e.Connections[key][color] = append(e.Connections[key][color], connection)
}
