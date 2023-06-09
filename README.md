# konmari / CLI tool to clean up code comment

<p align='left'>
  <img alt="GitHub release (latest by date)" src="https://img.shields.io/github/v/release/JunNishimura/konmari">
  <img alt="GitHub" src="https://img.shields.io/github/license/JunNishimura/konmari">
  <a href="https://github.com/JunNishimura/konmari/actions/workflows/test.yml"><img src="https://github.com/JunNishimura/konmari/actions/workflows/test.yml/badge.svg" alt="test"></a>
  <a href="https://goreportcard.com/report/github.com/JunNishimura/konmari"><img src="https://goreportcard.com/badge/github.com/JunNishimura/konmari" alt="Go Report Card"></a>
</p>

## 📖 Overview
konmari is a CLI tool that allows you to delete comments in source code in bulk.

## 💻 Installation
### Homebrew Tap
```
brew install JunNishimura/tap/konmari
```

### go intall
```
go install github.com/JunNishimura/konmari@latest
```

## 🔨 Options
```console
$ konmari -h 
clean up source code comments

Usage:
  konmari [flags]

Flags:
  -h, --help             help for konmari
  -o, --overwrite        overwrite existing files
  -p, --postfix string   postfix for cleaned files
```
### `-o`, `--overwrite`
By default, konmari outputs the result of comment deletion to a separate file. However, by specifying this option, existing files can be overwritten.

### `-p`, `--postfix`
Specify the postfix of the file that is the output destination of the comment deletion result.

## 👀 Example
To apply konmari to the following main.go file,
```go:main.go
package main

// import package
import "fmt"

/*
	this is a test go file
*/

// main function
func main() {
	fmt.Println("Hello, Wolrd") // print out hello world
}

```

The output will look like this
```go:main.go
package main


import "fmt"




func main() {
	fmt.Println("Hello, Wolrd") 
}

```

## 😢 Limitation
konmari has flows. konmari does not work as expected in the following case.
```go:main.go
package main

// import package
import "fmt"

// main function
func main() {
	fmt.Println("this // will be treated as comment") // print out hello world
}
```

This is the output.
```go:main.go
package main


import "fmt"


func main() {
	fmt.Println("this 
}

```

konmari incorrectly recognizes comment symbols in strings as comments as well. This will be fixed in the future, so please be patient.

## 🪧 License
konmari is released under MIT License. See [MIT](https://raw.githubusercontent.com/JunNishimura/konmari/main/LICENSE)
