package structure

import (
	"testing"
)

func TestNewKeyboard(t *testing.T) {

	jsonTest := "[[{\"default\":\"p\",\"shift\":\"P\",\"caps\":\"P\"},{\"default\":\"Enter\",\"execute\":\"Return\"}],[{\"default\":\"x\",\"shift\":\"X\",\"caps\":\"X\"}]]"

	k := NewKeyboard([]byte(jsonTest))

	if k.GetJSON() != jsonTest {
		t.Error("Different JSON resulted")
	}

	dictionary := k.GetDictionary()

	testData := map[string]string{"x": "x", "Enter": "Return", "na": ""}

	for input, expected := range testData {
		received := dictionary(input)
		if received != expected {
			t.Errorf("Expected command \"%s\" got \"%s\"", expected, received)
		}
	}
}
