package repo

import (
	"encoding/json"
	"os"
	"sync"
	"task-reminder/model"
)

var dataPath = "data/tasks.json"
var mu sync.Mutex

func LoadTasks() ([]model.Task, error) {
	mu.Lock()
	defer mu.Unlock()

	if _, err := os.Stat(dataPath); os.IsNotExist(err) {
		return []model.Task{}, nil
	}

	b, err := os.ReadFile(dataPath)
	if err != nil {
		return nil, err
	}

	var tasks []model.Task
	err = json.Unmarshal(b, &tasks)
	return tasks, err
}

func SaveTasks(tasks []model.Task) error {
	mu.Lock()
	defer mu.Unlock()

	b, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}

	os.MkdirAll("data", 0755)
	return os.WriteFile(dataPath, b, 0644)
}
