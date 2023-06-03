package file

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"

	"github.com/JunNishimura/konmari/internal/extension"
)

var (
	ErrNotAcceptibleExtension = errors.New("not acceptible file extension")
	regexpMap                 = map[extension.Mark]string{
		extension.Go: `\/\*[\s\S]*?\*\/|\/\/.*`,
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

func (c *Cleaner) Execute() error {
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
	parentDir := filepath.Dir(c.filePath)
	fileName := extractFileName(c.filePath)
	cleanedFileName := addPostfixToFileName(fileName, "cleaned")
	newFilePath := filepath.Join(parentDir, cleanedFileName)
	f, err := os.Create(newFilePath)
	if err != nil {
		return err
	}
	defer f.Close()
	if _, err := f.Write(cleanedSrc); err != nil {
		return err
	}

	return nil
}