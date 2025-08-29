package model

import "time"

type Task struct {
	Title       string
	Description string
	Date        time.Time
	Status      bool
}
