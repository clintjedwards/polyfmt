package polyfmt

// silentFormatter prints absolutely nothing and is essentially an implementation for --silent on
// command lines.
type silentFormatter struct{}

func newSilentFormatter() (*silentFormatter, error) {
	return &silentFormatter{}, nil
}

func (f *silentFormatter) Print(msg any, filter ...Mode)        {}
func (f *silentFormatter) PrintErr(msg any, filter ...Mode)     {}
func (f *silentFormatter) PrintSuccess(msg any, filter ...Mode) {}
func (f *silentFormatter) PrintWarning(msg any, filter ...Mode) {}
func (f *silentFormatter) PrintQuestion(msg any, filter ...Mode) string {
	return ""
}
func (f *silentFormatter) Println(msg any, filter ...Mode) {}
func (f *silentFormatter) Finish()                         {}
