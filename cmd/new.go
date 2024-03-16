/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "new create a new golang project",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("new called")
		projectName, err := cmd.Flags().GetString("name")
		if err != nil {
			fmt.Println("error parsing project name")
			return err
		}

		usr, err := user.Current()
		if err != nil {
			fmt.Println("error getting system user")
			return err
		}
		// Define path for new directory
		dir := filepath.Join(usr.HomeDir, projectName)

		// creating directory at the root user home directory
		err = os.Mkdir(dir, 0755)
		if err != nil {
			fmt.Printf("error creating new directory %v\n", err)
			return err
		}
		log.Println("Directory created at:", dir)

		// creates main.go file
		file, err := os.Create(dir + "/main.go")
		if err != nil {
			return err
		}
		defer file.Close()

		mainFileContent :=
			`package main

			import (
				"fmt"
			)
			
			func main() {
				fmt.Println("Hello world!")
			}`

		_, err = fmt.Fprintf(file, mainFileContent)
		if err != nil {
			return err
		}

		err = initGoModule(dir)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	newCmd.Flags().StringP("name", "n", "", "project name")
	rootCmd.AddCommand(newCmd)

}

// initGoModule initialises a Go module for the project located in the given project directory.
// It creates a command to run the "go mod init" command and sets the project directory as the working directory.
// If an error occurs while executing the command, it prints an error message and returns the error.
// On successful initialisation, it prints a success message and returns nil.
func initGoModule(projectDir string) error {
	cmd := exec.Command("go", "mod", "init",
		"github.com/username/your_project_name")

	cmd.Dir = projectDir
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
		return err
	}
	//err := cmd.Run()
	//if err != nil {
	//	fmt.Println("error initializing go module")
	//	fmt.Println("here the error", err)
	//	return err
	//}
	fmt.Println("Go module initialized successfully")
	return nil
}
