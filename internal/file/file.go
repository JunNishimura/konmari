package file

import (
	"os"
	"path/filepath"
	"strings"
)

func ExtractFileName(path string) string {
	cleanedPath := strings.ReplaceAll(path, `\`, "/")
	pathSplit := strings.Split(cleanedPath, "/")
	return pathSplit[len(pathSplit)-1]
}

func ExtractFilesFromDirectory(dirPath string) ([]string, error) {
	files, err := os.ReadDir(dirPath)
	if err != nil {
		return nil, err
	}

	filePaths := make([]string, 0)
	for _, file := range files {
		path := strings.ReplaceAll(filepath.Join(dirPath, file.Name()), `\`, "/")
		if file.IsDir() {
			extractedFiles, err := ExtractFilesFromDirectory(path)
			if err != nil {
				return nil, err
			}
			filePaths = append(filePaths, extractedFiles...)
		} else {
			filePaths = append(filePaths, path)
		}
	}

	return filePaths, nil
}
