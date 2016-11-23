package main

import (
	"flag"
	"fmt"
	"github.com/claudiu-persoiu/webremote/processor"
	"github.com/claudiu-persoiu/webremote/structure"
	"github.com/claudiu-persoiu/webremote/builder"
	"html/template"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

func handleWebSocket(path string, commands chan structure.Message) {

	wsHandler := func(ws *websocket.Conn) {
		fmt.Println("connect")
		defer func() {
			log.Println("closing client")
			ws.Close()
		}()

		msg := new(structure.Message)
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

func mainHandler(data *structure.PageData) func(http.ResponseWriter, *http.Request) {
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

	data := &structure.PageData{Title: "Web remote", Address: *address + websocketPath}

	http.HandleFunc("/", mainHandler(data))
	http.Handle("/js/", http.FileServer(http.Dir("public")))

	messages := make(chan structure.Message)
	keyboard := make(chan []string)
	mouseMove := make(chan structure.Offset)
	mouseClick := make(chan string)
	commands := make(chan string)

	handleWebSocket(websocketPath, messages)

	go builder.Dispatcher(messages, keyboard, mouseMove, mouseClick)
	go builder.KeyboardCommands(keyboard, commands)
	go builder.MouseClickCommands(mouseClick, commands)
	go builder.MouseMoveCommands(mouseMove, commands)
	go processor.ProcessCommands(commands)
	log.Printf("Starting listtening on %s... \n", *address)
	log.Fatal(http.ListenAndServe(*address, nil))
}
