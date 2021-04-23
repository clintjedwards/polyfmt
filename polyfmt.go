package polyfmt

type Mode string

const (
	Pretty Mode = "pretty"
	Plain  Mode = "plain"
	JSON   Mode = "json"
)

type Formatter interface {
	// Print will attempt to intelligently print objects passed to it.
	// Adding modes to the filter restricts the object being printed only
	// to those modes.
	Print(msg interface{}, filter []Mode)
}

func NewFormatter(mode Mode) (Formatter, error) {

	switch mode {
	case Pretty:
		return newPrettyFormatter(), nil
	case Plain:
		return newPlainFormatter(), nil
	case JSON:
		return newJSONFormatter(), nil
	}
	return nil, nil
}
