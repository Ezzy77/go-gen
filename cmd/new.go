/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
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

		return nil
	},
}

func init() {
	newCmd.Flags().StringP("name", "n", "", "project name")
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
