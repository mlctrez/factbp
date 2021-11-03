package blueprint

type Connection struct {
	EntityId  *uint64 `json:"entity_id,omitempty"`
	CircuitId *uint64 `json:"circuit_id,omitempty"`
}
