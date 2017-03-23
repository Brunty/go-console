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
var tags = make(map[string]string)

func colors() map[string]int {
	colors := make(map[string]int)

	colors["black"] = 0
	colors["red"] = 1
	colors["green"] = 2
	colors["yellow"] = 3
	colors["blue"] = 4
	colors["magenta"] = 5
	colors["cyan"] = 6
	colors["white"] = 7

	return colors
}

func styles() map[string]string {
	tags["info"] = "green"
	tags["note"] = "blue"
	tags["comment"] = "cyan"
	tags["warning"] = "yellow"
	tags["error"] = "red"
	tags["black"] = "black"
	tags["red"] = "red"
	tags["green"] = "green"
	tags["yellow"] = "yellow"
	tags["blue"] = "blue"
	tags["magenta"] = "magenta"
	tags["cyan"] = "cyan"
	tags["white"] = "white"

	return tags
}

// colorCode returns the value of the given color name
func colorCode(colorName string) int {
	colors := colors()

	return colors[colorName]
}

func fontColorCode(colorName string) int {
	return 30 + colorCode(colorName)
}

func bgColorCode(colorName string) int {
	return 40 + colorCode(colorName)
}

// Title outputs the string given underlined by equals
func Title(s string) {
	fmt.Fprintln(out, s)
	fmt.Fprintln(out, strings.Repeat("=", len(s)))
}

// Question reads input based on the question you pass to it
func Question(question string) (string, error) {
	reader := bufio.NewReader(in)
	fmt.Fprint(out, question+" \x1b[32m>\x1b[0m ")
	text, err := reader.ReadString('\n')
	text = strings.Replace(text, "\n", "", -1)
	return text, err
}

// Black outputs in ansi color black
func Black(s string) {
	PrintLine(fontColorCode("black"), s)
}

// Red outputs in ansi color red
func Red(s string) {
	PrintLine(fontColorCode("red"), s)
}

// Green outputs in ansi color green
func Green(s string) {
	PrintLine(fontColorCode("green"), s)
}

// Yellow outputs in ansi color yellow
func Yellow(s string) {
	PrintLine(fontColorCode("yellow"), s)
}

// Blue outputs in ansi color blue
func Blue(s string) {
	PrintLine(fontColorCode("blue"), s)
}

// Magenta outputs in ansi color magenta
func Magenta(s string) {
	PrintLine(fontColorCode("magenta"), s)
}

// Cyan outputs in ansi color cyan
func Cyan(s string) {
	PrintLine(fontColorCode("cyan"), s)
}

// White outputs in ansi color white
func White(s string) {
	PrintLine(fontColorCode("white"), s)
}

// AddStyle adds a new <tag> to be replaced with the <color> of choice
func AddStyle(tag string, color string) {
	tags[tag] = color
}

// WriteLn allows us to write a line to the console with markup in <tags> for colored output
func WriteLn(s string) {

	tags := styles()
	output := s

	for tag, color := range tags {
		output = strings.Replace(output, "<"+tag+">", "\x1b["+strconv.Itoa(fontColorCode(color))+"m", -1)
		output = strings.Replace(output, "</"+tag+">", "\x1b[0m", -1)
	}

	output = strings.Replace(output, "</>", "\x1b[0m", -1)

	fmt.Fprintln(out, output)
}

// Info prints green text
func Info(s string) {
	Green(s)
}

// Note prints cyan text
func Note(s string) {
	Cyan(s)
}

// Warning prints yellow text
func Warning(s string) {
	Yellow(s)
}

// Error prints red text
func Error(s string) {
	Red(s)
}

// PrintLine outputs in ansi color specified by the code
func PrintLine(code int, s string) {
	fmt.Fprint(out, "\x1b["+strconv.Itoa(code)+"m")
	fmt.Fprint(out, s)
	fmt.Fprintln(out, "\x1b[0m")
}

// BlackPanel outputs white text on a black background
func BlackPanel(s string) {
	PrintBox(bgColorCode("black"), fontColorCode("white"), s)
}

// RedPanel outputs white text on a red background
func RedPanel(s string) {
	PrintBox(bgColorCode("red"), fontColorCode("white"), s)
}

// GreenPanel outputs black text on a green background
func GreenPanel(s string) {
	PrintBox(bgColorCode("green"), fontColorCode("black"), s)
}

// YellowPanel outputs black text on a yellow background
func YellowPanel(s string) {
	PrintBox(bgColorCode("yellow"), fontColorCode("black"), s)
}

// BluePanel outputs black text on a blue background
func BluePanel(s string) {
	PrintBox(bgColorCode("blue"), fontColorCode("black"), s)
}

// MagentaPanel outputs white text on a magenta background
func MagentaPanel(s string) {
	PrintBox(bgColorCode("magenta"), fontColorCode("white"), s)
}

// CyanPanel outputs black text on a cyan background
func CyanPanel(s string) {
	PrintBox(bgColorCode("cyan"), fontColorCode("black"), s)
}

// WhitePanel outputs black text on a white background
func WhitePanel(s string) {
	PrintBox(bgColorCode("white"), fontColorCode("black"), s)
}

// PrintBox outputs colored text on a colored background
func PrintBox(backgroundCode int, fontCode int, s string) {
	fmt.Fprint(out, "\x1b["+strconv.Itoa(backgroundCode)+";"+strconv.Itoa(fontCode)+"m")
	fmt.Fprint(out, "\n\n")
	fmt.Fprintln(out, " "+s+" ")
	fmt.Fprintln(out, "\x1b[0m")
}
