package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/3rr0r-505/Xi/internal/task"
	"github.com/3rr0r-505/Xi/internal/ui"
	"github.com/3rr0r-505/Xi/internal/utils"
)

// helper functions

func showHeader() {
	fmt.Printf(ui.HalfTab+"%-5s %-10s %-50s %-11s %-13s %s\n", "#", "ID", "Title", "Status", "# SubTasks", "Priority")
	fmt.Println(strings.Repeat(ui.Dash, 109))
}

func printTask(idx int, tsk task.Task, noOfSubtasks int) {
	var prio string

	// shortening long titles
	title := tsk.Title
	if len(title) > 45 {
		title = title[:43] + "..."
	}

	// checking for prioritised task
	if tsk.Prioritised {
		prio = ui.UpArrow
	}

	var icon string
	switch tsk.Status {
	case task.StatusPending:
		icon = ui.CrossMark
	case task.StatusDone:
		icon = ui.TickMark
	default:
		icon = ui.Circle
	}
	fmt.Printf(ui.HalfTab+"%-5d %-10s %-50s %-1s %-10s %5d %11s\n", idx, tsk.Id[:6], title, icon, tsk.Status, noOfSubtasks, prio)
}

func listTasks(tasks []task.Task, subTaskCountMap map[string]int, filter func(task.Task) bool) {
	idx := 0
	for _, tsk := range tasks {
		if filter(tsk) {
			idx++
			printTask(idx, tsk, subTaskCountMap[tsk.Id])
		}
	}
}

// handler functions
func ListAllTasks() {
	// loading task list
	tasks, err := utils.LoadTask()
	if err != nil {
		fmt.Fprintln(os.Stderr, ui.ColorRed+"["+ui.XiUpper+"]"+ui.ColorReset, "error loading task:", err)
		os.Exit(1)
	}

	// making the header of the tasks
	fmt.Println()
	showHeader()

	// printing tasks
	listTasks(tasks, countSubtasks(tasks), func(t task.Task) bool { return t.ParentId == "" })
	fmt.Println()
}

func ListDoneTasks() {
	// loading task list
	tasks, err := utils.LoadTask()
	if err != nil {
		fmt.Fprintln(os.Stderr, ui.ColorRed+"["+ui.XiUpper+"]"+ui.ColorReset, "error loading task:", err)
		os.Exit(1)
	}

	// making the header of the tasks
	fmt.Println()
	showHeader()

	// printing tasks
	listTasks(tasks, countSubtasks(tasks), func(t task.Task) bool { return t.ParentId == "" && t.Status == task.StatusDone })
	fmt.Println()
}

func ListPendingTasks() {
	// loading task list
	tasks, err := utils.LoadTask()
	if err != nil {
		fmt.Fprintln(os.Stderr, ui.ColorRed+"["+ui.XiUpper+"]"+ui.ColorReset, "error loading task:", err)
		os.Exit(1)
	}

	// making the header of the tasks
	fmt.Println()
	showHeader()

	// printing tasks
	listTasks(tasks, countSubtasks(tasks), func(t task.Task) bool { return t.ParentId == "" && t.Status == task.StatusPending })
	fmt.Println()
}

func ListCompletedOn(date string) {
	// loading task list
	tasks, err := utils.LoadTask()
	if err != nil {
		fmt.Fprintln(os.Stderr, ui.ColorRed+"["+ui.XiUpper+"]"+ui.ColorReset, "error loading task:", err)
		os.Exit(1)
	}

	// making the header of the tasks
	fmt.Println()
	showHeader()

	// printing tasks
	listTasks(tasks, countSubtasks(tasks), func(t task.Task) bool {
		return t.ParentId == "" && t.CompletedOn != nil && date == t.CompletedOn.Format("2006-01-02")
	})
	fmt.Println()
}
