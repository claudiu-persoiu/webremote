package main

import (
	"embed"
	"flag"
	"fmt"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/claudiu-persoiu/webremote/builder"
	"github.com/claudiu-persoiu/webremote/logger"
	"github.com/claudiu-persoiu/webremote/processor"
	"github.com/claudiu-persoiu/webremote/processor/uinput"
	"github.com/claudiu-persoiu/webremote/processor/xdotool"
	"github.com/claudiu-persoiu/webremote/structure"

	"golang.org/x/net/websocket"
)

func handleWebSocket(path string, commands chan structure.Message) {

	wsHandler := func(ws *websocket.Conn) {
		logger.Log("connect")
		defer func() {
			logger.Log("closing client")
			ws.Close()
		}()

		msg := new(structure.Message)
		for {
			err := websocket.JSON.Receive(ws, msg)

			if err != nil {
				logger.Log(err)
				return
			}
			logger.Log(msg)
			logger.Log("Receive: %s\n", msg.Commands)
			commands <- *msg
		}
	}

	http.Handle(path, websocket.Handler(wsHandler))
}

func mainPageHandler(data *structure.PageData) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFS(templateBox, "template/index.html")

		if err != nil {
			log.Fatal(err)
		}

		err = t.Execute(w, *data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

//go:embed public
var publicBox embed.FS

//go:embed template
var templateBox embed.FS

//go:embed keyboard
var keyboardBox embed.FS

func handleWebServer(page *structure.PageData) {
	http.HandleFunc("/", mainPageHandler(page))
	http.Handle("/public/", http.FileServer(http.FS(publicBox)))
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
	keyboardData, err := keyboardBox.ReadFile(file)

	if err != nil {
		log.Fatal("Could not read keyboard file")
	}

	return structure.NewKeyboard(keyboardData)
}

// Get preferred outbound ip of this machine
func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

var port = flag.String("port", "8765", "http service port")
var exec = flag.String("exec", "uinput", "command executor, options are uinput and xdotool")
var verbose = flag.Bool("verbose", false, "make verbose")
var jiggler = flag.Bool("jiggler", false, "just jiggle the mouse to prevent sleep")

func main() {
	flag.Parse()

	keyboard := buildKeyboard("keyboard/default.json")

	var b processor.Processor
	switch *exec {
	case "uinput":
		b = uinput.NewBuilder()
		break
	case "xdotool":
		b = xdotool.NewBuilder(keyboard)
		break
	default:
		log.Fatal("Invalid command executor")
	}
	defer b.Close()

	if *jiggler {
		fmt.Println("Starting mouse jiggler...")
		go func() {
			mouseMoveChan := make(chan structure.Offset)
			go b.MouseMoveCommands(mouseMoveChan)
			for {
				mouseMoveChan <- structure.Offset{X: 10, Y: 10}
				mouseMoveChan <- structure.Offset{X: -10, Y: -10}
				// Sleep for a minute
				select {
				case <-time.After(10 * time.Second):
				}
			}
		}()
		fmt.Scanln()
	} else {
		websocketPath := "/echo"
		logger.SetVerbose(*verbose)

		pageData := &structure.PageData{Title: "Web remote", Address: *port + websocketPath, Keyboard: keyboard.GetJSON()}

		handleWebServer(pageData)

		messagesChan := make(chan structure.Message)
		handleWebSocket(websocketPath, messagesChan)

		handleMessageBuilders(b, messagesChan)

		address := fmt.Sprintf("%s:%s", getOutboundIP(), *port)
		fmt.Printf("Starting listening on...\n http://%s... \n", address)
		if host, err := os.Hostname(); err == nil {
			fmt.Printf(" http://%s:%s... \n", host, *port)
		}
		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", *port), nil))
	}
}
