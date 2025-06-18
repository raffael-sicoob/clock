/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/fatih/color"
	"github.com/raffael-sicoob/clock/database"
	rhapi "github.com/raffael-sicoob/clock/rhApi"
	"github.com/raffael-sicoob/clock/utils"
	"github.com/spf13/cobra"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Aliases: []string{"s"},
	Short: "Get the balance summary between two dates",
	Run: func(cmd *cobra.Command, args []string) {
		startDate, _ := cmd.Flags().GetString("start")
		endDate, _ := cmd.Flags().GetString("end")

		user := database.GetUser()
		token := database.GetToken()

		if user.Username == "" {
			fmt.Println("User not found, please run 'clock login'")
			os.Exit(1)
		}

		spinner := spinner.New(spinner.CharSets[14], 100*time.Millisecond, spinner.WithColor("blue"))
		spinner.Suffix = " Requesting balance summary..."
		spinner.Start()

		balanceSummary := rhapi.GetBalanceSummary(user, token, startDate, endDate)
		spinner.Stop()

		if balanceSummary == nil {
			fmt.Println("Error getting balance summary")
			os.Exit(1)
		}




		fmt.Println(strings.Repeat(color.HiBlackString("-"), 40))
		fmt.Println("Balance Summary Period: ", color.CyanString(startDate), " to ", color.CyanString(endDate))
		fmt.Println(strings.Repeat(color.HiBlackString("-"), 40))
		fmt.Println("Previous Balance: ", strings.Repeat(" ", 12), printColorSummary(balanceSummary.Previous))
		fmt.Println("Current Balance: ", strings.Repeat(" ", 13 ), printColorSummary(balanceSummary.Current))
		fmt.Println("Next Balance: ", strings.Repeat(" ", 16), printColorSummary(balanceSummary.Next))
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)

	currentDate := time.Now().Format("2006-01-02")
	startDayOfMonth := time.Now().AddDate(0, 0, -time.Now().Day() + 1).Format("2006-01-02")
	summaryCmd.Flags().StringP("start", "s", startDayOfMonth, "Start date (YYYY-MM-DD)")
	summaryCmd.Flags().StringP("end", "e", currentDate, "End date (YYYY-MM-DD)")
	
}


func printColorSummary(duration int32) string{
	colorGreenBold := color.New(color.FgGreen, color.Bold).SprintFunc()
	colorRedBold := color.New(color.FgRed, color.Bold).SprintFunc()

	if duration > 0 {
		return colorGreenBold(utils.FormatDuration(duration))
	} else if duration == 0 {
		return utils.FormatDuration(duration)
	}

	return colorRedBold(utils.FormatDuration(duration))
	
}