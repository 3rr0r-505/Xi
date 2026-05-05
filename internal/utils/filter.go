package utils

import (
	"fmt"
	"strings"

	"github.com/3rr0r-505/Xi/internal/task"
)

// finding task ID(SHA-1 hash) by the prefix
func FindTaskId(tasks []task.Task, prefix string) (string, error) {
	var matches []task.Task
	for _, t := range tasks {
		if strings.HasPrefix(t.Id, prefix) {
			matches = append(matches, t)
		}
	}

	if len(matches) == 0 {
		return "", fmt.Errorf("no task found with id starting: %s", prefix)
	}

	if len(matches) > 1 {
		return "", fmt.Errorf("ambiguous id, %d tasks match — use more characters", len(matches))
	}

	return matches[0].Id, nil
}

// finding task index in Task slice by the prefix
func FindTaskIndex(tasks []task.Task, prefix string) (int, string, error) {
	var matches []task.Task
	var index []int
	for i, t := range tasks {
		if strings.HasPrefix(t.Id, prefix) {
			matches = append(matches, t)
			index = append(index, i)
		}
	}

	if len(matches) == 0 {
		return -1, "", fmt.Errorf("no task found with id starting: %s", prefix)
	}

	if len(matches) > 1 {
		return -1, "", fmt.Errorf("ambiguous id, %d tasks match — use more characters", len(matches))
	}

	return index[0], matches[0].Id, nil
}
