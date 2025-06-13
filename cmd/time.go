/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/raffael-sicoob/clock/database"
	rhapi "github.com/raffael-sicoob/clock/rhApi"
	"github.com/raffael-sicoob/clock/utils"
	"github.com/spf13/cobra"
)

// timeCmd represents the time command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Get the current time",
	Long: `Get the current time.
	
	`,
	PreRun: func(cmd *cobra.Command, args []string)  {
		user := database.GetUser()
		if user.Username == "" {
			fmt.Println("user not found, please run 'clock login'")
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		user := database.GetUser()
    spinner := spinner.New(spinner.CharSets[14], 100*time.Millisecond, spinner.WithColor("blue"))
		spinner.Suffix = " Getting time..."
		spinner.Start()
		
		token := database.GetToken()
		dataTime := rhapi.GetTime(user, token)

		spinner.Stop()

		if dataTime.ActualDate == "" {
			fmt.Println("error getting time")
			os.Exit(1)
		}

		colorGreenBold := color.New(color.FgGreen, color.Bold).SprintFunc()

		formattedTime,_ := utils.FormatTime(dataTime.ActualDate, dataTime.ActualTime)

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
