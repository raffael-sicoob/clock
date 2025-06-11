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

	// actualDate := "2025-06-11T15:00:00Z"
	// actualTime := 72666254 // Duração em microssegundos

	// // Parse a data
	// parsedDate, err := time.Parse(time.RFC3339, actualDate)
	// if err != nil {
	// 	fmt.Println("Erro ao parsear a data:", err)
	// 	return
	// }


	// duration := time.Duration(int64(actualTime)) * time.Millisecond

	// formattedTime :=  duration.String()

	// // Imprimir a data e a hora
	// fmt.Printf("Data: %s\n", parsedDate.Format("2006-01-02"))
	// fmt.Printf("Hora: %s\n", formattedTime)
}
