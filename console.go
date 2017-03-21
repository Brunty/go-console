package console

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

var out io.Writer = os.Stdout

// Black outputs in ansi color black
func Black(s string) {
	PrintLine(30, s)
}

// Red outputs in ansi color red
func Red(s string) {
	PrintLine(31, s)
}

// Green outputs in ansi color green
func Green(s string) {
	PrintLine(32, s)
}

// Yellow outputs in ansi color yellow
func Yellow(s string) {
	PrintLine(33, s)
}

// Blue outputs in ansi color blue
func Blue(s string) {
	PrintLine(34, s)
}

// Magenta outputs in ansi color magenta
func Magenta(s string) {
	PrintLine(35, s)
}

// Cyan outputs in ansi color cyan
func Cyan(s string) {
	PrintLine(36, s)
}

// White outputs in ansi color white
func White(s string) {
	PrintLine(37, s)
}

// PrintLine outputs in ansi color specified by the code
func PrintLine(code int, s string) {
	fmt.Fprint(out, "\x1b["+strconv.Itoa(code)+"m")
	fmt.Fprint(out, s)
	fmt.Fprintln(out, "\x1b[0m")
}
