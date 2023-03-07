package polyfmt

import (
	"fmt"

	"github.com/fatih/color"
)

// plainFormatter wraps the passed in interface and prints it.
type plainFormatter struct{}

func newPlainFormatter() (*plainFormatter, error) {
	return &plainFormatter{}, nil
}

func (f *plainFormatter) Print(msg interface{}, filter ...Mode) {
	if isFiltered(Plain, filter) {
		return
	}

	fmt.Print(msg)
}

func (f *plainFormatter) PrintErr(msg interface{}, filter ...Mode) {
	if isFiltered(Plain, filter) {
		return
	}

	fmt.Printf("%s %s\n", color.RedString("x"), msg)
}

func (f *plainFormatter) PrintSuccess(msg interface{}, filter ...Mode) {
	if isFiltered(Plain, filter) {
		return
	}

	fmt.Printf("%s %s\n", color.GreenString("✓"), msg)
}

func (f *plainFormatter) PrintWarning(msg interface{}, filter ...Mode) {
	if isFiltered(Plain, filter) {
		return
	}

	fmt.Printf("%s %s\n", color.YellowString("!!"), msg)
}

func (f *plainFormatter) Println(msg interface{}, filter ...Mode) {
	if isFiltered(Plain, filter) {
		return
	}

	fmt.Println(msg)
}

func (f *plainFormatter) Finish() {}
