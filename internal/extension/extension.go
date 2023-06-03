package extension

import (
	"strings"
)

type extension string

const (
	Undefined extension = "undefined"
	Go        extension = "go"
)

var (
	extensionMap = map[string]extension{
		"go": Go,
	}
)

func New(filename string) extension {
	filenameSplit := strings.Split(filename, ".")
	if len(filenameSplit) <= 1 {
		return Undefined
	}
	extension := filenameSplit[len(filenameSplit)-1]
	normExtension := strings.ToLower(extension)
	got, ok := extensionMap[normExtension]
	if ok {
		return got
	} else {
		return Undefined
	}
}
