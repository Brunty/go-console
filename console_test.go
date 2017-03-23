package console

import (
	"bytes"
	"fmt"
	"testing"
)

func TestItGetsInputFromAQuestion(t *testing.T) {
	// assemble
	buf := &bytes.Buffer{}
	out = buf

	// act
	result, _ := Question("Where would you like to go today?")

	// assert
	expected := ""

	if result != expected {
		t.Error(result, expected)
	}
}

func TestItPrintsATitle(t *testing.T) {
	// assemble
	buf := &bytes.Buffer{}
	out = buf

	// act
	Title("My string goes here")

	// assert
	result := buf.String()
	expected := "My string goes here\n===================\n"

	fmt.Print(result)

	if result != expected {
		t.Error(result, expected)
	}
}

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

func TestItPrintsColorFromTags(t *testing.T) {
	// assemble
	buf := &bytes.Buffer{}
	out = buf

	// act
	WriteLn("my <info>string</info> <error>goes</error> here")

	// assert
	result := buf.String()
	expected := "my \x1b[32mstring\x1b[0m \x1b[31mgoes\x1b[0m here\n"

	fmt.Println(result)

	if result != expected {
		t.Error(result, expected)
	}
}

func TestItPrintsColorFromCustomTags(t *testing.T) {
	// assemble
	buf := &bytes.Buffer{}
	out = buf

	// act
	AddStyle("mytag", "red")
	WriteLn("my <mytag>string</mytag> goes here")

	// assert
	result := buf.String()
	expected := "my \x1b[31mstring\x1b[0m goes here\n"

	fmt.Println(result)

	if result != expected {
		t.Error(result, expected)
	}
}

func TestItPrintsInfoText(t *testing.T) {
	// assemble
	buf := &bytes.Buffer{}
	out = buf

	// act
	Info("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[32mmy string goes here\x1b[0m\n"

	fmt.Println(result)

	if result != expected {
		t.Error(result, expected)
	}
}

func TestItPrintsNoteText(t *testing.T) {
	// assemble
	buf := &bytes.Buffer{}
	out = buf

	// act
	Note("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[36mmy string goes here\x1b[0m\n"

	fmt.Println(result)

	if result != expected {
		t.Error(result, expected)
	}
}

func TestItPrintsWarningText(t *testing.T) {
	// assemble
	buf := &bytes.Buffer{}
	out = buf

	// act
	Warning("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[33mmy string goes here\x1b[0m\n"

	fmt.Println(result)

	if result != expected {
		t.Error(result, expected)
	}
}

func TestItPrintsErrorText(t *testing.T) {
	// assemble
	buf := &bytes.Buffer{}
	out = buf

	// act
	Error("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[31mmy string goes here\x1b[0m\n"

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
	BlackPanel("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[40;37m\n\n my string goes here \n\x1b[0m\n"

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
	RedPanel("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[41;37m\n\n my string goes here \n\x1b[0m\n"

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
	GreenPanel("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[42;30m\n\n my string goes here \n\x1b[0m\n"

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
	YellowPanel("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[43;30m\n\n my string goes here \n\x1b[0m\n"

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
	BluePanel("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[44;30m\n\n my string goes here \n\x1b[0m\n"

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
	MagentaPanel("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[45;37m\n\n my string goes here \n\x1b[0m\n"

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
	CyanPanel("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[46;30m\n\n my string goes here \n\x1b[0m\n"

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
	WhitePanel("my string goes here")

	// assert
	result := buf.String()
	expected := "\x1b[47;30m\n\n my string goes here \n\x1b[0m\n"

	fmt.Println(result)

	if result != expected {
		t.Error(result, expected)
	}
}
