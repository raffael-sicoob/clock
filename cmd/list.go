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
	"github.com/raffael-sicoob/clock/utils"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Aliases: []string{"l"},
	Short: "List all clockings for an range of dates",
	Long: `List all clockings for an range of dates.
	Default range is today.
	Example:
	clock list --start 2025-01-01 --end 2025-01-31`,
	
	Run: func(cmd *cobra.Command, args []string) {
		startDate, _ := cmd.Flags().GetString("start")
		endDate, _ := cmd.Flags().GetString("end")
		
		user := database.GetUser()
		token := database.GetToken()

		if( endDate == ""	 &&startDate  != "") {
			endDate = startDate
		}

		if( startDate == ""	 &&endDate  != "") {
			startDate = endDate
		}

		spinner := spinner.New(spinner.CharSets[14], 100*time.Millisecond, spinner.WithColor("blue"))
		spinner.Suffix = " Requesting clockings..."
		spinner.Start()

		 listClockings := rhapi.GetClockings(user, token, startDate, endDate)
		 spinner.Stop()
		 
		 if len(listClockings.Clockings) == 0 {
			fmt.Println("No clockings found")
			return
		 }

		 colorCyanBold := color.New(color.FgCyan, color.Bold).SprintFunc()
		 for _, clocking := range listClockings.Clockings {
			_, dateClocking := utils.FormatTime(clocking.Date, clocking.Hour)

			fmt.Println(
				dateClocking.Format("02/01/2006"),
				"--------", 
				"🕑",
				colorCyanBold(dateClocking.Format("15:04:05")),
					"--------", 
				colorDirection(clocking.Direction))
		 }

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	currentDate := time.Now().Format("2006-01-02")
	listCmd.Flags().StringP("start", "s", currentDate, "Start date (YYYY-MM-DD)")
	listCmd.Flags().StringP("end", "e", currentDate, "End date (YYYY-MM-DD)")


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


func colorDirection(direction string) string {
	if direction == "entry" {
		return color.GreenString("🟢 Entry")
	}
	return color.RedString("🔴 Exit")
}


