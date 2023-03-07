package polyfmt

// silentFormatter prints absolutely nothing and is essentially an implementation for --silent on
// command lines.
type silentFormatter struct{}

func newSilentFormatter() (*silentFormatter, error) {
	return &silentFormatter{}, nil
}

func (f *silentFormatter) Print(msg interface{}, filter ...Mode)        {}
func (f *silentFormatter) PrintErr(msg interface{}, filter ...Mode)     {}
func (f *silentFormatter) PrintSuccess(msg interface{}, filter ...Mode) {}
func (f *silentFormatter) PrintWarning(msg interface{}, filter ...Mode) {}
func (f *silentFormatter) Println(msg interface{}, filter ...Mode)      {}
func (f *silentFormatter) Finish()                                      {}
