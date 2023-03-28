package polyfmt

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// jsonFormatter wraps the passed in interface and prints it.
type jsonFormatter struct {
	debug bool
}

func newJSONFormatter(debug bool) (*jsonFormatter, error) {
	return &jsonFormatter{
		debug: debug,
	}, nil
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

func (f *jsonFormatter) Err(msg any, filter ...Mode) {
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

func (f *jsonFormatter) Success(msg any, filter ...Mode) {
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

func (f *jsonFormatter) Warning(msg any, filter ...Mode) {
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

func (f *jsonFormatter) Question(msg any, filter ...Mode) string {
	if isFiltered(JSON, filter) {
		return ""
	}

	tmp := map[string]any{
		"label": "question",
		"data":  msg,
	}
	b, _ := json.Marshal(&tmp)

	var input string
	fmt.Println(string(b))
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		input = scanner.Text()
	}

	return input
}

func (f *jsonFormatter) Debugln(msg any, filter ...Mode) {
	if isFiltered(JSON, filter) || !f.debug {
		return
	}

	tmp := map[string]any{
		"label": "debug",
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
