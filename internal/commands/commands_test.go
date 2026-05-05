package commands_test

import (
	"os"
	"testing"

	"github.com/3rr0r-505/Xi/internal/commands"
	"github.com/3rr0r-505/Xi/internal/task"
)

func suppressStdout() func() {
	nullFile, _ := os.OpenFile(os.DevNull, os.O_WRONLY, 0)
	old := os.Stdout
	os.Stdout = nullFile
	return func() {
		os.Stdout = old
		nullFile.Close()
	}
}

func countSubtasks(tasks []task.Task) map[string]int {
	subTaskCount := map[string]int{}
	for _, t := range tasks {
		if t.ParentId != "" {
			subTaskCount[t.ParentId]++
		}
	}
	return subTaskCount
}

// testing helper functions of commands
func TestCountSubtasks(t *testing.T) {
	tasks := []task.Task{
		{Id: "parent1", ParentId: ""},
		{Id: "child1", ParentId: "parent1"},
		{Id: "child2", ParentId: "parent1"},
		{Id: "parent2", ParentId: ""},
	}

	subTaskCountMap := countSubtasks(tasks)
	if subTaskCountMap["parent1"] != 2 {
		t.Errorf("expected 2 subtasks, got: %d", subTaskCountMap["parent1"])
	}
	if subTaskCountMap["parent2"] != 0 {
		t.Errorf("expected 0 subtasks, got: %d", subTaskCountMap["parent2"])
	}
}

// benchmarking commands
func BenchmarkAddTask(b *testing.B) {
	defer suppressStdout()()
	for b.Loop() {
		commands.AddTask("Testing Benchmarks")
	}
}

func BenchmarkDoneTask(b *testing.B) {
	defer suppressStdout()()
	for b.Loop() {
		commands.DoneTask("e829f5")
	}
}

func BenchmarkPendingTask(b *testing.B) {
	defer suppressStdout()()
	for b.Loop() {
		commands.PendingTask("4efe75")
	}
}

func BenchmarkTaskDetail(b *testing.B) {
	defer suppressStdout()()
	for b.Loop() {
		commands.TaskDetail("f51f193")
	}
}

func BenchmarkAddSubTask(b *testing.B) {
	defer suppressStdout()()
	for b.Loop() {
		commands.AddSubTask("acf39c", "Domain Purchase")
	}
}

func BenchmarkTaskPrioritise(b *testing.B) {
	defer suppressStdout()()
	for b.Loop() {
		commands.TaskPrioritise("acf39c")
	}
}

func BenchmarkListDoneTasks(b *testing.B) {
	defer suppressStdout()()
	for b.Loop() {
		commands.ListDoneTasks()
	}
}

func BenchmarkListPendingTasks(b *testing.B) {
	defer suppressStdout()()
	for b.Loop() {
		commands.ListPendingTasks()
	}
}

func BenchmarkListCompletedOn(b *testing.B) {
	defer suppressStdout()()
	for b.Loop() {
		commands.ListCompletedOn("2026-04-14")
	}
}

func BenchmarkListAllTasks(b *testing.B) {
	defer suppressStdout()()
	for b.Loop() {
		commands.ListAllTasks()
	}
}
