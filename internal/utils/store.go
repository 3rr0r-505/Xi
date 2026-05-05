package utils

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/3rr0r-505/Xi/internal/task"
)

func checkFile() (string, error) {
	// finding home dir of user
	homedir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	// checking for whether json file exists or not
	jsonPath := filepath.Join(homedir, ".xi", "tasks.json")
	_, err = os.Stat(jsonPath)
	if err == nil {
		return jsonPath, nil
	}

	// checking if file exists but any other problem occurs
	if !os.IsNotExist(err) {
		return "", err
	}

	// create the dir
	err = os.MkdirAll(filepath.Dir(jsonPath), 0755)
	if err != nil {
		return "", err
	}

	// create the json file
	file, err := os.Create(jsonPath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// initialise empty json array
	_, err = file.WriteString("[]")
	if err != nil {
		return "", err
	}

	return jsonPath, nil
}

func StoreTask(updatedTaskList []task.Task) error {
	jsonPath, err := checkFile()
	if err != nil {
		return err
	}

	// formatting the json
	updatedJSON, err := json.MarshalIndent(updatedTaskList, "", "\t")
	if err != nil {
		return err
	}
	err = os.WriteFile(jsonPath, updatedJSON, 0644)
	if err != nil {
		return err
	}

	return nil
}
