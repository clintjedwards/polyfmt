package polyfmt

type jsonFormatter struct{}

func newJSONFormatter() (*jsonFormatter, error) {
	return &jsonFormatter{}, nil
}

func (f *jsonFormatter) Print(msg interface{}, filter ...Mode)        {}
func (f *jsonFormatter) PrintErr(msg interface{}, filter ...Mode)     {}
func (f *jsonFormatter) PrintSuccess(msg interface{}, filter ...Mode) {}
func (f *jsonFormatter) Println(msg interface{}, filter ...Mode)      {}
func (f *jsonFormatter) Finish()                                      {}
