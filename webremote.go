package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"golang.org/x/net/websocket"
)

type Offset struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Message struct {
	Type     string
	Commands []string `json:"commands,omitempty"`
	Offset   Offset   `json:"offset,omitempty"`
}

type PageData struct {
	Title   string
	Address string
}

func handleWebSocket(path string, commands chan Message) {

	wsHandler := func(ws *websocket.Conn) {
		fmt.Println("connect")
		defer func() {
			log.Println("closing client")
			ws.Close()
		}()

		msg := new(Message)
		for {
			err := websocket.JSON.Receive(ws, msg)

			if err != nil {
				log.Println(err)
				return
			}
			fmt.Println(msg)
			fmt.Printf("Receive: %s\n", msg.Commands)
			commands <- *msg
		}
	}

	http.Handle(path, websocket.Handler(wsHandler))
}

func execCommand(cmd string) ([]byte, error) {
	fmt.Println("command is ", cmd)

	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:]

	return exec.Command(head, parts...).Output()
}

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

func executeCommands(commands chan string) {
	for {
		command := <-commands
		out, err := execCommand("xdotool " + command)
		if err != nil {
			log.Println("Eroare", err)
		}
		log.Println(out)
	}
}

func buildKeyboardCommands(keyboard chan Message, commands chan string) {
	dictionary := makeDictionary()

	for {
		keyboardMessage := <-keyboard
		command := ""
		for _, key := range keyboardMessage.Commands {
			command += dictionary(key) + "+"
		}
		commands <- "key " + command[0:len(command)-1]
	}
}

func buildMouseMoveCommands(mouseMove chan Message, commands chan string) {
	for {
		mouseMoveMessage := <-mouseMove
		commands <- "mousemove_relative -- " + strconv.Itoa(mouseMoveMessage.Offset.X) + " " + strconv.Itoa(mouseMoveMessage.Offset.Y)
	}
}

func buildMouseClickCommands(mouseClick chan Message, commands chan string) {
	for {
		mouseClickMessage := <-mouseClick
		commands <- "click " + mouseClickMessage.Commands[0]
	}
}

func commandsDispatcher(messages chan Message, keyboard chan Message, mouseMove chan Message, mouseClick chan Message) {
	for {
		message := <-messages

		switch message.Type {
		case "keyboard":
			keyboard <- message
		case "mousemove":
			mouseMove <- message
		case "mouseclick":
			mouseClick <- message
		}
	}
}

func mainHandler(data *PageData) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("public/index.html")

		if err != nil {
			log.Panic(err)
		}

		err = t.Execute(w, *data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

var address = flag.String("addr", "localhost:8000", "http service address")

func main() {
	flag.Parse()

	websocketPath := "/echo"

	data := &PageData{Title: "Web remote", Address: *address + websocketPath}

	http.HandleFunc("/", mainHandler(data))
	http.Handle("/js/", http.FileServer(http.Dir("public")))
	messages := make(chan Message)
	keyboard := make(chan Message)
	mouseMove := make(chan Message)
	mouseClick := make(chan Message)
	commands := make(chan string)

	handleWebSocket(websocketPath, messages)

	// https://jan.newmarch.name/go/template/chapter-template.html

	go commandsDispatcher(messages, keyboard, mouseMove, mouseClick)
	go buildKeyboardCommands(keyboard, commands)
	go buildMouseClickCommands(mouseClick, commands)
	go buildMouseMoveCommands(mouseMove, commands)
	go executeCommands(commands)
	log.Printf("Starting listtening on %s... \n", *address)
	log.Fatal(http.ListenAndServe(*address, nil))
}
