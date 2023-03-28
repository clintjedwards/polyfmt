package polyfmt

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

// plainFormatter wraps the passed in interface and prints it.
type plainFormatter struct {
	debug bool
}

func newPlainFormatter(debug bool) (*plainFormatter, error) {
	return &plainFormatter{
		debug: debug,
	}, nil
}

func (f *plainFormatter) Print(msg any, filter ...Mode) {
	if isFiltered(Plain, filter) {
		return
	}

	fmt.Print(msg)
}

func (f *plainFormatter) Err(msg any, filter ...Mode) {
	if isFiltered(Plain, filter) {
		return
	}

	fmt.Printf("%s %s\n", color.RedString("x"), msg)
}

func (f *plainFormatter) Success(msg any, filter ...Mode) {
	if isFiltered(Plain, filter) {
		return
	}

	fmt.Printf("%s %s\n", color.GreenString("✓"), msg)
}

func (f *plainFormatter) Warning(msg any, filter ...Mode) {
	if isFiltered(Plain, filter) {
		return
	}

	fmt.Printf("%s %s\n", color.YellowString("!!"), msg)
}

func (f *plainFormatter) Question(msg any, filter ...Mode) string {
	if isFiltered(Plain, filter) {
		return ""
	}

	var input string
	fmt.Printf("%s %s", color.MagentaString("?"), msg)
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input = scanner.Text()
	}

	return input
}

func (f *plainFormatter) Debugln(msg any, filter ...Mode) {
	if isFiltered(Plain, filter) || !f.debug {
		return
	}

	c := color.New(color.BgYellow).Add(color.Faint)

	fmt.Printf("%s %s\n", c.Sprint("DEBUG"), msg)
}

func (f *plainFormatter) Println(msg any, filter ...Mode) {
	if isFiltered(Plain, filter) {
		return
	}

	fmt.Println(msg)
}

func (f *plainFormatter) Finish() {}
