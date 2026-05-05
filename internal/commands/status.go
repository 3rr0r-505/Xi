package commands

import (
	"fmt"
	"os"
	"time"

	"github.com/3rr0r-505/Xi/internal/task"
	"github.com/3rr0r-505/Xi/internal/ui"
	"github.com/3rr0r-505/Xi/internal/utils"
)

func DoneTask(taskIdPrefix string) {
	//loading existed tasks list
	tasks, err := utils.LoadTask()
	if err != nil {
		fmt.Fprintln(os.Stderr, ui.ColorRed+"["+ui.XiUpper+"]"+ui.ColorReset, "error loading task:", err)
		os.Exit(1)
	}

	// getting task id & current time
	taskId, err := utils.FindTaskId(tasks, taskIdPrefix)
	if err != nil {
		fmt.Fprintln(os.Stderr, ui.ColorRed+"["+ui.XiUpper+"]"+ui.ColorReset, err)
		os.Exit(1)
	}
	currentTime := time.Now()

	// Updating Tasks list
	for idx := range tasks {
		if tasks[idx].Id == taskId {
			tasks[idx].Status = task.StatusDone
			tasks[idx].CompletedOn = &currentTime
			break
		}
	}

	err = utils.StoreTask(tasks)
	if err != nil {
		fmt.Fprintln(os.Stderr, ui.ColorRed+"["+ui.XiUpper+"]"+ui.ColorReset, "error saving task:", err)
		os.Exit(1)
	}
	fmt.Printf(ui.ColorGreen+"["+ui.TickMark+"] Status Updated for task Id [%s]"+ui.ColorReset+"\n", taskIdPrefix)
}

func PendingTask(taskIdPrefix string) {
	//loading existed tasks list
	tasks, err := utils.LoadTask()
	if err != nil {
		fmt.Fprintln(os.Stderr, ui.ColorRed+"["+ui.XiUpper+"]"+ui.ColorReset, "error loading task:", err)
		os.Exit(1)
	}

	// getting task id
	taskId, err := utils.FindTaskId(tasks, taskIdPrefix)
	if err != nil {
		fmt.Fprintln(os.Stderr, ui.ColorRed+"["+ui.XiUpper+"]"+ui.ColorReset, err)
		os.Exit(1)
	}

	// Updating Tasks list
	for idx := range tasks {
		if tasks[idx].Id == taskId {
			tasks[idx].Status = task.StatusPending
			tasks[idx].CompletedOn = nil
			break
		}
	}

	err = utils.StoreTask(tasks)
	if err != nil {
		fmt.Fprintln(os.Stderr, ui.ColorRed+"["+ui.XiUpper+"]"+ui.ColorReset, "error saving task:", err)
		os.Exit(1)
	}
	fmt.Printf(ui.ColorGreen+"["+ui.TickMark+"] Status Updated for task Id [%s]"+ui.ColorReset+"\n", taskIdPrefix)
}
