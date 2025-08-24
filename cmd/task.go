package main

import "time"

type Task struct {
	Title       string
	Description string
	Date        time.Time
	Status      bool
}
