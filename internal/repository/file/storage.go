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
	os.WriteFile("data/tasks.json", data, 0755)
}

func LoadTask() {
	file, err := os.ReadFile("data/tasks.json")
	if err != nil {
		fmt.Println("Loading failed: ", err)
		return
	}

	if err := json.Unmarshal(file, &memorycache.Tasks); err != nil {
		fmt.Println("Unmarshalling failed: ", err)
		return
	}

}
