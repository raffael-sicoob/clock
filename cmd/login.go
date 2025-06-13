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
	Short: "Login to the system",
	Long: `Login to the system.`,


	Run: func(cmd *cobra.Command, args []string) {

			user := database.GetUser()
		

			if user.Username != "" {
				fmt.Println("User already logged in with name", user.Name)
				return
			}

			username, _ := cmd.Flags().GetString("username")
			password, _ := cmd.Flags().GetString("password")
			name, _ := cmd.Flags().GetString("name")
			token := rhapi.Login(username, password)

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
