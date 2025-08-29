package filestorage

import (
	"encoding/json"
	"fmt"
	"os"
	"task-manager/internal/repository/memorycache"
)

func SaveTasks() {
	data, err := json.MarshalIndent(memorycache.Tasks, "", "  ")
	if err != nil {
		fmt.Println("saving failed: ", err)
		return
	}
	os.WriteFile("Save-json/tasks.json", data, 0644)
}

func LoadTask() {
	file, err := os.ReadFile("Save-json/tasks.json")
	if err != nil {
		fmt.Println("Loading failed: ", err)
		return
	}

	if err := json.Unmarshal(file, &memorycache.Tasks); err != nil {
		fmt.Println("Unmarshalling failed: ", err)
		return
	}

}
