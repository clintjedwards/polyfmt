package polyfmt

import (
	"fmt"
	"os"

	"github.com/fatih/color"
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

type Options struct {
	// Allows the user to opt-in to the ability for polyfmt to detect non interactive terminals and
	// auto switch to JSON output. Default is false.
	AutoDetectTTY *bool

	// Turn on printing for debug lines
	Debug *bool
}

func DefaultOptions() Options {
	return Options{
		AutoDetectTTY: ptr(false),
		Debug:         ptr(false),
	}
}

type Formatter interface {
	// Print will attempt to intelligently print objects passed to it.
	// Adding modes to the filter restricts the object being printed only
	// to those modes.
	Print(msg any, filter ...Mode)

	// Println prints the message adding a newline to the end.
	// Adding modes to the filter restricts the object being printed only
	// to those modes.
	Println(msg any, filter ...Mode)

	// Err prints the message noting it as an error to the user.
	// Adding modes to the filter restricts the object being printed only
	// to those modes.
	Err(msg any, filter ...Mode)

	// Success prints the message noting it as an error to the user.
	// Adding modes to the filter restricts the object being printed only
	// to those modes.
	Success(msg any, filter ...Mode)

	// Warning prints the message noting it as a warning to the user.
	// Adding modes to the filter restricts the object being printed only
	// to those modes.
	Warning(msg any, filter ...Mode)

	// Debugln prints a message only if debug is turned on in the formatter options.
	// Newline included.
	Debugln(msg any, filter ...Mode)

	// Question prints the message noting it as a question to the user.
	// It also collects user input using bufio.Scanner and returns it.
	//
	// Adding modes to the filter restricts the object being printed only
	// to those modes. This is especially important for this mode,
	// since even in JSON output it will stop and wait for user input.
	//
	// If filtered will return an empty string.
	Question(msg any, filter ...Mode) string

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
// if we detect that the output is being piped into a non-interactive context.
// (as in the case of piping to another command)
func NewFormatter(mode Mode, options Options) (Formatter, error) {
	opts := DefaultOptions()

	if options.AutoDetectTTY != nil {
		opts.AutoDetectTTY = options.AutoDetectTTY
	}

	// The pretty mode does not print well when not in an interactive terminal. This check is
	// here mostly to cover situations where the user has forgotten the application is in
	// pretty mode.
	if mode == Pretty && !isTTY() && *opts.AutoDetectTTY {
		mode = JSON
	}

	// Explicitly set noColor since the backing library attempts to detect NonInteractive ttys and turns color off
	// automagically. This is a problem for non-interactive environments that can actually support color.
	noColor := os.Getenv("NO_COLOR")
	if noColor != "" {
		color.NoColor = true
	} else {
		color.NoColor = false
	}

	switch mode {
	case Plain:
		f, err := newPlainFormatter(*options.Debug)
		if err != nil {
			return nil, err
		}
		return f, nil
	case Pretty:
		f, err := newPrettyFormatter(*options.Debug)
		if err != nil {
			return nil, err
		}
		return f, nil
	case JSON:
		f, err := newJSONFormatter(*options.Debug)
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

func ptr[T any](v T) *T {
	return &v
}
