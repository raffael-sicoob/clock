/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

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
		fmt.Println("punch called")
	},
}

func init() {
	rootCmd.AddCommand(punchCmd)

	
}
