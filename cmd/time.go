/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/raffael-sicoob/clock/database"
	rhapi "github.com/raffael-sicoob/clock/rhApi"
	"github.com/raffael-sicoob/clock/utils"
	"github.com/spf13/cobra"
)

// timeCmd represents the time command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	PreRun: func(cmd *cobra.Command, args []string)  {
		user := database.GetUser()
		if user.Username == "" {
			fmt.Println("user not found, please run 'clock login'")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		user := database.GetUser()
		token := database.GetToken()
		dataTime := rhapi.GetTime(user, token)

		if dataTime.ActualDate == "" {
			fmt.Println("error getting time")
			os.Exit(1)
		}

		colorGreenBold := color.New(color.FgGreen, color.Bold).SprintFunc()

		formattedTime := utils.FormatTime(dataTime.ActualDate, dataTime.ActualTime)

		fmt.Println("Actual time: 🕑", colorGreenBold(formattedTime))
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
