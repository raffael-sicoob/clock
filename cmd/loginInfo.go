/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/raffael-sicoob/clock/database"
	"github.com/spf13/cobra"
)

// loginInfoCmd represents the loginInfo command
var loginInfoCmd = &cobra.Command{
	Use:   "info",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		user := database.GetUser()
		if user.Username == "" {
			fmt.Printf("user not found, please run 'clock login'\n")
			os.Exit(1)
		}
		fmt.Printf("Name: %s; Username: %s\n", user.Name, user.Username)
	},
}

func init() {
	loginCmd.AddCommand(loginInfoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginInfoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginInfoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
