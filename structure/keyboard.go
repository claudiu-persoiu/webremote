package structure

import (
	"encoding/json"
	"fmt"
	"os"
)

type Key struct {
	Default string `json:"default"`
	Shift   string `json:"shift,omitempty"`
	Caps    string `json:"caps,omitempty"`
	Execute string `json:"execute,omitempty"`
}

type Keyboard struct {
	Keys [][]Key
}

func NewKeyboard(rawJSON []byte) *Keyboard {

	k := new(Keyboard)

	err := json.Unmarshal(rawJSON, &k.Keys)
	if err != nil {
		fmt.Println("Unable to decode JSON")
		os.Exit(1)
	}

	return k
}

func (k *Keyboard) GetDictionary() func(string) string {
	commands := make(map[string]string)

	for _, row := range k.Keys {
		for _, key := range row {
			if len(key.Execute) > 0 {
				commands[key.Default] = key.Execute
			} else {
				commands[key.Default] = key.Default
			}
		}
	}

	return func(key string) string {
		if translate, found := commands[key]; found {
			return translate
		}

		return ""
	}
}

func (k *Keyboard) GetJSON() string {
	res, _ := json.Marshal(k.Keys)

	return string(res)
}
