package xdotool

import "github.com/claudiu-persoiu/webremote/structure"

type Builder struct {
	commands chan string
	keyboard *structure.Keyboard
}

func NewBuilder(keyboard *structure.Keyboard) *Builder {
	commands := make(chan string)
	go processCommands(commands)

	return &Builder{
		commands: commands,
		keyboard: keyboard,
	}
}

func (b *Builder) Close() {
	// close resources
}
