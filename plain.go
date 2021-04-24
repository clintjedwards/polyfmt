package polyfmt

type plainFormatter struct{}

func newPlainFormatter() (*plainFormatter, error) {
	return &plainFormatter{}, nil
}

func (f *plainFormatter) Print(msg interface{}, filter ...Mode)        {}
func (f *plainFormatter) PrintErr(msg interface{}, filter ...Mode)     {}
func (f *plainFormatter) PrintSuccess(msg interface{}, filter ...Mode) {}
func (f *plainFormatter) Println(msg interface{}, filter ...Mode)      {}
func (f *plainFormatter) Finish()                                      {}
