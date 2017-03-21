package console

import (
	"bytes"
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

	if result != expected {
		t.Error(result, expected)
	}
}
