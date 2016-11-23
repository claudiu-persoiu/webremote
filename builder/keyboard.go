package builder

func makeDictionary() func(string) string {
	commands := map[string]string{
		"Bksp":   "BackSpace",
		"Esc":    "Escape",
		"Enter":  "Return",
		"Men":    "Menu",
		"Win":    "Super",
		"&larr;": "Left",
		"&rarr;": "Right",
		"&uarr;": "Up",
		"&darr;": "Down",
		"Space":  "space"}

	return func(key string) string {
		if translate, found := commands[key]; found {
			return translate
		}

		return key
	}
}

func KeyboardCommands(keyboard chan []string, commands chan string) {
	dictionary := makeDictionary()

	for {
		keyboardMessage := <-keyboard
		command := ""
		for _, key := range keyboardMessage {
			command += dictionary(key) + "+"
		}
		commands <- "key " + command[0:len(command)-1]
	}
}