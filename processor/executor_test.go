package processor

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
	"time"
)

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}

func TestProcessCommands(t *testing.T) {
	lineToPrint := ""
	execCommand = fakeExecCommand
	logLine = func(lines ...interface{}) (int, error) {
		if line, ok := lines[len(lines)-1].(string); ok {
			lineToPrint = line
		}
		return 0, nil
	}

	defer func() {
		execCommand = exec.Command
		logLine = fmt.Println
	}()

	command := "multiple arguments"
	expectedResult := "xdotool " + command
	commands := make(chan string, 1)
	commands <- command

	go ProcessCommands(commands)

	time.Sleep(time.Millisecond + 10)
	if lineToPrint != expectedResult {
		t.Errorf("Wrong command, expected: \"%s\"", expectedResult)
	}
}
