package main

import (
	"bufio"
	"fmt"
	"os"
	"io"
	"strings"
	"encoding/json"
)

type Task struct {
	ID			int		`json:"id"`
	Title		string 	`json:"title"`
	Priority	string 	`json:"priority"`
}

func main(){
	fmt.Println("Welcome to task tracker!\n")

	for true {
		var	command string 
		fmt.Print("Provide a command (\"help\" for help and \"exit\" to quit) : ")
		fmt.Scanln(&command)

		switch command{
		case "help":
			printHelp()
		case "create":
			createJson()
		case "list":
			listTasks()
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
	// Define each command [x]
	// Make it look good [x]

	fmt.Println("\t---COMMANDS---\n")	
	fmt.Println("help - Display help")
	fmt.Println("create - Create .json file in running directory")
	fmt.Println("list - List the tasks currently in the .json file")
	fmt.Println("add - Add a task to tasks.json")
	fmt.Println("exit - Quit the application")
	fmt.Println()
}

func createJson(){
	// Check if file exists
	_, err := os.Stat("tasks.json")
	if os.IsNotExist(err) {
		outFile, err := os.Create("tasks.json")	
		if err != nil {
			fmt.Println("ERR | Failed to create JSON file")
			return
		}
		defer outFile.Close()
	}else{
		fmt.Println("Creating this file will overwrite the current tasks file.")
		fmt.Print("Are you sure [y/n]: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		answer := strings.ToLower(strings.TrimSpace(input))

		if len(answer) > 0 && answer[0] == 'y' {
			outFile, err := os.Create("tasks.json")
			if err != nil {
				fmt.Println("ERR | Failed to create JSON file:", err)
				return
			}
			defer outFile.Close()
			fmt.Println("tasks.json overwritten successfully.")
		} else {
			fmt.Println("File not created.")
			return
		}
	}

}

func listTasks(){
	file, err := os.OpenFile("tasks.json", os.O_RDONLY, 0644)
	if err != nil {
		fmt.Println("ERR | Failed to open tasks.json", err)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("ERR | Failed to read file", err)
		return
	}

	var tasks []Task
	err = json.Unmarshal(fileBytes, &tasks)
	if err != nil {
		fmt.Println("ERR | Invalid JSON: ", err)
		return
	}

	for _, t := range tasks {
		fmt.Println(t)
	}
}

func addTask(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Provide a task name: ")
	taskName, _ := reader.ReadString('\n')
	taskName = strings.TrimSpace(taskName)
	fmt.Print("Provide task priority (low/med/high): ")
	priority, _ := reader.ReadString('\n')
	priority = strings.TrimSpace(priority)
	if priority != "low" && priority != "med" && priority != "high" {
		fmt.Println("Provide a valid priority (low/med/high)!")
		return
	}

	file, err := os.OpenFile("tasks.json", os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("ERR | Failed open tasks.json", err)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("ERR | Failed to read file", err)
		return
	}

	var tasks []Task
	if len(fileBytes) > 0 {
		if err := json.Unmarshal(fileBytes, &tasks); err != nil {
			fmt.Println("ERR | File contains invalid JSON", err)
			return
		}
	}
	newID := len(tasks) + 1

	newTask := Task {
		ID:			newID,
		Title: 		taskName,
		Priority: 	priority,
	}

	tasks = append(tasks, newTask)

	output, err := json.MarshalIndent(tasks, "", "\t")
	if err != nil {
		fmt.Println("ERR | Failed to format/marshal file", err)
		return
	}
	if err := file.Truncate(0); err != nil {
		fmt.Println("ERR | Failed to clear file", err)
		return
	}
	if _, err := file.Seek(0, 0); err != nil {
		fmt.Println("ERR | Failed to reset file pointer", err)
		return
	}
	if _, err := file.Write(output); err != nil {
		fmt.Println("ERR | Coudln't write tasks", err)
		return
	}
}

func delTask(){
	file, err := os.OpenFile("tasks.json", os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
}
