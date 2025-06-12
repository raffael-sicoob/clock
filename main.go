/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/raffael-sicoob/clock/cmd"
	"github.com/raffael-sicoob/clock/database"
)

func main() {

	database.InitDB()
	defer database.CloseDB()
	cmd.Execute()
	
}
