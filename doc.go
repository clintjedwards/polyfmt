/*
Package polyfmt is a convenience package that provides multiple forms of formatted output.
Useful for CLI applications where you might want to provide JSON output for machine users,
but pretty output for interactive users.

# Why

In a command line application you usually want to provide some well-formatted output to users. This
may include progress bars, timers, spinners, or tables so that interactive users can better parse
your programs output. For non-interactive users or automation this might make your CLI application
difficult to parse, build automation around, or just unnecessarily verbose. To this end, you might want to
provide a common serialization format to users who use your CLI app from within a non-interactive environment.

Polyfmt aims to simplify the API around multiple formatting options and make it easy to switch between them.

# Usage

Polyfmt provides a very simple API, full of print functions.

Initiate a new formatter instance, passing in what type of formatter you want back. This is usually passed in
by your user at runtime via flags or config.

	pfmt, _ := polyfmt.NewFormatter(JSON, false)
	defer pfmt.Finish() // Finish flushes the output and cleans up safely.

Use the returned formatter to print a simple json formatted hello.

	pfmt.Println("hello")
	// Output:
	// {"label":"info","data":"hello"}

You can also pass the printer any interface and it will attempt to print it (providing that it is printable).

	pfmt.Println(struct {
	    Test string `json:"test"`
	}{
	    Test: "Some text",
	})
	// Output:
	// {"label":"info","data":{"test":"Some text"}}

Sometimes you'll want to output something only for specific formatters. Most commands take a list of formatters, which
tells the command to only print for those formatters.

	pfmt.Println("hello", Pretty)
	// Output:
	//

# Additional Details

You can turn off color by using the popular `NO_COLOR` environment variable.
*/
package polyfmt
