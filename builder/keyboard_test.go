package builder

import (
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

	go KeyboardCommands(keyboard, commands)

	time.Sleep(time.Millisecond * 10)

	for _, expected := range expectedCommands {
		received := <-commands
		if expected != received {
			t.Errorf("Expected command \"%s\" got \"%s\"", expected, received)
		}
	}
}
