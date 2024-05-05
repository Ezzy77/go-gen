package cmd

//
//import (
//	"fmt"
//	"github.com/spf13/cobra"
//	template2 "go-gen/cmd/template/default/handler"
//	template "go-gen/cmd/template/default/httpServer"
//	"go-gen/cmd/template/default/route"
//	"go-gen/cmd/utils"
//	"log"
//	"os"
//	"os/user"
//	"path/filepath"
//)
//
//// newCmd represents the new command
//var newCmd = &cobra.Command{
//	Use:   "new",
//	Short: "new create a new golang project",
//	Long: `new create a new golang project:
//.`,
//	RunE: func(cmd *cobra.Command, args []string) error {
//		projectName, err := cmd.Flags().GetString("name")
//		if err != nil {
//			fmt.Println("error parsing project name")
//			return err
//		}
//
//		usr, err := user.Current()
//		if err != nil {
//			fmt.Println("error getting system user")
//			return err
//		}
//		// Define path for new directory
//		dirPath := filepath.Join(usr.HomeDir, projectName)
//
//		// creating directory at the root user home directory
//		err = os.Mkdir(dirPath, 0755)
//		if err != nil {
//			fmt.Printf("error creating new directory %v\n", err)
//			return err
//		}
//		log.Println("Directory created at:", dirPath)
//
//		// generate main.go file and http server boilerplate
//		err = template.GenerateHttpServer(dirPath)
//		if err != nil {
//			return err
//		}
//
//		err = template2.GenerateHandler(dirPath)
//		if err != nil {
//			return err
//		}
//
//		err = route.GenerateRoutes(dirPath)
//		if err != nil {
//			return err
//		}
//
//		err = utils.InitGoModule(dirPath)
//		if err != nil {
//			return err
//		}
//
//		return nil
//	},
//}
//
//func init() {
//	newCmd.Flags().StringP("name", "n", "", "project name")
//	rootCmd.AddCommand(newCmd)
//
//}
