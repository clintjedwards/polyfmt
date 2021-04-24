package polyfmt

import "os"

type Mode string
type FormatterOption func()

const (
	// Pretty outputs text in a more humanized fashion and provides spinners for longer actions.
	Pretty Mode = "pretty"
	// Plain outputs text as pretty printed json.
	Plain Mode = "plain"
	// JSON outputs json formatted text, mainly suitable to be read by computers.
	JSON Mode = "json"
)

type Formatter interface {
	// Print will attempt to intelligently print objects passed to it.
	// Adding modes to the filter restricts the object being printed only
	// to those modes.
	Print(msg interface{}, filter ...Mode)
	// PrintErr prints the message noting it as an error to the user.
	// Adding modes to the filter restricts the object being printed only
	// to those modes.
	PrintErr(msg interface{}, filter ...Mode)
	// PrintSuccess prints the message noting it as an error to the user.
	// Adding modes to the filter restricts the object being printed only
	// to those modes.
	PrintSuccess(msg interface{}, filter ...Mode)
	// Println prints the message adding a newline to the end.
	// Adding modes to the filter restricts the object being printed only
	// to those modes.
	Println(msg interface{}, filter ...Mode)
	// Cleans up and flushes any last bit of formatting.
	// Should be called as the before program exit.
	Finish()
}

// isTTY determines if program is being run from terminal
func isTTY() bool {
	if fileInfo, _ := os.Stdout.Stat(); (fileInfo.Mode() & os.ModeCharDevice) != 0 {
		return true
	}

	return false
}

// NewFormatter create a new formatter with the appropriate mode. If mode is pretty and the environment
// this is run in is not interactive, it will intelligently revert to plain mode.
func NewFormatter(mode Mode) (Formatter, error) {
	if mode == Pretty && !isTTY() {
		mode = Plain
	}

	switch mode {
	case Pretty:
		f, err := newPrettyFormatter()
		if err != nil {
			return nil, err
		}
		return f, nil
	case Plain:
		f, err := newPlainFormatter()
		if err != nil {
			return nil, err
		}
		return f, nil
	case JSON:
		f, err := newJSONFormatter()
		if err != nil {
			return nil, err
		}
		return f, nil
	}
	return nil, nil
}
