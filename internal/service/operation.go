package service

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	filestorage "task-manager/internal/repository/file"
	"time"
)

var taskID = 0

func AddTask() {
	filestorage.LoadTask()
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
	fmt.Println("-------Task succeccfuly added-------")
}

func ListTask() {
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

func DeleteTask() {
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

func StatusTask() {
	loadTask()
	var input string
	fmt.Printf("Enter task ID: ")
	fmt.Scan(&input)
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("---Invalid ID---")
	}
	task, exists := task_map[id]
	if exists {
		task.Status = true
		task_map[id] = task
		fmt.Println("Task status changed to true")
	} else {
		fmt.Println("Task not found")
	}
	saveTasks()
}
