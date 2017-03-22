package console

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var out io.Writer = os.Stdout
var in io.Reader = os.Stdin

func Title(s string) {
	fmt.Fprintln(out, s)
	fmt.Fprintln(out, strings.Repeat("=", len(s)))
}

// Question reads input based on the question you pass to it
func Question(question string) (string, error) {
	reader := bufio.NewReader(in)
	fmt.Fprint(out, question+" \x1b[32m>\x1b[0m")
	text, err := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text, err
}

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

// BlackBox outputs white text on a black background
func BlackBox(s string) {
	PrintBox(40, 37, s)
}

// RedBox outputs white text on a red background
func RedBox(s string) {
	PrintBox(41, 37, s)
}

// GreenBox outputs black text on a green background
func GreenBox(s string) {
	PrintBox(42, 30, s)
}

// YellowBox outputs black text on a yellow background
func YellowBox(s string) {
	PrintBox(43, 30, s)
}

// BlueBox outputs black text on a blue background
func BlueBox(s string) {
	PrintBox(44, 30, s)
}

// MagentaBox outputs white text on a magenta background
func MagentaBox(s string) {
	PrintBox(45, 37, s)
}

// CyanBox outputs black text on a cyan background
func CyanBox(s string) {
	PrintBox(46, 30, s)
}

// WhiteBox outputs black text on a white background
func WhiteBox(s string) {
	PrintBox(47, 30, s)
}

// PrintBox outputs colored text on a colored background
func PrintBox(backgroundCode int, fontCode int, s string) {
	fmt.Fprint(out, "\x1b["+strconv.Itoa(backgroundCode)+";"+strconv.Itoa(fontCode)+"m")
	fmt.Fprint(out, " "+s+" ")
	fmt.Fprintln(out, "\x1b[0m")
}
