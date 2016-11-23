package structure


type Offset struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Message struct {
	Type     string
	Commands []string `json:"commands,omitempty"`
	Offset   Offset   `json:"offset,omitempty"`
}