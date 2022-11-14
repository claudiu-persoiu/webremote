package xdotool

import (
	"strconv"

	"github.com/claudiu-persoiu/webremote/structure"
)

func (b *Builder) MouseMoveCommands(mouseMove chan structure.Offset) {
	for {
		offsetMessage := <-mouseMove
		b.commands <- "mousemove_relative -- " + strconv.Itoa(offsetMessage.X) + " " + strconv.Itoa(offsetMessage.Y)
	}
}

func (b *Builder) MouseClickCommands(mouseClick chan string) {
	for {
		b.commands <- "click " + <-mouseClick
	}
}
