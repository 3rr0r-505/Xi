package utils

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/3rr0r-505/Xi/internal/task"
)

func LoadTask() ([]task.Task, error) {
	// finding home dir of user
	homedir, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	// checking for either json file exists or not
	jsonPath := filepath.Join(homedir, ".xi", "tasks.json")
	_, err = os.Stat(jsonPath)
	if os.IsNotExist(err) {
		return []task.Task{}, nil
	}

	// reading the data from json file
	var data []task.Task
	taskData, err := os.ReadFile(jsonPath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(taskData, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
