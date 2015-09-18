package goerrui

import (
	"fmt"
	"os"

	"github.com/shiena/ansicolor"
)

var w = ansicolor.NewAnsiColorWriter(os.Stderr)

var Levels = struct {
	Alert    int
	Critical int
	Warn     int
	Info     int
	Message  int
}{50, 40, 30, 20, 10}

var LogLevel = Levels.Warn

func SetLevel(level int) {
	LogLevel = level
}

//50
func Alert(format string, a ...interface{}) {
	if LogLevel > Levels.Alert {
		return
	}

	text := "%s%v%s\n"
	msg := fmt.Sprintf(format, a...)
	fmt.Fprintf(w, text, "\x1b[31m", msg, "\x1b[0m")
}

//10
func Message(format string, a ...interface{}) {
	if LogLevel > Levels.Message {
		return
	}

	text := "%s%v%s\n"
	msg := fmt.Sprintf(format, a...)
	fmt.Fprintf(w, text, "\x1b[36m", msg, "\x1b[0m")
}

//30
func Warn(format string, a ...interface{}) {
	if LogLevel > Levels.Warn {
		return
	}

	text := "%s%v%s\n"
	msg := fmt.Sprintf(format, a...)
	fmt.Fprintf(w, text, "\x1b[33m", msg, "\x1b[0m")
}
