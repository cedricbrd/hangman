package hangman

import (
	"os"
	"strings"
)

func listFilesInDirectory(directory, extension string) ([]string, error) {
	var files []string
	dir, err := os.Open(directory)
	if err != nil {
		return nil, err
	}
	defer dir.Close()
	fileInfos, err := dir.Readdir(0)
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range fileInfos {
		if !fileInfo.IsDir() && strings.HasSuffix(fileInfo.Name(), extension) {
			files = append(files, fileInfo.Name())
		}
	}
	return files, nil
}
