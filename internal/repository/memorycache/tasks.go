package memorycache

import "task-manager/internal/model"

var Tasks *map[int]model.Task

func init() {
	Tasks = &map[int]model.Task{}
}
