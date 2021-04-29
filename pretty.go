package polyfmt

import (
	"fmt"
	"os"
	"time"

	"github.com/theckman/yacspin"
)

// prettyFormatter is a printer with a spinner attached.
//
// This current implementation has an auto suffix field that we ignore so we can simplify the API
// in order to provide a cleaner interface for other types of formatters.
type prettyFormatter struct {
	spinner        *yacspin.Spinner
	cfg            yacspin.Config
	currentMessage string // We store this so that we can restore it when we start a new spinner
}

func newPrettyFormatter() (*prettyFormatter, error) {
	cfg := yacspin.Config{
		Writer:            os.Stdout,
		Frequency:         100 * time.Millisecond,
		CharSet:           yacspin.CharSets[14],
		SuffixAutoColon:   false,
		StopCharacter:     "✓",
		StopColors:        []string{"fgGreen"},
		StopFailCharacter: "x",
		StopFailColors:    []string{"fgRed"},
	}

	f := &prettyFormatter{
		cfg: cfg,
	}

	err := f.newSpinner()
	if err != nil {
		return nil, err
	}

	return f, nil
}

func (f *prettyFormatter) newSpinner() error {
	spinner, err := yacspin.New(f.cfg)
	if err != nil {
		return err
	}
	err = spinner.Start()
	if err != nil {
		return err
	}
	spinner.Message(f.currentMessage)
	f.spinner = spinner
	return nil
}

func (f *prettyFormatter) Print(msg interface{}, filter ...Mode) {
	if isFiltered(Pretty, filter) {
		return
	}

	formattedMsg := fmt.Sprintf(" %s", msg)
	f.spinner.Message(formattedMsg)
	f.currentMessage = formattedMsg
}

func (f *prettyFormatter) Println(msg interface{}, filter ...Mode) {
	if isFiltered(Pretty, filter) {
		return
	}

	f.spinner.StopCharacter("")
	_ = f.spinner.Stop()
	fmt.Println(msg)
	_ = f.newSpinner()
}

func (f *prettyFormatter) PrintErr(msg interface{}, filter ...Mode) {
	if isFiltered(Pretty, filter) {
		return
	}

	f.spinner.StopFailMessage(fmt.Sprintf(" %s", msg))
	_ = f.spinner.StopFail()
	_ = f.newSpinner()
}

func (f *prettyFormatter) PrintSuccess(msg interface{}, filter ...Mode) {
	if isFiltered(Pretty, filter) {
		return
	}

	f.spinner.Suffix(fmt.Sprintf(" %s", msg))
	_ = f.spinner.Stop()
	_ = f.newSpinner()
}

func (f *prettyFormatter) Finish() {
	f.spinner.StopCharacter("")
	_ = f.spinner.Stop()
}
