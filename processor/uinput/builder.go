package uinput

import (
	"log"

	"github.com/bendahl/uinput"
	"github.com/claudiu-persoiu/webremote/logger"
	"github.com/claudiu-persoiu/webremote/structure"
)

type Builder struct {
	keyboard uinput.Keyboard
	mouse    uinput.Mouse
}

func NewBuilder() *Builder {
	keyboard, err := uinput.CreateKeyboard("/dev/uinput", []byte("testkeyboard"))
	if err != nil {
		log.Fatal("Unable to created virtual keyboard: " + err.Error())
	}
	mouse, err := uinput.CreateMouse("/dev/uinput", []byte("testmouse"))
	if err != nil {
		log.Fatal("Unable to created virtual mouse: " + err.Error())
	}
	return &Builder{
		keyboard: keyboard,
		mouse:    mouse,
	}
}

func (b *Builder) KeyboardCommands(keyboardChan chan []string) {
	for {
		keyboardMessage := <-keyboardChan
		combList := make([]int, 0)
		for _, key := range keyboardMessage {
			filteredComb, okComb := combinationKeysMap[key]
			if okComb {
				b.keyboard.KeyDown(filteredComb)
				combList = append(combList, filteredComb)
			}
			filtered, ok := keyboardMap[key]
			if ok {
				b.keyboard.KeyPress(filtered)
			}
		}
		for _, key := range combList {
			b.keyboard.KeyUp(key)
		}
	}
}

func (b *Builder) MouseMoveCommands(mouseMoveChan chan structure.Offset) {
	for {
		offset := <-mouseMoveChan
		if offset.X != 0 {
			if offset.X > 0 {
				b.mouse.MoveRight(int32(offset.X))
			} else {
				b.mouse.MoveLeft(int32(-offset.X))
			}
		}
		if offset.Y != 0 {
			if offset.Y > 0 {
				b.mouse.MoveDown(int32(offset.Y))
			} else {
				b.mouse.MoveUp(int32(-offset.Y))
			}
		}
	}
}
func (b *Builder) MouseClickCommands(mouseClickChan chan string) {
	for {
		action := <-mouseClickChan
		logger.Log(action)
		switch action {
		case "1":
			b.mouse.LeftClick()
			break
		case "2":
			b.mouse.MiddleClick()
			break
		case "3":
			b.mouse.RightClick()
			break
		case "4":
			b.mouse.Wheel(false, 5)
			break
		case "5":
			b.mouse.Wheel(false, -5)
			break
		}
	}
}

func (b *Builder) Close() {
	b.keyboard.Close()
	b.mouse.Close()
}
