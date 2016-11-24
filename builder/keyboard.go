package builder

import "github.com/claudiu-persoiu/webremote/structure"

func KeyboardCommands(keyboardChan chan []string, commands chan string, keyboard *structure.Keyboard) {
	dictionary := keyboard.GetDictionary()

	for {
		keyboardMessage := <-keyboardChan
		command := ""
		for _, key := range keyboardMessage {
			command += dictionary(key) + "+"
		}
		commands <- "key " + command[0:len(command)-1]
	}
}
