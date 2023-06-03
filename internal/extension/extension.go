package extension

import (
	"strings"

	"golang.org/x/exp/slices"
)

const (
	GO = "go"
)

var (
	extensions = []string{
		GO,
	}
)

func IsAcceptible(filename string) bool {
	filenameSplit := strings.Split(filename, ".")
	if len(filenameSplit) <= 1 {
		return false
	}
	extension := filenameSplit[len(filenameSplit)-1]
	normExtension := strings.ToLower(extension)
	return slices.Contains(extensions, normExtension)
}
