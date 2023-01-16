package logger

import "fmt"

var verbose bool = false

func SetVerbose(v bool) {
	verbose = v
}

func Log(a ...interface{}) (int, error) {
	if verbose {
		return fmt.Println(a...)
	}
	return 0, nil
}
