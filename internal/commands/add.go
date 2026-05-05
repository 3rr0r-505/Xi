package commands

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"os"
	"time"

	"github.com/3rr0r-505/Xi/internal/task"
	"github.com/3rr0r-505/Xi/internal/ui"
	"github.com/3rr0r-505/Xi/internal/utils"
)

func hashSHA1(rawInput string) string {
	hasher := sha1.New()
	hasher.Write([]byte(rawInput))
	hash := hasher.Sum(nil)
	hashString := hex.EncodeToString(hash[:])
	return hashString
}

func AddTask(desc string) {
	//loading existed tasks list
	tasks, err := utils.LoadTask()
	if err != nil {
		fmt.Fprintln(os.Stderr, ui.ColorRed+"["+ui.XiUpper+"]"+ui.ColorReset, "error loading task:", err)
		os.Exit(1)
	}

	// calculating id & time for new task
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	hashId := hashSHA1(desc + formattedTime)

	// creating new task
	var newTask task.Task
	newTask.Id = hashId
	newTask.Title = desc
	newTask.Status = task.StatusPending
	newTask.CreatedOn = currentTime

	// updating the tasks list
	updatedTasks := append(tasks, newTask)

	err = utils.StoreTask(updatedTasks)
	if err != nil {
		fmt.Fprintln(os.Stderr, ui.ColorRed+"["+ui.XiUpper+"]"+ui.ColorReset, "error saving task:", err)
		os.Exit(1)
	}
	fmt.Printf(ui.ColorGreen+"["+ui.TickMark+"] Task added [%s]"+ui.ColorReset+"\n", newTask.Id[:6])
}

func AddSubTask(parentId string, desc string) {
	// loading existed tasks list
	tasks, err := utils.LoadTask()
	if err != nil {
		fmt.Fprintln(os.Stderr, ui.ColorRed+"["+ui.XiUpper+"]"+ui.ColorReset, "loading task:", err)
		os.Exit(1)
	}

	// getting parent task
	parentTaskId, err := utils.FindTaskId(tasks, parentId)
	if err != nil {
		fmt.Fprintln(os.Stderr, ui.ColorRed+"["+ui.XiUpper+"]"+ui.ColorReset, err)
		os.Exit(1)
	}

	// calculating id & time for new task
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	hashId := hashSHA1(desc + formattedTime)

	// creating new task
	var newSubTask task.Task
	newSubTask.Id = hashId
	newSubTask.ParentId = parentTaskId
	newSubTask.Title = desc
	newSubTask.Status = task.StatusPending
	newSubTask.CreatedOn = currentTime

	// updating the tasks list
	updatedTasks := append(tasks, newSubTask)

	err = utils.StoreTask(updatedTasks)
	if err != nil {
		fmt.Fprintln(os.Stderr, ui.ColorRed+"["+ui.XiUpper+"]"+ui.ColorReset, "error saving task:", err)
		os.Exit(1)
	}
	fmt.Printf(ui.ColorGreen+"["+ui.TickMark+"] SubTask added [%s]"+ui.ColorReset+"\n", newSubTask.Id[:6])
}
