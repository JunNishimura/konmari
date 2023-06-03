package extension

import (
	"strings"
)

type Mark string

const (
	Undefined Mark = "undefined"
	Go        Mark = "go"
	Python    Mark = "python"
)

var (
	extensionMap = map[string]Mark{
		"go": Go,
		"py": Python,
	}
)

func New(filename string) Mark {
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
