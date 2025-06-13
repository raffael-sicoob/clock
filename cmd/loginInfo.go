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
	Short: "Get the current user logged info",
	Long: `Get the current user info.`,
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
