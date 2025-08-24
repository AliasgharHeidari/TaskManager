package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var task_map = map[int]Task{}

func saveTasks() {
	data, err := json.MarshalIndent(task_map, "", "  ")
	if err != nil {
		fmt.Println("saving failed: ", err)
		return
	}
	os.WriteFile("Save-json/tasks.json", data, 0644)
}

func loadTask() {
	file, err := os.ReadFile("Save-json/tasks.json")
	if err != nil {
		fmt.Println("Loading failed: ", err)
		return
	}

	json.Unmarshal(file, &task_map)
}
