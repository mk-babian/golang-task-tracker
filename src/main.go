package main

import (
	"fmt"
	"os"
	"encoding/json"
)

type Task struct {
	ID			int		`json:"id"`
	Title		string 	`json:"title"`
}

func main(){
	fmt.Println("Welcome to task tracker!\n")

	for 1 < 2{
		var number int
		fmt.Print("Provide a number as a command (0 for help, -1 to exit): ")
		fmt.Scan(&number)

		switch number{
		case 0:
			printHelp()
		case 1:
			addTask()
		case 2:

		case -1:
			os.Exit(0)
		default:

		}
	}	
}

func printHelp(){
	
	// TODO:
	// Define each command (-1, 0, 1, etc...) []
	// Make it look good []

	fmt.Println("The help page !\n---\tCommands:\t---")
}

func addTask(){

	// TODO:
		// Create .json file [x]
		// Write task(s) to file [x]
		// Save file [x]

	outFile, err := os.Create("out.json")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	quest := Task {
		ID:		1,
		Title: 	"Quest 1",
	}

	jsonData, err := json.MarshalIndent(quest, "", "\t")
	if err != nil {
		panic(err)
	}

	outFile.Write(jsonData)
}

func delTask(){
	
	// TODO:
		// Read from existing .json file []
		// Delete task from .json file []
		// Save file []
}
