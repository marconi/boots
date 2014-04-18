package boots

import (
	"errors"
	"fmt"
	"strings"

	"github.com/marconi/boots/constants"
)

type Frame struct {
	Command string
	Headers Headers
	Body    string
}

func NewFrame(cmd string, h Headers, body string) *Frame {
	return &Frame{
		Command: cmd,
		Headers: h,
		Body:    body,
	}
}

func (f *Frame) Build() string {
	return f.Command +
		constants.EOL +
		f.Headers.Prepare(f.Command) +
		f.Body +
		string(constants.NULL)
}

func ParseFrame(rawFrame string) (*Frame, error) {
	lines := strings.Split(rawFrame, constants.EOL)

	cmd := lines[0]
	if _, ok := constants.COMMANDS[cmd]; !ok {
		return nil, errors.New(fmt.Sprintf("Invalid command: %", cmd))
	}

	// look for the ending index of header and body lines
	var headerEndIndex, bodyEndIndex int
	for i, line := range lines {
		if line == "" {
			headerEndIndex = i
		} else if line == string(constants.NULL) {
			bodyEndIndex = i
		}
	}

	// build headers
	headers := make(Headers)
	rawHeaders := lines[1:headerEndIndex]
	for _, header := range rawHeaders {
		h := strings.Split(header, ":")
		name, value := h[0], h[1]

		// if the command is not CONNECT/CONNECTED, un-escape header values.
		if cmd != constants.CONNECT && cmd != constants.CONNECTED {
			value = strings.Replace(value, "\\r", "\r", -1)
			value = strings.Replace(value, "\\n", "\n", -1)
			value = strings.Replace(value, "\\c", ":", -1)
		}

		headers[name] = value
	}

	body := strings.Join(lines[headerEndIndex+1:bodyEndIndex], "\n")

	return NewFrame(cmd, headers, body), nil
}
