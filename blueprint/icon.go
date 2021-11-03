package blueprint

type Icon struct {
	Index  uint64            `json:"index,omitempty"`
	Signal map[string]string `json:"signal,omitempty"`
}
