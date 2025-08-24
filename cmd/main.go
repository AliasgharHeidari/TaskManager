package main

import "fmt"

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
