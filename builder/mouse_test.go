package builder

import (
	"testing"
	"github.com/claudiu-persoiu/webremote/structure"
	"time"
)

func TestMouseMoveCommands(t *testing.T) {

	inputCommands := []structure.Offset{structure.Offset{X: 10, Y: 20},
		structure.Offset{X: 20, Y: 30}}
	expectedCommands := []string{"mousemove_relative -- 10 20", "mousemove_relative -- 20 30"}

	mouseMove := make(chan structure.Offset, len(inputCommands))
	commands := make(chan string, 3)

	for _, input := range inputCommands {
		mouseMove <- input
	}

	go MouseMoveCommands(mouseMove, commands)

	time.Sleep(time.Millisecond * 200)

	for _, expected := range expectedCommands {
		received := <-commands
		if expected != received {
			t.Errorf("Expected command \"%s\" got \"%s\"", expected, received)
		}
	}
}

func TestMouseClickCommands(t *testing.T) {

	inputCommands := []string{"1", "3"}
	expectedCommands := []string{"click 1", "click 3"}

	mouseMove := make(chan string, len(inputCommands))
	commands := make(chan string, 3)

	for _, input := range inputCommands {
		mouseMove <- input
	}

	go MouseClickCommands(mouseMove, commands)

	time.Sleep(time.Millisecond * 200)

	for _, expected := range expectedCommands {
		received := <-commands
		if expected != received {
			t.Errorf("Expected command \"%s\" got \"%s\"", expected, received)
		}
	}
}
