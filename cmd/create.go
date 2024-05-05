package cmd

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"go-gen/cmd/utils"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
)

// createCmd represents the custom command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create command allows you to generate a new Rest API backend using Golang.",
	Long: `Create command allows you to generate a new Rest API backend using Golang, 
			with options to chose preferred routing framework.

			Example: go-gen create -n <your_project_name> 
				or go-gen create --name <your_project_name>	

`,
	RunE: func(cmd *cobra.Command, args []string) error {
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
		dirPath := filepath.Join(usr.HomeDir, projectName)

		// creating project directory at the root user home directory
		err = os.Mkdir(dirPath, 0755)
		if err != nil {
			fmt.Printf("error creating new directory %v\n", err)
			return err
		}
		log.Println("Directory created at:", dirPath)

		// Prompt the user for options
		var options struct {
			RoutingFramework string
			Database         string
		}
		qs := []*survey.Question{
			{
				Name: "RoutingFramework",
				Prompt: &survey.Select{
					Message: "Select a routing framework:",
					Options: []string{"gorilla/mux", "Gin", "Echo"},
				},
			},
			{
				Name: "Database",
				Prompt: &survey.Select{
					Message: "Select a database:",
					Options: []string{"None"},
				},
			},
		}
		if err := survey.Ask(qs, &options); err != nil {
			return err
		}

		// Print selected options
		fmt.Println("Selected routing framework:", options.RoutingFramework)
		fmt.Println("Selected database:", options.Database)

		// Generate project folder
		err = utils.GenerateProjectFolder(dirPath)
		if err != nil {
			return err
		}
		// Generate boilerplate code based on selected options
		err = utils.GenerateCustomBoilerplate(options.RoutingFramework, options.Database, dirPath)
		if err != nil {
			return err
		}

		err = utils.InitGoModule(dirPath)
		if err != nil {
			return err
		}

		return nil
	},
}

func init() {
	createCmd.Flags().StringP("name", "n", "", "project name")
	rootCmd.AddCommand(createCmd)
}
