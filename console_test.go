package console

import (
	"bytes"
	"fmt"
	"testing"
)

func TestItPrintsInBlack(t *testing.T) {
	// assemble
	buf := &bytes.Buffer{}
	out = buf

	// act
	Black("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[30mmy string goes here\x1b[0m\n"

	fmt.Println(result)

	if result != expected {
		t.Error(result, expected)
	}
}

func TestItPrintsInRed(t *testing.T) {
	// assemble
	buf := &bytes.Buffer{}
	out = buf

	// act
	Red("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[31mmy string goes here\x1b[0m\n"

	fmt.Println(result)

	if result != expected {
		t.Error(result, expected)
	}
}

func TestItPrintsInGreen(t *testing.T) {
	// assemble
	buf := &bytes.Buffer{}
	out = buf

	// act
	Green("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[32mmy string goes here\x1b[0m\n"

	fmt.Println(result)

	if result != expected {
		t.Error(result, expected)
	}
}

func TestItPrintsInYellow(t *testing.T) {
	// assemble
	buf := &bytes.Buffer{}
	out = buf

	// act
	Yellow("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[33mmy string goes here\x1b[0m\n"

	fmt.Println(result)

	if result != expected {
		t.Error(result, expected)
	}
}

func TestItPrintsInBlue(t *testing.T) {
	// assemble
	buf := &bytes.Buffer{}
	out = buf

	// act
	Blue("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[34mmy string goes here\x1b[0m\n"

	fmt.Println(result)

	if result != expected {
		t.Error(result, expected)
	}
}

func TestItPrintsInMagenta(t *testing.T) {
	// assemble
	buf := &bytes.Buffer{}
	out = buf

	// act
	Magenta("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[35mmy string goes here\x1b[0m\n"

	fmt.Println(result)

	if result != expected {
		t.Error(result, expected)
	}
}

func TestItPrintsInCyan(t *testing.T) {
	// assemble
	buf := &bytes.Buffer{}
	out = buf

	// act
	Cyan("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[36mmy string goes here\x1b[0m\n"

	fmt.Println(result)

	if result != expected {
		t.Error(result, expected)
	}
}

func TestItPrintsInWhite(t *testing.T) {
	// assemble
	buf := &bytes.Buffer{}
	out = buf

	// act
	White("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[37mmy string goes here\x1b[0m\n"

	fmt.Println(result)

	if result != expected {
		t.Error(result, expected)
	}
}

func TestItPrintsABlackBox(t *testing.T) {
	// assemble
	buf := &bytes.Buffer{}
	out = buf

	// act
	BlackBox("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[40;37m my string goes here \x1b[0m\n"

	fmt.Println(result)

	if result != expected {
		t.Error(result, expected)
	}
}

func TestItPrintsARedBox(t *testing.T) {
	// assemble
	buf := &bytes.Buffer{}
	out = buf

	// act
	RedBox("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[41;37m my string goes here \x1b[0m\n"

	fmt.Println(result)

	if result != expected {
		t.Error(result, expected)
	}
}

func TestItPrintsAGreenBox(t *testing.T) {
	// assemble
	buf := &bytes.Buffer{}
	out = buf

	// act
	GreenBox("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[42;30m my string goes here \x1b[0m\n"

	fmt.Println(result)

	if result != expected {
		t.Error(result, expected)
	}
}

func TestItPrintsAYellowBox(t *testing.T) {
	// assemble
	buf := &bytes.Buffer{}
	out = buf

	// act
	YellowBox("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[43;30m my string goes here \x1b[0m\n"

	fmt.Println(result)

	if result != expected {
		t.Error(result, expected)
	}
}

func TestItPrintsABlueBox(t *testing.T) {
	// assemble
	buf := &bytes.Buffer{}
	out = buf

	// act
	BlueBox("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[44;30m my string goes here \x1b[0m\n"

	fmt.Println(result)

	if result != expected {
		t.Error(result, expected)
	}
}

func TestItPrintsAMagentaBox(t *testing.T) {
	// assemble
	buf := &bytes.Buffer{}
	out = buf

	// act
	MagentaBox("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[45;37m my string goes here \x1b[0m\n"

	fmt.Println(result)

	if result != expected {
		t.Error(result, expected)
	}
}

func TestItPrintsACyanBox(t *testing.T) {
	// assemble
	buf := &bytes.Buffer{}
	out = buf

	// act
	CyanBox("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[46;30m my string goes here \x1b[0m\n"

	fmt.Println(result)

	if result != expected {
		t.Error(result, expected)
	}
}

func TestItPrintsAWhiteBox(t *testing.T) {
	// assemble
	buf := &bytes.Buffer{}
	out = buf

	// act
	WhiteBox("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[47;30m my string goes here \x1b[0m\n"

	fmt.Println(result)

	if result != expected {
		t.Error(result, expected)
	}
}
