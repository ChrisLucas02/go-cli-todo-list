package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type ToDo struct {
	name    string
	dueDate time.Time
}

var todoList []ToDo

func main() {
	clearTerminal()
	printHeader()
	printList()
	var isRunning = true

	for isRunning {
		runApp(&isRunning)
	}
	clearTerminal()
}

func runApp(isRunning *bool) {

	var option = 0
	fmt.Print("Type 0 for help. How can i help you ? ")
	fmt.Scan(&option)
	switch option {
	case 0:
		printHelp()
		return
	case 1:
		fmt.Print("Add a new item <name,date> : ")
		var line = ""
		fmt.Scan(&line)
		var split = strings.Split(line, ",")
		var text = split[0]
		var date = split[1]
		addItem(text, date)
	case 2:
		clearTerminal()
		printHeader()
		printList()
		var number = 0
		fmt.Print("Select the task to be updated <number> : ")
		fmt.Scan(&number)

		var nameChange = ""
		fmt.Print("Do you want the update the name ? <N,y>")
		fmt.Scan(&nameChange)

		var name = ""
		if strings.ToLower(nameChange) == "y" {
			fmt.Print("Set new name <name> : ")
			fmt.Scan(&name)
		}

		var dateChange = ""
		fmt.Print("Do you want the update the date ? <N,y>")
		fmt.Scan(&dateChange)

		var date = ""
		if strings.ToLower(dateChange) == "y" {
			fmt.Print("Set new date <date> : ")
			fmt.Scan(&date)
		}

		updateItem(number-1, name, date)
	case 3:
		var number = 0
		fmt.Print("Remove a task <number> : ")
		fmt.Scan(&number)
		deleteItem(number)
	case 4:
		printList()
	case 5:
		*isRunning = false
	}
	printHeader()
	printList()
}

func addItem(text string, date string) {
	var formatedText = capitalizeFirstLetter(text)
	formatedDate, _ := time.Parse("2006-01-02", date)
	var newItem = ToDo{formatedText, formatedDate}
	todoList = append(todoList, newItem)
}

func updateItem(number int, name string, date string) {
	var toBeUpdated = &todoList[number]
	if len(name) > 0 {
		toBeUpdated.name = capitalizeFirstLetter(name)
	}
	if len(date) > 0 {
		formatedDate, _ := time.Parse("2006-01-02", date)
		toBeUpdated.dueDate = formatedDate
	}
}

func deleteItem(number int) {
	todoList = append(todoList[:number-1], todoList[number:]...)
}

func printList() {
	if len(todoList) == 0 {
		fmt.Println("No task need to be done")
		fmt.Println("")
		fmt.Println("")
		return
	}
	for i, item := range todoList {
		fmt.Printf("%d)	%s\n", i+1, item.name)
		fmt.Printf("	Due Date: %s\n", item.dueDate.Format("02, Jan 2006"))
		fmt.Println("")
	}
	fmt.Println("")
}

func printHeader() {
	clearTerminal()
	fmt.Println("\nCLI To-Do List !")
	fmt.Println("--------------------------------")
	fmt.Println("")
}

func printHelp() {
	printHeader()
	fmt.Println("1) Add new item to the list")
	fmt.Println("2) Update an item on the list")
	fmt.Println("3) Delete an item on the list")
	fmt.Println("4) Display the To-Do list")
	fmt.Println("")
}

func clearTerminal() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	lower := strings.ToLower(s)
	return strings.ToUpper(string(lower[0])) + lower[1:]
}
