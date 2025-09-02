package main

import (
	"fmt"
	"task-manager/internal/service"
)

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
			service.AddTask()
		case 2:
			service.ListTask()
		case 3:
			service.DeleteTask()
		case 4:
			service.StatusTask()
		default:
			fmt.Println("---Invalid choice---")
			continue
		}
	}
}
