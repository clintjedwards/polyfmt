package polyfmt

// silentFormatter prints absolutely nothing and is essentially an implementation for --silent on
// command lines.
type silentFormatter struct{}

func newSilentFormatter() (*silentFormatter, error) {
	return &silentFormatter{}, nil
}

func (f *silentFormatter) Print(msg any, filter ...Mode)   {}
func (f *silentFormatter) Err(msg any, filter ...Mode)     {}
func (f *silentFormatter) Success(msg any, filter ...Mode) {}
func (f *silentFormatter) Warning(msg any, filter ...Mode) {}
func (f *silentFormatter) Question(msg any, filter ...Mode) string {
	return ""
}
func (f *silentFormatter) Println(msg any, filter ...Mode) {}
func (f *silentFormatter) Debugln(msg any, filter ...Mode) {}
func (f *silentFormatter) Finish()                         {}
