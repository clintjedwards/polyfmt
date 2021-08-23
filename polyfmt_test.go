package polyfmt

func ExampleNewFormatter() {
	// Invoke the JSON printer. This is usually supplied by the user
	// at runtime. Ex. --json, --pretty, --silent
	pfmt, _ := NewFormatter(JSON)
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
	pfmt.Println("hello", Pretty)

	// Output: {"data":"hello","level":"info"}
	// {"data":{"test":"Some text"},"level":"info"}
}
