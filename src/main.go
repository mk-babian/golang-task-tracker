package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"encoding/json"
)

type Task struct {
	ID			int		`json:"id"`
	Title		string 	`json:"title"`
}

func main(){
	fmt.Println("Welcome to task tracker!\n")

	for 1 < 2{
		var	command string 
		fmt.Print("Provide a command (\"help\" for help and \"exit\" to quit) : ")
		fmt.Scanln(&command)
		
		switch command{
			case "help":
				printHelp()
			case "create":
				createJson()
			case "add":
				addTask()
			case "exit":
				os.Exit(0)
			default:
				os.Exit(0)
		}
	}	
}

func printHelp(){

	// TODO:
	// Define each command (-1, 0, 1, etc...) []
	// Make it look good []
	
	fmt.Println("Hello, world!")
}

func createJson(){
	outFile, err := os.Create("tasks.json")	
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	outFile.Close()
}

func addTask(){

	// TODO
		// Append to a .json file using the data given by the user (task name and stuff) []

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Provide a task name: ")
	taskNameNewLine, err := reader.ReadString('\n')

	file, err := os.OpenFile("tasks.json", os.O_RDWR, 0644)
	if err != nil {
		panic(err)
		return
	}
	defer file.Close()

	taskName := strings.TrimSpace(taskNameNewLine)
	
	fmt.Println(taskName)

	var idCount int = 1
	quest := Task {
		ID:		idCount,
		Title: 	taskName,
	}
	idCount++

	jsonData, err := json.MarshalIndent(quest, "", "\t")
	if err != nil {
		panic(err)
		return
	}
	
	file.Write(jsonData)

	file.Close()
}

func delTask(){
	file, err := os.OpenFile("tasks.json", os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	file.Close()
}
