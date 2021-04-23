package polyfmt

type plainFormatter struct{}

func newPlainFormatter() *plainFormatter {
	return &plainFormatter{}
}

func (f *plainFormatter) Print(msg interface{}, filter []Mode) {}
