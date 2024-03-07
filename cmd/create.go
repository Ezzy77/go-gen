/*
Copyright © 2024 Elisio Cabral
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create command create a new Go project",
	Long: `The create command initialises a new Golang backend project, when
			run without argument, default structure will be generated.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("create called")

		// Get current user
		usr, err := user.Current()
		if err != nil {
			log.Fatalf("Cannot get current user: %v", err)
		}
		// Define path for new directory
		dir := filepath.Join(usr.HomeDir, "testDir")

		// Creating directory
		err = os.Mkdir(dir, 0755)
		if err != nil {
			log.Fatalf("Error creating directory: %v", err)
		}
		log.Println("Directory created at:", dir)
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
