package builder

import (
	"github.com/claudiu-persoiu/webremote/structure"
	"strconv"
)

func MouseMoveCommands(mouseMove chan structure.Offset, commands chan string) {
	for {
		offsetMessage := <-mouseMove
		commands <- "mousemove_relative -- " + strconv.Itoa(offsetMessage.X) + " " + strconv.Itoa(offsetMessage.Y)
	}
}

func MouseClickCommands(mouseClick chan string, commands chan string) {
	for {
		commands <- "click " + <-mouseClick
	}
}
