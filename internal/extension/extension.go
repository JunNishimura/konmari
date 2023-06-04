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
	Csharp     Mark = "c#"
	Rust       Mark = "rust"
	Swift      Mark = "swift"
	Kotlin     Mark = "kotlin"
	Dart       Mark = "dart"
	HTML       Mark = "html"
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
		"cs":    Csharp,
		"rs":    Rust,
		"swift": Swift,
		"kt":    Kotlin,
		"dart":  Dart,
		"html":  HTML,
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
