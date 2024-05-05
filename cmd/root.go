/*
Copyright © 2024 Elisio Sa
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-gen",
	Short: "go-gen is a CLI tool for generating Restful API in Golang",
	Long: `go-gen is a CLI library for Go that empowers applications.
			This application is a tool to generate the needed files
			to quickly create a Cobra application.

			Example: go-gen create -n <your_project_name> 
				or go-gen create --name <your_project_name>		
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main()
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
