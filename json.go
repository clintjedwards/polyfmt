package polyfmt

import (
	"encoding/json"
	"fmt"
)

// jsonFormatter wraps the passed in interface and prints it.
type jsonFormatter struct{}

func newJSONFormatter() (*jsonFormatter, error) {
	return &jsonFormatter{}, nil
}

func (f *jsonFormatter) Print(msg interface{}, filter ...Mode) {
	if isFiltered(JSON, filter) {
		return
	}

	tmp := map[string]interface{}{
		"label": "info",
		"data":  msg,
	}
	b, _ := json.Marshal(&tmp)
	fmt.Println(string(b))
}

func (f *jsonFormatter) PrintErr(msg interface{}, filter ...Mode) {
	if isFiltered(JSON, filter) {
		return
	}

	tmp := map[string]interface{}{
		"label": "error",
		"data":  msg,
	}
	b, _ := json.Marshal(&tmp)
	fmt.Println(string(b))
}

func (f *jsonFormatter) PrintSuccess(msg interface{}, filter ...Mode) {
	if isFiltered(JSON, filter) {
		return
	}

	tmp := map[string]interface{}{
		"label": "success",
		"data":  msg,
	}
	b, _ := json.Marshal(&tmp)
	fmt.Println(string(b))
}

func (f *jsonFormatter) Println(msg interface{}, filter ...Mode) {
	if isFiltered(JSON, filter) {
		return
	}

	tmp := map[string]interface{}{
		"label": "info",
		"data":  msg,
	}
	b, _ := json.Marshal(&tmp)
	fmt.Println(string(b))
}

func (f *jsonFormatter) Finish() {}
