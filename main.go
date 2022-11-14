package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/claudiu-persoiu/webremote/builder"
	"github.com/claudiu-persoiu/webremote/processor"
	"github.com/claudiu-persoiu/webremote/processor/uinput"
	"github.com/claudiu-persoiu/webremote/structure"

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

func mainPageHandler(data *structure.PageData) func(http.ResponseWriter, *http.Request) {
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

func handleWebServer(page *structure.PageData) {
	http.HandleFunc("/", mainPageHandler(page))
	http.Handle("/js/", http.FileServer(http.Dir("public")))
}

func handleMessageBuilders(b processor.Processor, messagesChan chan structure.Message) {
	keyboardChan := make(chan []string)
	mouseMoveChan := make(chan structure.Offset)
	mouseClickChan := make(chan string)

	go builder.Dispatcher(messagesChan, keyboardChan, mouseMoveChan, mouseClickChan)
	go b.KeyboardCommands(keyboardChan)
	go b.MouseMoveCommands(mouseMoveChan)
	go b.MouseClickCommands(mouseClickChan)
}

func buildKeyboard(file string) *structure.Keyboard {
	keyboardData, err := ioutil.ReadFile(file)

	if err != nil {
		log.Fatal("Could not read keyboard file")
	}

	return structure.NewKeyboard(keyboardData)
}

var address = flag.String("addr", "192.168.1.215:8000", "http service address")

func main() {
	flag.Parse()

	websocketPath := "/echo"

	keyboard := buildKeyboard("keyboard/default.json")

	pageData := &structure.PageData{Title: "Web remote", Address: *address + websocketPath, Keyboard: keyboard.GetJSON()}

	handleWebServer(pageData)

	messagesChan := make(chan structure.Message)
	handleWebSocket(websocketPath, messagesChan)

	//b := xdotool.NewBuilder(keyboard)
	b := uinput.NewBuilder()
	defer b.Close()

	handleMessageBuilders(b, messagesChan)

	log.Printf("Starting listtening on %s... \n", *address)
	log.Fatal(http.ListenAndServe(*address, nil))
}
