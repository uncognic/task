package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args

	if len(args) <= 1 {
		printHelp()
		os.Exit(1)
	}

	var tasks []Task

	LoadTasks("tasks.json", &tasks)

	operation := args[1]
	switch operation {
	case "add":
		err := checkArgCount(args)
		if err != nil {
			printErr(err)
			os.Exit(1)
		}
		input := args[2]
		tasks = append(tasks, Task{input, false})
		SaveTasks("tasks.json", tasks)

	case "del":
		err := checkArgCount(args)
		if err != nil {
			printErr(err)
			os.Exit(1)
		}
		index, err := strconv.Atoi(args[2])
		if err != nil {
			index = -1
		}
		newList, err := DeleteTask(tasks, index-1)
		if err != nil {
			printErr(err)
		} else {
			tasks = newList
			SaveTasks("tasks.json", tasks)
		}

	case "done":
		err := checkArgCount(args)
		if err != nil {
			printErr(err)
			os.Exit(1)
		}
		index, err := strconv.Atoi(args[2])
		if err != nil {
			index = -1
		}
		newTasks, err := ToggleTask(tasks, index-1)
		if err != nil {
			printErr(err)
		} else {
			tasks = newTasks
			SaveTasks("tasks.json", tasks)
		}

	case "list":
		err := ListTasks(tasks)
		if err != nil {
			printErr(err)
		}

	case "clear":
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Clear task file? (y/N): ")
		scanner.Scan()

		switch scanner.Text() {
		case "Y", "y":
			fmt.Println("Task file cleared")
			tasks = []Task{}
			SaveTasks("tasks.json", tasks)
		default:
			os.Exit(1)
		}

	default:
		printHelp()
	}
}

func checkArgCount(args []string) error {
	if len(args) < 3 {
		return errors.New("too few arguments")
	}

	return nil
}

func printErr(err error) {
	fmt.Printf("Error: %s\n", err)
}

func printHelp() {
	fmt.Printf(`help			show this message
add "<task-body>"	add a task
del <task-id>		delete a task
done <task-id>		toggle task completion
list 			list all tasks
`)
}
