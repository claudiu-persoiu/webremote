package main

import (
	"flag"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/claudiu-persoiu/webremote/builder"
	"github.com/claudiu-persoiu/webremote/processor"
	"github.com/claudiu-persoiu/webremote/structure"
	"github.com/gobuffalo/packr"

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

var box = packr.NewBox("./public")
var keyboardBox = packr.NewBox("./keyboard")

func mainPageHandler(data *structure.PageData) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		b, err := box.FindString("index.html")
		if err != nil {
			log.Panic(err)
		}

		t, err := template.New("hello").Parse(b)

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
	http.Handle("/js/", http.FileServer(box))
}

func handleMessageBuilders(messagesChan chan structure.Message, commandsChan chan string, keyboard *structure.Keyboard) {
	keyboardChan := make(chan []string)
	mouseMoveChan := make(chan structure.Offset)
	mouseClickChan := make(chan string)

	go builder.Dispatcher(messagesChan, keyboardChan, mouseMoveChan, mouseClickChan)
	go builder.KeyboardCommands(keyboardChan, commandsChan, keyboard)
	go builder.MouseClickCommands(mouseClickChan, commandsChan)
	go builder.MouseMoveCommands(mouseMoveChan, commandsChan)
}

func buildKeyboard(file string) *structure.Keyboard {
	// keyboardData, err := ioutil.ReadFile(file)

	keyboardData, err := keyboardBox.FindString(file)

	if err != nil {
		log.Fatal("Could not read keyboard file")
	}

	return structure.NewKeyboard([]byte(keyboardData))
}

var address = flag.String("addr", "localhost:8000", "http service address")

func main() {
	flag.Parse()

	websocketPath := "/echo"

	keyboard := buildKeyboard("default.json")

	pageData := &structure.PageData{Title: "Web remote", Address: *address + websocketPath, Keyboard: keyboard.GetJSON()}

	handleWebServer(pageData)

	messagesChan := make(chan structure.Message)
	commandsChan := make(chan string)
	handleWebSocket(websocketPath, messagesChan)
	handleMessageBuilders(messagesChan, commandsChan, keyboard)
	go processor.ProcessCommands(commandsChan)

	log.Printf("Starting listtening on %s... \n", *address)
	log.Fatal(http.ListenAndServe(*address, nil))
}
