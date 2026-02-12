package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Task struct {
	Text string `json:"text"`
	Done bool   `json:"done"`
}

func ToggleTask(taskList []Task, index int) ([]Task, error) {
	if index >= len(taskList) || index < 0 {
		return taskList, errors.New("invalid index")
	}

	taskList[index].Done = !taskList[index].Done
	return taskList, nil

}

func DeleteTask(taskList []Task, index int) ([]Task, error) {
	if index >= len(taskList) || index < 0 {
		return taskList, errors.New("invalid index")
	}

	newList := []Task{}

	for i, task := range taskList {
		if i != index {
			newList = append(newList, task)
		}
	}

	return newList, nil
}

func LoadTasks(fileName string, taskList *[]Task) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
	} else {
		json.Unmarshal(data, taskList)
	}
}

func SaveTasks(fileName string, taskList []Task) {
	data, err := json.Marshal(taskList)
	if err != nil {
		printErr(err)
	} else {
		err = os.WriteFile(fileName, data, 0644)
		if err != nil {
			printErr(err)
		}
	}
}

func ListTasks(taskList []Task) error {
	if len(taskList) == 0 {
		return errors.New("task list empty")
	}
	for i, task := range taskList {
		if task.Done {
			fmt.Print("[X] ")
		} else {
			fmt.Print("[ ] ")
		}

		fmt.Printf("#%d %s\n", i+1, task.Text)
	}

	return nil
}
