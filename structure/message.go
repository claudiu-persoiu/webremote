package structure

import (
	"fmt"
	"strconv"
)

type Offset struct {
	X int `json:"x"`
	Y int `json:"y"`
}

func (o Offset) String() string {
	return fmt.Sprintf("x: %s, y: %s", strconv.Itoa(o.X), strconv.Itoa(o.Y))
}

type Message struct {
	Type     string
	Commands []string `json:"commands,omitempty"`
	Offset   Offset   `json:"offset,omitempty"`
}
