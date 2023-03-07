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

func (f *jsonFormatter) Print(msg any, filter ...Mode) {
	if isFiltered(JSON, filter) {
		return
	}

	tmp := map[string]any{
		"label": "info",
		"data":  msg,
	}
	b, _ := json.Marshal(&tmp)
	fmt.Println(string(b))
}

func (f *jsonFormatter) PrintErr(msg any, filter ...Mode) {
	if isFiltered(JSON, filter) {
		return
	}

	tmp := map[string]any{
		"label": "error",
		"data":  msg,
	}
	b, _ := json.Marshal(&tmp)
	fmt.Println(string(b))
}

func (f *jsonFormatter) PrintSuccess(msg any, filter ...Mode) {
	if isFiltered(JSON, filter) {
		return
	}

	tmp := map[string]any{
		"label": "success",
		"data":  msg,
	}
	b, _ := json.Marshal(&tmp)
	fmt.Println(string(b))
}

func (f *jsonFormatter) PrintWarning(msg any, filter ...Mode) {
	if isFiltered(JSON, filter) {
		return
	}

	tmp := map[string]any{
		"label": "warning",
		"data":  msg,
	}
	b, _ := json.Marshal(&tmp)
	fmt.Println(string(b))
}

func (f *jsonFormatter) Println(msg any, filter ...Mode) {
	if isFiltered(JSON, filter) {
		return
	}

	tmp := map[string]any{
		"label": "info",
		"data":  msg,
	}
	b, _ := json.Marshal(&tmp)
	fmt.Println(string(b))
}

func (f *jsonFormatter) Finish() {}
