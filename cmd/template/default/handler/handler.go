package template

import (
	"os"
	"path/filepath"
)

// function to generate handler boilerplate

func generateHandler(dirPath string) error {
	// create handler directory inside dirPath
	err := os.MkdirAll(filepath.Join(dirPath, "handler"), 0755)
	if err != nil {
		return err
	}

	// create  todo.go file inside handler folder that contains todos handlers using os package
	file, err := os.Create(filepath.Join(dirPath, "handler", "todo.go"))
	if err != nil {
		return err
	}
	defer file.Close()

	// write the handler code to the file
	_, err = file.WriteString(`package handler

	import (
		"fmt"
		"net/http"
	)
	
	func TodoHandler(w http.ResponseWriter, r *http.Request) {
		// handle todos
		fmt.Println(w, "Handling todos...")
	}`)

	if err != nil {
		return err
	}

	return nil
}
