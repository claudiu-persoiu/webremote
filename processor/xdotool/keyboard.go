package xdotool

func (b *Builder) KeyboardCommands(keyboardChan chan []string) {
	dictionary := b.keyboard.GetDictionary()

	for {
		keyboardMessage := <-keyboardChan
		command := ""
		for _, key := range keyboardMessage {
			filtered := dictionary(key)
			if len(filtered) > 0 {
				command += filtered + "+"
			}
		}

		if len(command) > 0 {
			b.commands <- "key " + command[0:len(command)-1]
		}
	}
}
