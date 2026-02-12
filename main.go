package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type Operation int

const (
	ADD = iota
	EDIT
	DEL
	TOGGLE
	LIST
	CLEAR
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		printHelp()
		os.Exit(1)
	}

	var tasks []Task

	savePath, err := getSavePath()
	if err != nil {
		printErr(err)
		os.Exit(1)
	}

	LoadTasks(savePath, &tasks)
	handleOperation(args, &tasks)
}

func getSavePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return homeDir, err
	}

	return homeDir + "/.task.json", nil
}

func checkArgCount(args []string, op Operation) error {
	switch op {
	case ADD, DEL, TOGGLE:
		if len(args) < 3 {
			return errors.New("too few arguments")
		}
	case EDIT:
		if len(args) < 4 {
			return errors.New("too few arguments")
		}
	default:
		return errors.New("invalid operation")
	}

	return nil
}

func printErr(err error) {
	fmt.Printf("Error: %s\n", err)
}

func printHelp() {
	fmt.Printf(`
task <command> [arguments]

help			show this message
add [text]		add a task
edit <id> <text>	edit a task
del [id]		delete a task
done [id]		toggle task completion
list 			list all tasks

`)
}

func handleOperation(args []string, tasks *[]Task) {
	operation := args[1]
	switch operation {
	case "add":
		handleAdd(args, tasks)
	case "edit":
		handleEdit(args, tasks)
	case "del":
		handleDelete(args, tasks)
	case "done":
		handleToggle(args, tasks)
	case "clear":
		handleClear(tasks)
	case "list":
		err := ListTasks(*tasks)
		if err != nil {
			printErr(err)
		}
	default:
		printHelp()
	}
}

func handleAdd(args []string, tasks *[]Task) {
	err := checkArgCount(args, ADD)
	if err != nil {
		printErr(err)
		os.Exit(1)
	}
	input := args[2]
	*tasks = append(*tasks, Task{input, false})
	saveToFile(tasks)
}

func handleEdit(args []string, tasks *[]Task) {
	err := checkArgCount(args, EDIT)
	if err != nil {
		printErr(err)
		os.Exit(1)
	}

	index, err := strconv.Atoi(args[2])
	if err != nil {
		printErr(err)
		os.Exit(1)
	}
	text := args[3]
	*tasks, err = EditTask(*tasks, index-1, text)
	if err != nil {
		printErr(err)
		os.Exit(1)
	}
	saveToFile(tasks)
}

func handleDelete(args []string, tasks *[]Task) {
	err := checkArgCount(args, DEL)
	if err != nil {
		printErr(err)
		os.Exit(1)
	}
	index, err := strconv.Atoi(args[2])
	if err != nil {
		index = -1
	}
	newList, err := DeleteTask(*tasks, index-1)
	if err != nil {
		printErr(err)
	} else {
		*tasks = newList
		saveToFile(tasks)
	}
}

func handleToggle(args []string, tasks *[]Task) {
	err := checkArgCount(args, TOGGLE)
	if err != nil {
		printErr(err)
		os.Exit(1)
	}
	index, err := strconv.Atoi(args[2])
	if err != nil {
		index = -1
	}
	newTasks, err := ToggleTask(*tasks, index-1)
	if err != nil {
		printErr(err)
	} else {
		*tasks = newTasks
		saveToFile(tasks)
	}
}

func handleClear(tasks *[]Task) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Clear task file? (y/N): ")
	scanner.Scan()

	switch scanner.Text() {
	case "Y", "y":
		fmt.Println("Task file cleared")
		*tasks = []Task{}
		saveToFile(tasks)
	default:
		os.Exit(1)
	}
}

func saveToFile(tasks *[]Task) {
	savePath, err := getSavePath()
	if err != nil {
		printErr(err)
		os.Exit(1)
	}
	SaveTasks(savePath, *tasks)
}
