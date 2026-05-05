package utils_test

import (
	"os"
    "testing"
	"path/filepath"

    "github.com/3rr0r-505/Xi/internal/utils"
    "github.com/3rr0r-505/Xi/internal/task"
)

func TestFindTaskId(t *testing.T) {
    tasks := []task.Task{
        {Id: "abc123def456", Title: "Task one"},
        {Id: "abc789ghi012", Title: "Task two"},
        {Id: "xyz111aaa222", Title: "Task three"},
    }

    // ambiguous prefix
    _, err := utils.FindTaskId(tasks, "abc")
    if err == nil {
        t.Errorf("expected ambiguous error, got nil")
    }

    // not found
    _, err = utils.FindTaskId(tasks, "zzz")
    if err == nil {
        t.Errorf("expected not found error, got nil")
    }

    // valid match
    id, err := utils.FindTaskId(tasks, "xyz")
    if err != nil {
        t.Errorf("expected match, got error: %v", err)
    }
    if id != "xyz111aaa222" {
        t.Errorf("expected xyz111aaa222, got %s", id)
    }
}

func TestLoadEmpty(t *testing.T){
	defer os.Remove(filepath.Join(os.Getenv("HOME"), ".xi", "tasks.json"))
	tasks, err := utils.LoadTask()
	if err != nil {
		t.Fatalf("expected nil error, got: %v", err)
	}
	if len(tasks) != 0 {
		t.Errorf("Expected empty list slice, got error: %v", err)
	}
}

func TestSaveAndLoad(t *testing.T){
	defer os.Remove(filepath.Join(os.Getenv("HOME"), ".xi", "tasks.json"))
	tasks := []task.Task{
        {Id: "abc123def456", Title: "Task one"},
        {Id: "abc789ghi012", Title: "Task two"},
        {Id: "xyz111aaa222", Title: "Task three"},
    }

	err := utils.StoreTask(tasks)
	if err != nil {
		t.Fatalf("Expected nil error, got: %v", err)
	}

	tasks, err = utils.LoadTask()
	if err != nil {
		t.Fatalf("expected nil error, got: %v", err)
	}
	if len(tasks) != 3 {
		t.Errorf("Expected list slice having 3 elements, got: %v", err)
	}
}

func BenchmarkLoadTask(b *testing.B) {
	src, _ := os.ReadFile(filepath.Join(os.Getenv("HOME"), ".xi", "demo_tasks.json"))
	os.WriteFile(filepath.Join(os.Getenv("HOME"), ".xi", "tasks.json"), src, 0644)

    for b.Loop() {
        utils.LoadTask()
    }
}

func BenchmarkStoreTask(b *testing.B){
	tasks, _ := utils.LoadTask()
	os.Remove(filepath.Join(os.Getenv("HOME"), ".xi", "tasks.json"))
	defer os.Remove(filepath.Join(os.Getenv("HOME"), ".xi", "tasks.json"))

    for b.Loop() {
        utils.StoreTask(tasks)
    }
}

