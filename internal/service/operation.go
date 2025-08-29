package service

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"task-manager/internal/model"
	filestorage "task-manager/internal/repository/file"
	"task-manager/internal/repository/memorycache"
	"time"
)

var taskID = 0

func AddTask() {
	filestorage.LoadTask()
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	reader.ReadString('\n')
	var task model.Task
	fmt.Printf("Enter the name of task: ")
	title, _ := reader.ReadString('\n')
	task.Title = strings.TrimSpace(title)

	fmt.Printf("Enter description: ")
	desc, _ := reader.ReadString('\n')
	task.Description = strings.TrimSpace(desc)
	task.Date = time.Now()
	task.Status = false
	taskID = rand.Intn(9000) + 1000
	(*memorycache.Tasks)[taskID] = task
	filestorage.SaveTasks()
	fmt.Println("-------Task successfully added-------")
}

func ListTask() {
	filestorage.LoadTask()
	if len(*memorycache.Tasks) == 0 {
		fmt.Println("-----------------------------------------------")
		fmt.Println("----------------No task avalible---------------")
	}
	fmt.Println("-----------------------------------------------")
	for id, task := range *memorycache.Tasks {
		fmt.Println("ID:", id)
		fmt.Println("Title:", task.Title)
		fmt.Println("Description:", task.Description)
		fmt.Println("Date-modified:", task.Date.Format("2006-01-02 15:04:05"))
		fmt.Println("Status:", task.Status)
		fmt.Println("-----------------------------------------------")
	}
}

func DeleteTask() {
	filestorage.LoadTask()
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

	_, ok := (*memorycache.Tasks)[id]
	if ok {
		delete(*memorycache.Tasks, id)
		fmt.Println("---Task successfully deleted---")
	} else {
		fmt.Println("Invalid ID")
	}
	filestorage.SaveTasks()
}

func StatusTask() {
	filestorage.LoadTask()
	var input string
	fmt.Printf("Enter task ID: ")
	fmt.Scan(&input)
	id, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("---Invalid ID---")
	}
	task, exists := (*memorycache.Tasks)[id]
	if exists {
		task.Status = true
		(*memorycache.Tasks)[id] = task
		fmt.Println("Task status changed to true")
	} else {
		fmt.Println("Task not found")
	}
	filestorage.SaveTasks()
}
