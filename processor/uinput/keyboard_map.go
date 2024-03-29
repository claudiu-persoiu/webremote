package uinput

import "github.com/bendahl/uinput"

var keyboardMap = map[string]int{
	"Esc":    uinput.KeyEsc,
	"F1":     uinput.KeyF1,
	"F2":     uinput.KeyF2,
	"F3":     uinput.KeyF3,
	"F4":     uinput.KeyF4,
	"F5":     uinput.KeyF5,
	"F6":     uinput.KeyF6,
	"F7":     uinput.KeyF7,
	"F8":     uinput.KeyF8,
	"F9":     uinput.KeyF9,
	"F10":    uinput.KeyF10,
	"F11":    uinput.KeyF11,
	"F12":    uinput.KeyF12,
	"`":      uinput.KeyGrave,
	"1":      uinput.Key1,
	"2":      uinput.Key2,
	"3":      uinput.Key3,
	"4":      uinput.Key4,
	"5":      uinput.Key5,
	"6":      uinput.Key6,
	"7":      uinput.Key7,
	"8":      uinput.Key8,
	"9":      uinput.Key9,
	"0":      uinput.Key0,
	"-":      uinput.KeyMinus,
	"=":      uinput.KeyEqual,
	"Bksp":   uinput.KeyBackspace,
	"Tab":    uinput.KeyTab,
	"q":      uinput.KeyQ,
	"w":      uinput.KeyW,
	"e":      uinput.KeyE,
	"r":      uinput.KeyR,
	"t":      uinput.KeyT,
	"y":      uinput.KeyY,
	"u":      uinput.KeyU,
	"i":      uinput.KeyI,
	"o":      uinput.KeyO,
	"p":      uinput.KeyP,
	"[":      uinput.KeyLeftbrace,
	"]":      uinput.KeyRightbrace,
	"\\":     uinput.KeyBackslash,
	"a":      uinput.KeyA,
	"s":      uinput.KeyS,
	"d":      uinput.KeyD,
	"f":      uinput.KeyF,
	"g":      uinput.KeyG,
	"h":      uinput.KeyH,
	"j":      uinput.KeyJ,
	"k":      uinput.KeyK,
	"l":      uinput.KeyL,
	";":      uinput.KeySemicolon,
	"'":      uinput.KeyApostrophe,
	"Enter":  uinput.KeyEnter,
	"z":      uinput.KeyZ,
	"x":      uinput.KeyX,
	"c":      uinput.KeyC,
	"b":      uinput.KeyB,
	"n":      uinput.KeyN,
	"m":      uinput.KeyM,
	",":      uinput.KeyComma,
	".":      uinput.KeyDot,
	"/":      uinput.KeySlash,
	"Win":    uinput.KeyLeftmeta,
	"Space":  uinput.KeySpace,
	"&larr;": uinput.KeyLeft,
	"&uarr;": uinput.KeyUp,
	"&darr;": uinput.KeyDown,
	"&rarr;": uinput.KeyRight,
}

var combinationKeysMap = map[string]int{
	"Shift": uinput.KeyLeftshift,
	"Ctrl":  uinput.KeyLeftctrl,
	"Alt":   uinput.KeyLeftalt,
}
