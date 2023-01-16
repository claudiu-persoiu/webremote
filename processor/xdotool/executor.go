package xdotool

import (
	"os/exec"
	"strings"

	"github.com/claudiu-persoiu/webremote/logger"
)

var execCommand = exec.Command
var logLine = logger.Log

func processCommands(commands chan string) {
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
