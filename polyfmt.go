package polyfmt

import (
	"fmt"
	"os"
)

type Mode string

const (
	// Plain outputs text in a humanized fashion without spinners.
	Plain Mode = "plain"
	// Pretty outputs text in a more humanized fashion and provides spinners for longer actions.
	Pretty Mode = "pretty"
	// JSON outputs json formatted text, mainly suitable to be read by computers.
	JSON Mode = "json"
	// Dummy formatter that doesn't print anything
	Silent Mode = "silent"
)

type Formatter interface {
	// Print will attempt to intelligently print objects passed to it.
	// Adding modes to the filter restricts the object being printed only
	// to those modes.
	Print(msg any, filter ...Mode)

	// PrintErr prints the message noting it as an error to the user.
	// Adding modes to the filter restricts the object being printed only
	// to those modes.
	PrintErr(msg any, filter ...Mode)

	// PrintSuccess prints the message noting it as an error to the user.
	// Adding modes to the filter restricts the object being printed only
	// to those modes.
	PrintSuccess(msg any, filter ...Mode)

	// PrintWarning prints the message noting it as a warning to the user.
	// Adding modes to the filter restricts the object being printed only
	// to those modes.
	PrintWarning(msg any, filter ...Mode)

	// PrintQuestion prints the message noting it as a question to the user.
	// It also collects user input using bufio.Scanner and returns it.
	//
	// Adding modes to the filter restricts the object being printed only
	// to those modes. This is especially important for this mode,
	// since even in JSON output it will stop and wait for user input.
	//
	// If filtered will return an empty string.
	PrintQuestion(msg any, filter ...Mode) string

	// Println prints the message adding a newline to the end.
	// Adding modes to the filter restricts the object being printed only
	// to those modes.
	Println(msg any, filter ...Mode)

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

// isFiltered is a convenience function for detecting if the current mode is in the list of modes to print
func isFiltered(currMode Mode, filterList []Mode) bool {
	if len(filterList) < 1 {
		return false
	}

	for _, mode := range filterList {
		if mode == currMode {
			return false
		}
	}

	return true
}

// NewFormatter create a new formatter with the appropriate mode.
// detectNonInteractive allows the user to opt-in to the ability to auto switch to JSON output
// if we detect that the output is being piped into a non-interactive context.
// (as in the case of piping to another command)
func NewFormatter(mode Mode, detectNonInteractive bool) (Formatter, error) {
	if mode == Pretty && !isTTY() && detectNonInteractive {
		mode = JSON
	}

	switch mode {
	case Plain:
		f, err := newPlainFormatter()
		if err != nil {
			return nil, err
		}
		return f, nil
	case Pretty:
		f, err := newPrettyFormatter()
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
	case Silent:
		f, err := newSilentFormatter()
		if err != nil {
			return nil, err
		}
		return f, nil
	default:
		return nil, fmt.Errorf("invalid formatter: %q", mode)
	}
}
