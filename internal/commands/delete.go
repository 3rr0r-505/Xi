package commands

import (
	"fmt"
	"os"

	"github.com/3rr0r-505/Xi/internal/task"
	"github.com/3rr0r-505/Xi/internal/ui"
	"github.com/3rr0r-505/Xi/internal/utils"
)

func DeleteTask(taskIdPrefix string) {
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
	var updatedTasks []task.Task
	for idx := range tasks {
		if tasks[idx].Id != taskId && tasks[idx].ParentId != taskId {
			updatedTasks = append(updatedTasks, tasks[idx])
		}
	}

	err = utils.StoreTask(updatedTasks)
	if err != nil {
		fmt.Fprintln(os.Stderr, ui.ColorRed+"["+ui.XiUpper+"]"+ui.ColorReset, "error saving task:", err)
		os.Exit(1)
	}
	fmt.Printf(ui.ColorGreen+"["+ui.TickMark+"] deleted task having Id [%s]"+ui.ColorReset+"\n", taskIdPrefix)
}
