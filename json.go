package polyfmt

type jsonFormatter struct{}

func newJSONFormatter() *jsonFormatter {
	return &jsonFormatter{}
}

func (f *jsonFormatter) Print(msg interface{}, filter []Mode) {}
