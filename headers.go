package boots

import (
	"fmt"
	"strings"

	"github.com/marconi/boots/constants"
)

type Headers map[string]string

func (h Headers) Prepare(cmd string) string {
	var headers []string
	for name, value := range h {
		// if the command is not CONNECT/CONNECTED, escape header values.
		if cmd != constants.CONNECT && cmd != constants.CONNECTED {
			value = strings.Replace(value, "\r", "\\r", -1)
			value = strings.Replace(value, "\n", "\\n", -1)
			value = strings.Replace(value, ":", "\\c", -1)
		}
		header := fmt.Sprintf("%s:%s", name, value)
		headers = append(headers, header)
	}
	headers = append(headers, constants.EOL) // end of header
	return strings.Join(headers, constants.EOL)
}

// append extra headers
func (h Headers) Append(extraHead Headers) {
	for name, value := range extraHead {
		if _, ok := h[name]; !ok {
			h[name] = value
		}
	}
}
