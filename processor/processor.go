package processor

import "github.com/claudiu-persoiu/webremote/structure"

type Processor interface {
	KeyboardCommands(keyboardChan chan []string)
	MouseMoveCommands(mouseMoveChan chan structure.Offset)
	MouseClickCommands(mouseClickChan chan string)
	Close()
}
