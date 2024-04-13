package template

import (
	"os"
	"path/filepath"
)

// functions to generate http server boilerplate

func generateHttpServer(dirPath string) error {
	// Create main.go file
	mainFilePath := filepath.Join(dirPath, "main.go")
	mainFile, err := os.Create(mainFilePath)
	if err != nil {
		return err
	}
	defer mainFile.Close()
	return nil
}
