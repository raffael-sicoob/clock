/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/raffael-sicoob/clock/database"
	rhapi "github.com/raffael-sicoob/clock/rhApi"
	"github.com/spf13/cobra"
)

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit",
	Short: "Edit any clocking",

	Run: func(cmd *cobra.Command, args []string) {
		date, _ := cmd.Flags().GetString("date")
		hour, _ := cmd.Flags().GetString("time")

		user := database.GetUser()
		token := database.GetToken()

		spinner := spinner.New(spinner.CharSets[14], 100*time.Millisecond, spinner.WithColor("blue"))
		spinner.Suffix = " Requesting clocking edit..."
		spinner.Start()

		err := rhapi.EditClockings(user, token, date, hour)
		spinner.Stop()

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("✅ ", color.GreenString("Clocking edited successfully!"))

	},
}

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.Flags().StringP("date", "d", "", "Date (YYYY-MM-DD)")
	editCmd.MarkFlagRequired("date")
	editCmd.Flags().StringP("time", "t", "", "Time (1h30m for 1 hour and 30 minutes)")
	editCmd.MarkFlagRequired("time")

}
