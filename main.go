/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/raffael-sicoob/clock/cmd"
	"github.com/raffael-sicoob/clock/database"
)

func init() {
	database.InitDB()
}

func main() {

	defer database.CloseDB()
	cmd.Execute()
	
}
