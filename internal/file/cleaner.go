package file

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"

	"github.com/JunNishimura/konmari/internal/extension"
)

const (
	defautlPostfix = "cleaned"
)

var (
	ErrNotAcceptibleExtension = errors.New("not acceptible file extension")
	regexpMap                 = map[extension.Mark]string{
		extension.Go:         `\/\*[\s\S]*?\*\/|\/\/.*`,
		extension.Python:     `#.*|#.*|'''[\s\S]*?'''`,
		extension.JavaScript: `\/\*[\s\S]*?\*\/|\/\/.*`,
		extension.TypeScript: `\/\*[\s\S]*?\*\/|\/\/.*`,
		extension.Ruby:       `#.*|=begin[\s\S]*?=end`,
		extension.PHP:        `\/\*[\s\S]*?\*\/|\/\/.*|#.*`,
		extension.Java:       `\/\*[\s\S]*?\*\/|\/\/.*`,
		extension.C:          `\/\*[\s\S]*?\*\/|\/\/.*`,
		extension.Cpp:        `\/\*[\s\S]*?\*\/|\/\/.*`,
		extension.Rust:       `\/\*[\s\S]*?\*\/|\/\/.*`,
		extension.Swift:      `\/\*[\s\S]*?\*\/|\/\/.*`,
		extension.Kotlin:     `\/\*[\s\S]*?\*\/|\/\/.*`,
		extension.Dart:       `\/\*[\s\S]*?\*\/|\/\/.*`,
		extension.HTML:       `<!--[\s\S]*?-->`,
	}
)

type Cleaner struct {
	extension extension.Mark
	filePath  string
}

func NewCleaner(filePath string) *Cleaner {
	cleaner := newCleaner(filePath)

	fileName := extractFileName(filePath)
	ext := extension.New(fileName)
	cleaner.extension = ext

	return cleaner
}

func newCleaner(filePath string) *Cleaner {
	return &Cleaner{
		filePath: filePath,
	}
}

func (c *Cleaner) Execute(isOverWrite bool, postfix string) error {
	if c.extension == extension.Undefined {
		return ErrNotAcceptibleExtension
	}

	src, err := os.ReadFile(c.filePath)
	if err != nil {
		return err
	}

	// regexp replacement
	pattern, ok := regexpMap[c.extension]
	if !ok {
		return ErrNotAcceptibleExtension
	}
	re := regexp.MustCompile(pattern)
	cleanedSrc := re.ReplaceAll(src, []byte(""))

	// write to file
	if isOverWrite {
		f, err := os.Create(c.filePath)
		if err != nil {
			return err
		}
		defer f.Close()
		if _, err := f.Write(cleanedSrc); err != nil {
			return err
		}
	} else {
		parentDir := filepath.Dir(c.filePath)
		fileName := extractFileName(c.filePath)
		var cleanedFileName string
		if postfix == "" {
			cleanedFileName = addPostfixToFileName(fileName, defautlPostfix)
		} else {
			cleanedFileName = addPostfixToFileName(fileName, postfix)
		}
		newFilePath := filepath.Join(parentDir, cleanedFileName)
		f, err := os.Create(newFilePath)
		if err != nil {
			return err
		}
		defer f.Close()
		if _, err := f.Write(cleanedSrc); err != nil {
			return err
		}
	}

	return nil
}
