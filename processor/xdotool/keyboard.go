package xdotool

import "strings"

func (b *Builder) KeyboardCommands(keyboardChan chan []string) {
	dictionary := b.keyboard.GetDictionary()

	for {
		keyboardMessage := <-keyboardChan
		var command strings.Builder
		for _, key := range keyboardMessage {
			filtered := dictionary(key)
			if len(filtered) > 0 {
				command.WriteString(filtered + "+")
			}
		}

		if len(command.String()) > 0 {
			b.commands <- "key " + command.String()[0:len(command.String())-1]
		}
	}
}
