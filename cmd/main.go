package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	Title       string
	Description string
	Date        time.Time
	Status      bool
}

var task_map = map[int]Task{}
var taskID = 0

func addTask() {
	loadTask()
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	reader.ReadString('\n')
	var task Task
	fmt.Printf("Enter the name of task: ")
	title, _ := reader.ReadString('\n')
	task.Title = strings.TrimSpace(title)

	fmt.Printf("Enter description: ")
	desc, _ := reader.ReadString('\n')
	task.Description = strings.TrimSpace(desc)
	task.Date = time.Now()
	task.Status = false
	taskID = rand.Intn(9000) + 1000
	task_map[taskID] = task
	saveTasks()
	fmt.Println("-------Task succeccfuly addedp-------")
}

func listTask() {
	loadTask()
	if len(task_map) == 0 {
		fmt.Println("-----------------------------------------------")
		fmt.Println("----------------No task avalible---------------")
	}
	fmt.Println("-----------------------------------------------")
	for id, task := range task_map {
		fmt.Println("ID:", id)
		fmt.Println("Title:", task.Title)
		fmt.Println("Description:", task.Description)
		fmt.Println("Date-modified:", task.Date.Format("2006-01-02 15:04:05"))
		fmt.Println("Status:", task.Status)
		fmt.Println("-----------------------------------------------")
	}

}

func deleteTask() {
	loadTask()
	var id int
	var err error
	for {
		var input string
		fmt.Printf("Enter task ID: ")
		fmt.Scan(&input)
		id, err = strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid ID, Please enter a number")
			continue
		}
		break
	}

	_, ok := task_map[id]
	if ok {
		delete(task_map, id)
		fmt.Println("---Task successfuly deleted---")
	} else {
		fmt.Println("Invalid ID")
	}
	saveTasks()
}

func statusTask() {
	loadTask()
	var task Task
	var input string
	fmt.Printf("Enter task ID: ")
	fmt.Scan(&input)
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("---Invalid ID---")
	}
	_, exists := task_map[id]
	if exists {
		task.Status = true
		task_map[id] = task
		fmt.Println("Task status changed to true")
	} else {
		fmt.Println("Task not found")
	}
	saveTasks()
}

func saveTasks() {
	data, err := json.MarshalIndent(task_map, "", "  ")
	if err != nil {
		fmt.Println("saving failed: ", err)
		return
	}
	os.WriteFile("tasks.json", data, 0644)
}

func loadTask() {
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		fmt.Println("Loading failed: ", err)
		return
	}

	json.Unmarshal(file, &task_map)

}

func main() {
	var choice int
	for {
		fmt.Println("1.Add")
		fmt.Println("2.List")
		fmt.Println("3.Delete")
		fmt.Println("4.Complete")
		fmt.Printf("Choose Option: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			addTask()
		case 2:
			listTask()
		case 3:
			deleteTask()
		case 4:
			statusTask()
		default:
			fmt.Println("---Invalid choice---")
			continue
		}
	}

}
