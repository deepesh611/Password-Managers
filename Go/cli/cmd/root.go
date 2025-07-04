package cmd

import (
	"cli/models"
	"cli/utils"
	"fmt"
)

var entries []models.Entry

func Run() {
	for {
		fmt.Println("\n1) Add entry")
		fmt.Println("2) List entries")
		fmt.Println("3) Exit")

		choice := utils.ReadInput("Choose option: ")

		switch choice {
		case "1":
			addEntry()
		case "2":
			listEntries()
		case "3":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func addEntry() {
	service := utils.ReadInput("Service: ")
	username := utils.ReadInput("Username: ")
	password := utils.ReadInput("Password: ")

	entry := models.Entry{
		Service:  service,
		Username: username,
		Password: password,
	}

	entries = append(entries, entry)
	fmt.Println("Entry added.")
}

func listEntries() {
	if len(entries) == 0 {
		fmt.Println("No entries found.")
		return
	}

	for i, entry := range entries {
		fmt.Printf("%d) %s | %s | %s\n", i+1, entry.Service, entry.Username, entry.Password)
	}
}
