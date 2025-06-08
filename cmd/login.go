/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/raffael-sicoob/clock/database"
	rhapi "github.com/raffael-sicoob/clock/rhApi"
	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,


	Run: func(cmd *cobra.Command, args []string) {

			user, err := database.GetUser()
			if err != nil {
				fmt.Println("Error getting user", err)
				return
			}

			if user.Username != "" {
				fmt.Println("User already logged in with name", user.Name)
				return
			}

			username, _ := cmd.Flags().GetString("username")
			password, _ := cmd.Flags().GetString("password")
			name, _ := cmd.Flags().GetString("name")
			token, err := rhapi.Login(username, password)
			if err != nil {
				fmt.Println("Error logging in", err)
				return
			}

			if name == "" {
				name = username
			}

			newUser := database.User{
				Name: name,
				Username: username,
				Password: password,
			}


			database.SaveToken(token)
			database.CreateUser(newUser)


			fmt.Println("Login successful with name", name)
	
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringP("username", "u", "", "Username")
	loginCmd.MarkFlagRequired("username")
	loginCmd.Flags().StringP("password", "p", "", "Password")
	loginCmd.MarkFlagRequired("password")
	loginCmd.Flags().StringP("name", "n", "", "Name")

}
