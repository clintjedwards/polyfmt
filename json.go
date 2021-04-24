package polyfmt

import (
	"encoding/json"
	"fmt"
)

type jsonFormatter struct{}

func newJSONFormatter() (*jsonFormatter, error) {
	return &jsonFormatter{}, nil
}

func (f *jsonFormatter) Print(msg interface{}, filter ...Mode) {
	if isFiltered(JSON, filter) {
		return
	}

	tmp := map[string]interface{}{"info": msg}
	b, _ := json.Marshal(&tmp)
	fmt.Println(string(b))
}
func (f *jsonFormatter) PrintErr(msg interface{}, filter ...Mode) {
	if isFiltered(JSON, filter) {
		return
	}

	tmp := map[string]interface{}{"error": msg}
	b, _ := json.Marshal(&tmp)
	fmt.Println(string(b))
}
func (f *jsonFormatter) PrintSuccess(msg interface{}, filter ...Mode) {
	if isFiltered(JSON, filter) {
		return
	}

	tmp := map[string]interface{}{"success": msg}
	b, _ := json.Marshal(&tmp)
	fmt.Println(string(b))
}
func (f *jsonFormatter) Println(msg interface{}, filter ...Mode) {
	if isFiltered(JSON, filter) {
		return
	}

	tmp := map[string]interface{}{"info": msg}
	b, _ := json.Marshal(&tmp)
	fmt.Println(string(b))
}
func (f *jsonFormatter) Finish() {}
