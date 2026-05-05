package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/3rr0r-505/Xi/internal/task"
	"github.com/3rr0r-505/Xi/internal/ui"
	"github.com/3rr0r-505/Xi/internal/utils"
)

func countSubtasks(tasks []task.Task) map[string]int {
	subTaskCount := map[string]int{}
	for _, t := range tasks {
		if t.ParentId != "" {
			subTaskCount[t.ParentId]++
		}
	}
	return subTaskCount
}

func printTaskDetails(taskDetail []task.Task, subTaskCountMap map[string]int) {
	// printing header
	fmt.Println()
	fmt.Printf(ui.HalfTab+"[%s] %s\n", taskDetail[0].Id[:6], taskDetail[0].Title)
	fmt.Println(strings.Repeat(ui.Dash, 50))

	// printing main task's details
	if taskDetail[0].Status == task.StatusPending {
		fmt.Printf(ui.HalfTab+"%-10s %-3s [%s] %s\n", "Status", ":", ui.CrossMark, taskDetail[0].Status)
	} else {
		fmt.Printf(ui.HalfTab+"%-10s %-3s [%s] %s\n", "Status", ":", ui.TickMark, taskDetail[0].Status)
	}
	fmt.Printf(ui.HalfTab+"%-10s %-3s %s\n", "Created", ":", taskDetail[0].CreatedOn.Format("2006-01-02 15:04:05"))
	if taskDetail[0].CompletedOn != nil {
		fmt.Printf(ui.HalfTab+"%-10s %-3s %s", "Completed", ":", taskDetail[0].CompletedOn.Format("2006-01-02 15:04:05"))
		fmt.Println()
	} else {
		fmt.Printf(ui.HalfTab+"%-10s %-3s \uFF0D", "Completed", ":")
		fmt.Println()
	}
	if taskDetail[0].Prioritised {
		fmt.Printf(ui.HalfTab+"%-10s %-3s %s", "Priority", ":", "High")
		fmt.Println()
	} else {
		fmt.Printf(ui.HalfTab+"%-10s %-3s %s", "Priority", ":", "Low")
		fmt.Println()
	}
	fmt.Printf(ui.HalfTab+"%-10s %-3s %d\n", "# SubTasks", ":", len(taskDetail)-1)

	// printing subtasks
	for i := 1; i < len(taskDetail); i++ {
		var icon string

		switch taskDetail[i].Status {
		case task.StatusPending:
			icon = ui.CrossMark
		case task.StatusDone:
			icon = ui.TickMark
		default:
			icon = ui.Circle
		}

		fmt.Printf(ui.HalfTab+ui.HalfTab+ui.SubArrow+" %-10s %-30s %-5d [%s] %s\n",
			taskDetail[i].Id[:6],
			taskDetail[i].Title,
			subTaskCountMap[taskDetail[i].Id],
			icon,
			taskDetail[i].Status,
		)
	}

	fmt.Println()
}

func TaskDetail(taskIdPrefix string) {
	// loading existed tasks list
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

	// building complete Task with subtasks list
	var detailedTask []task.Task
	for idx := range tasks {
		if tasks[idx].Id == taskId {
			detailedTask = append(detailedTask, tasks[idx])
		} else if tasks[idx].ParentId == taskId {
			detailedTask = append(detailedTask, tasks[idx])
		}
	}
	if len(detailedTask) == 0 {
		fmt.Fprintln(os.Stderr, "task not found")
		os.Exit(1)
	}
	printTaskDetails(detailedTask, countSubtasks(tasks))
}

func TaskPrioritise(taskIdPrefix string) {
	// loading existed tasks list
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
	state := "prioritised"
	for idx := range tasks {
		if tasks[idx].Id == taskId {
			tasks[idx].Prioritised = !tasks[idx].Prioritised
			if !tasks[idx].Prioritised {
				state = "unprioritised"
			}
			break
		}
	}

	err = utils.StoreTask(tasks)
	if err != nil {
		fmt.Fprintln(os.Stderr, ui.ColorRed+"["+ui.XiUpper+"]"+ui.ColorReset, "error saving task:", err)
		os.Exit(1)
	}
	fmt.Printf(ui.ColorGreen+"["+ui.TickMark+"] %s task having Id [%s]"+ui.ColorReset+"\n", state, taskIdPrefix)
}
