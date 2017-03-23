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

// Colors returns a map of colors that we can use to work out color codes for the app
func Colors() map[string]int {

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

// ColorCode returns the value of the given color name
func ColorCode(colorName string) int {
	colors := Colors()

	return colors[colorName]
}

// FontColorCode returns the value used for styling console fonts of the given color name
func FontColorCode(colorName string) int {
	return 30 + ColorCode(colorName)
}

// BgColorCode returns the value used for styling console fonts of the given color name
func BgColorCode(colorName string) int {
	return 40 + ColorCode(colorName)
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
	PrintLine(FontColorCode("black"), s)
}

// Red outputs in ansi color red
func Red(s string) {
	PrintLine(FontColorCode("red"), s)
}

// Green outputs in ansi color green
func Green(s string) {
	PrintLine(FontColorCode("green"), s)
}

// Yellow outputs in ansi color yellow
func Yellow(s string) {
	PrintLine(FontColorCode("yellow"), s)
}

// Blue outputs in ansi color blue
func Blue(s string) {
	PrintLine(FontColorCode("blue"), s)
}

// Magenta outputs in ansi color magenta
func Magenta(s string) {
	PrintLine(FontColorCode("magenta"), s)
}

// Cyan outputs in ansi color cyan
func Cyan(s string) {
	PrintLine(FontColorCode("cyan"), s)
}

// White outputs in ansi color white
func White(s string) {
	PrintLine(FontColorCode("white"), s)
}

func Styles() map[string]string {
	tags["info"] = "green"
	tags["note"] = "blue"
	tags["comment"] = "cyan"
	tags["warning"] = "yellow"
	tags["error"] = "red"
	return tags
}

func AddStyle(tag string, color string) {
	tags[tag] = color
}

func WriteLn(s string) {

	tags := Styles()

	output := s
	for tag, color := range tags {
		output = strings.Replace(output, "<"+tag+">", "\x1b["+strconv.Itoa(FontColorCode(color))+"m", -1)
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
	PrintBox(BgColorCode("black"), FontColorCode("white"), s)
}

// RedPanel outputs white text on a red background
func RedPanel(s string) {
	PrintBox(BgColorCode("red"), FontColorCode("white"), s)
}

// GreenPanel outputs black text on a green background
func GreenPanel(s string) {
	PrintBox(BgColorCode("green"), FontColorCode("black"), s)
}

// YellowPanel outputs black text on a yellow background
func YellowPanel(s string) {
	PrintBox(BgColorCode("yellow"), FontColorCode("black"), s)
}

// BluePanel outputs black text on a blue background
func BluePanel(s string) {
	PrintBox(BgColorCode("blue"), FontColorCode("black"), s)
}

// MagentaPanel outputs white text on a magenta background
func MagentaPanel(s string) {
	PrintBox(BgColorCode("magenta"), FontColorCode("white"), s)
}

// CyanPanel outputs black text on a cyan background
func CyanPanel(s string) {
	PrintBox(BgColorCode("cyan"), FontColorCode("black"), s)
}

// WhitePanel outputs black text on a white background
func WhitePanel(s string) {
	PrintBox(BgColorCode("white"), FontColorCode("black"), s)
}

// PrintBox outputs colored text on a colored background
func PrintBox(backgroundCode int, fontCode int, s string) {
	fmt.Fprint(out, "\x1b["+strconv.Itoa(backgroundCode)+";"+strconv.Itoa(fontCode)+"m")
	fmt.Fprintln(out, "\n")
	fmt.Fprintln(out, " "+s+" ")
	fmt.Fprintln(out, "\x1b[0m")
}
