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
		currentDateTime := rhapi.RequestClocking(user, token)
		

		if currentDateTime.ActualDate == "" {
			fmt.Println("Error getting time")
			os.Exit(1)
		}

		colorGreenBold := color.New(color.FgGreen, color.Bold).SprintFunc()

		formattedCurrentDateTime := utils.FormatTime(currentDateTime.ActualDate, currentDateTime.ActualTime)

		fmt.Println("Clocking requested: 🕑", colorGreenBold(formattedCurrentDateTime))
	},
}

func init() {
	rootCmd.AddCommand(punchCmd)

	
}
