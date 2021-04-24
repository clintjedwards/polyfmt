package polyfmt

type silentFormatter struct{}

func newSilentFormatter() (*silentFormatter, error) {
	return &silentFormatter{}, nil
}

func (f *silentFormatter) Print(msg interface{}, filter ...Mode)        {}
func (f *silentFormatter) PrintErr(msg interface{}, filter ...Mode)     {}
func (f *silentFormatter) PrintSuccess(msg interface{}, filter ...Mode) {}
func (f *silentFormatter) Println(msg interface{}, filter ...Mode)      {}
func (f *silentFormatter) Finish()                                      {}
