package builder

import (
	"github.com/claudiu-persoiu/webremote/structure"
	"testing"
	"time"
)

func TestKeyboardCommands(t *testing.T) {
	inputCommands := [][]string{
		[]string{"x"},
		[]string{"y", "z"},
		[]string{"&darr;"}}

	expectedCommands := []string{"key x", "key y+z", "key Down"}

	keyboard := make(chan []string, len(inputCommands))
	commands := make(chan string, len(inputCommands))

	for _, command := range inputCommands {
		keyboard <- command
	}

	testKeyboard := []byte("[[{\"default\":\"x\",\"shift\":\"X\",\"caps\":\"X\"},{\"default\":\"y\",\"shift\":\"Y\",\"caps\":\"Y\"},{\"default\":\"z\",\"shift\":\"Z\",\"caps\":\"Z\"},{\"default\":\"&darr;\",\"execute\":\"Down\"}]]")

	go KeyboardCommands(keyboard, commands, structure.NewKeyboard(testKeyboard))

	time.Sleep(time.Millisecond * 10)

	for _, expected := range expectedCommands {
		received := <-commands
		if expected != received {
			t.Errorf("Expected command \"%s\" got \"%s\"", expected, received)
		}
	}
}
