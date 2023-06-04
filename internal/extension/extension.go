package extension

import (
	"strings"
)

type Mark string

const (
	Undefined  Mark = "undefined"
	Go         Mark = "go"
	Python     Mark = "python"
	JavaScript Mark = "javascript"
	TypeScript Mark = "typescript"
	Ruby       Mark = "ruby"
	PHP        Mark = "php"
	Java       Mark = "java"
	C          Mark = "c"
	Cpp        Mark = "cpp"
	Rust       Mark = "rust"
	Swift      Mark = "swift"
	Kotlin     Mark = "kotlin"
)

var (
	extensionMap = map[string]Mark{
		"go":    Go,
		"py":    Python,
		"js":    JavaScript,
		"ts":    TypeScript,
		"rb":    Ruby,
		"php":   PHP,
		"java":  Java,
		"c":     C,
		"cpp":   Cpp,
		"rs":    Rust,
		"swift": Swift,
		"kt":    Kotlin,
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
