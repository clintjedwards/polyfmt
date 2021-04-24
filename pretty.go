package polyfmt

import (
	"fmt"
	"os"
	"time"

	"github.com/theckman/yacspin"
)

type prettyFormatter struct {
	spinner *yacspin.Spinner
	cfg     yacspin.Config
}

func newPrettyFormatter() (*prettyFormatter, error) {
	cfg := yacspin.Config{
		Writer:            os.Stdout,
		Frequency:         100 * time.Millisecond,
		CharSet:           yacspin.CharSets[14],
		Suffix:            " ",
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
	f.spinner = spinner
	return nil
}

func (f *prettyFormatter) Print(msg interface{}, filter ...Mode) {
	f.spinner.Message(fmt.Sprintf("%s", msg))
}

func (f *prettyFormatter) PrintErr(msg interface{}, filter ...Mode) {
	f.spinner.StopFailMessage(fmt.Sprintf("%s", msg))
	_ = f.spinner.StopFail()
	_ = f.newSpinner()
}

func (f *prettyFormatter) PrintSuccess(msg interface{}, filter ...Mode) {
	f.spinner.Suffix(fmt.Sprintf(" %s", msg))
	_ = f.spinner.Stop()
	f.spinner.Suffix(" ")
	_ = f.newSpinner()
}

func (f *prettyFormatter) Println(msg interface{}, filter ...Mode) {
	f.spinner.StopCharacter("")
	_ = f.spinner.Stop()
	fmt.Println(msg)
	_ = f.newSpinner()
}

func (f *prettyFormatter) Finish() {
	f.spinner.StopCharacter("")
	_ = f.spinner.Stop()
}
