package request

import (
	"fmt"
	"io"
	"strings"
	"unicode"
)

type Request struct {
	RequestLine RequestLine
}

type RequestLine struct {
	HttpVersion   string
	RequestTarget string
	Method        string
}

func RequestFromReader(reader io.Reader) (*Request, error) {
	req, err := io.ReadAll(reader)
	if err != nil {
		return nil, fmt.Errorf("could not read from reader: %v", err)
	}

	requestLine, err := parseRequestLine(string(req))
	if err != nil {
		return nil, fmt.Errorf("could not parse request line: %v", err)
	}

	for _, letter := range requestLine.Method {
		if !unicode.IsUpper(letter) {
			return nil, fmt.Errorf("method line malformed, not all uppercase alphabetic")
		}
	}

	if !strings.Contains(requestLine.HttpVersion, "1.1") {
		return nil, fmt.Errorf("http version does not match")
	}

	request := Request{RequestLine: requestLine}
	return &request, nil
}

func parseRequestLine(request string) (RequestLine, error) {
	reqParts := strings.Split(request, "\r\n")
	lineParts := strings.Split(reqParts[0], " ")
	if len(lineParts) < 3 {
		return RequestLine{}, fmt.Errorf("request malformed, not enough fields")
	}

	httpVersionParts := strings.Split(lineParts[2], "/")
	if httpVersionParts[0] != "HTTP" {
		return RequestLine{}, fmt.Errorf("not an http request")
	}

	return RequestLine{
		HttpVersion:   httpVersionParts[1],
		RequestTarget: lineParts[1],
		Method:        lineParts[0],
	}, nil
}
