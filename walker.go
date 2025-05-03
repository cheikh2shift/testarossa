package testarossa

import (
	"os"
	"path/filepath"
	"strings"
)

type SourceFile struct {
	Path    string
	Content string
}

func GetFileContents(root, suffix string) ([]SourceFile, error) {
	var sourceFiles []SourceFile

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // if there's an error accessing the path
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), suffix) {
			content, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			sourceFiles = append(sourceFiles, SourceFile{Path: path, Content: string(content)})
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return sourceFiles, nil
}
