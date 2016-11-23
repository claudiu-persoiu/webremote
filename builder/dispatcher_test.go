package builder

import (
	"testing"
	"github.com/claudiu-persoiu/webremote/structure"
	"time"
	"fmt"
)

func TestDispatcher(t *testing.T) {

	inputCommands := []structure.Message{structure.Message{Type: "keyboard", Commands: []string{"s"}},
		structure.Message{Type: "keyboard", Commands: []string{"Shift", "s"}},
		structure.Message{Type: "keyboard", Commands: []string{"Alt", "Shift", "s"}},
		structure.Message{Type: "mouseclick", Commands: []string{"1"}},
		structure.Message{Type: "mouseclick", Commands: []string{"3"}},
		structure.Message{Type: "mousemove", Offset: structure.Offset{X: 1, Y: 2}},
		structure.Message{Type: "mousemove", Offset: structure.Offset{X: 4, Y: 0}}}

	expectedKeyboard := [][]string{
		[]string{"s"},
		[]string{"Shift", "s"},
		[]string{"Alt", "Shift", "s"}}

	expectedMouseClick := []string{"1", "3"}

	expectedMouseMove := []structure.Offset{structure.Offset{X: 1, Y: 2},
		structure.Offset{X: 4, Y: 0}}


	messages := make(chan structure.Message, len(inputCommands))
	keyboard := make(chan []string, len(expectedKeyboard))
	mouseClick := make(chan string, len(expectedMouseClick))
	mouseMove := make(chan structure.Offset, len(expectedMouseMove))

	for _, input := range inputCommands {
		messages <- input
	}

	go Dispatcher(messages, keyboard, mouseMove, mouseClick)

	time.Sleep(time.Millisecond * 100)

	close(keyboard)
	close(mouseMove)
	close(mouseClick)

	for _, expected := range expectedKeyboard {
		received := <-keyboard
		if fmt.Sprintf("%s", expected) != fmt.Sprintf("%s", received) {
			t.Errorf("Expected command \"%s\" got \"%s\"", expected, received)
		}
	}

	for _, expected := range expectedMouseMove {
		received := <-mouseMove
		if expected != received {
			t.Errorf("Expected command \"%s\" got \"%s\"", expected, received)
		}
	}

	for _, expected := range expectedMouseClick {
		received := <-mouseClick
		if expected != received {
			t.Errorf("Expected command \"%s\" got \"%s\"", expected, received)
		}
	}
}