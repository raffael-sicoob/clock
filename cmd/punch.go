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

// punchCmd represents the punch command
var punchCmd = &cobra.Command{
	Use:   "punch",
	Aliases: []string{"p"},
	Example: `
	clock punch 
	clock p
	`,
	Short: "Punch enter or exit",
	Long: `
	Punch enter or exit
	Register a clocking with the current time
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		user := database.GetUser()
		token := database.GetToken()

		spinner := spinner.New(spinner.CharSets[14], 100*time.Millisecond, spinner.WithColor("blue"))
		spinner.Suffix = " Requesting clocking..."
		spinner.Start()

		currentDateTime := rhapi.RequestClocking(user, token)
		spinner.Stop()

		if currentDateTime.ActualDate == "" {
			fmt.Println("Error getting time")
			os.Exit(1)
		}

		colorGreenBold := color.New(color.FgGreen, color.Bold).SprintFunc()

		formattedCurrentDateTime, _ := utils.FormatTime(currentDateTime.ActualDate, currentDateTime.ActualTime)

		fmt.Println("Clocking registered: 🕑", colorGreenBold(formattedCurrentDateTime))
	},
}

func init() {
	rootCmd.AddCommand(punchCmd)

	
}
