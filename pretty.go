package polyfmt

type prettyFormatter struct{}

func newPrettyFormatter() *prettyFormatter {
	return &prettyFormatter{}
}

func (f *prettyFormatter) Print(msg interface{}, filter []Mode) {}
