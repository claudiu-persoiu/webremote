package processor

import (
	"os/exec"
	"strings"
	"fmt"
)

var execCommand = exec.Command;
var logLine = fmt.Println;

func ProcessCommands(commands chan string) {
	for {
		command := <-commands
		out, err := buildCommand("xdotool " + command)
		if err != nil {
			logLine("Error: ", err)
		}
		logLine(out)
	}
}

func buildCommand(cmd string) ([]byte, error) {
	logLine("command is ", cmd)

	parts := strings.Fields(cmd)
	head := parts[0]
	parts = parts[1:]

	return execCommand(head, parts...).Output()
}
