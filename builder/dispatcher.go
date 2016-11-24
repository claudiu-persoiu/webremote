package builder

import "github.com/claudiu-persoiu/webremote/structure"

func Dispatcher(messages chan structure.Message, keyboard chan []string, mouseMove chan structure.Offset, mouseClick chan string) {
	for {
		message := <-messages

		switch message.Type {
		case "keyboard":
			keyboard <- message.Commands
		case "mousemove":
			mouseMove <- message.Offset
		case "mouseclick":
			mouseClick <- message.Commands[0]
		}
	}
}
