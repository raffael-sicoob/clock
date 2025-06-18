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
	Aliases: []string{"ls"},
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
		 for i, clocking := range listClockings.Clockings {
			_, dateClocking := utils.FormatTime(clocking.Date, clocking.Hour)
			if i > 0 && listClockings.Clockings[i-1].Date != clocking.Date {
				fmt.Println(color.HiBlackString("--------"))
			}

			fmt.Println(
				dateClocking.Format("02/01/2006"),
				color.HiBlackString("--------"), 
				"🕑",
				colorCyanBold(dateClocking.Format("15:04:05")),
					color.HiBlackString("--------"), 
				colorDirection(clocking.Direction))
		 }

	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	currentDate := time.Now().Format("2006-01-02")
	listCmd.Flags().StringP("start", "s", currentDate, "Start date (YYYY-MM-DD)")
	listCmd.Flags().StringP("end", "e", currentDate, "End date (YYYY-MM-DD)")


}


func colorDirection(direction string) string {
	if direction == "entry" {
		return color.GreenString("🟢 Entry")
	}
	return color.RedString("🔴 Exit")
}


