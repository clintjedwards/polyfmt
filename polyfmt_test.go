package polyfmt_test

import (
	"github.com/clintjedwards/polyfmt/v2"
)

func ExampleNewFormatter() {
	// Invoke the JSON printer. This is usually supplied by the user
	// at runtime. Ex. --json, --pretty, --silent
	pfmt, _ := polyfmt.NewFormatter(polyfmt.JSON, polyfmt.DefaultOptions())
	defer pfmt.Finish() // Finish flushes the output and cleans up safely.

	// Prints a simple json formatted hello.
	pfmt.Println("hello")

	// Prints the data structure passed to it if it can be displayed.
	pfmt.Println(struct {
		Test string `json:"test"`
	}{
		Test: "Some text",
	})

	// Most commands also have the ability to not print under specific settings.
	// This example prints hello ONLY if the "Pretty" mode is on, if any other mode
	// is enabled it will be skipped.
	pfmt.Println("hello", polyfmt.Pretty)

	// Output: {"data":"hello","label":"info"}
	// {"data":{"test":"Some text"},"label":"info"}
}

func ExampleFormatter_Debugln() {
	// Sometimes you'll want your tooling to print debug messages if needed.
	// The debug option allows you to write those debug lines normally and
	// only show them when you want to.
	options := polyfmt.DefaultOptions()
	debug := true
	options.Debug = &debug
	pfmt, _ := polyfmt.NewFormatter(polyfmt.JSON, options)
	defer pfmt.Finish() // Finish flushes the output and cleans up safely.

	// Prints a normal JSON formatted hello.
	pfmt.Println("hello")

	// Prints a debug line
	pfmt.Debugln("hello!")

	// Output: {"data":"hello","label":"info"}
	// {"data":"hello!","label":"debug"}
}
