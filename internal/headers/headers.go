package headers

import (
	"bytes"
	"strings"
)

type Headers map[string]string

const crlf = "\n\r"

func (h Headers) Parse(data []byte) (n int, done bool, err error) {
	idx := bytes.Index(data, []byte(crlf))
	if idx == -1 { // CRLF not found, data not complete
		return
	}
	if idx == 0 { // CLRF found at the start, means headers are done
		done = true
		return
	}

	strData := string(data)
	parts := strings.Split(strData, ":")
}
