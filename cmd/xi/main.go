package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/3rr0r-505/Xi/internal/commands"
)

const usage = `
Usage: xi <subcommand> [flags]

Subcommands:
  	add         Add a new task
  	done        Mark a task as done
  	pending     Mark a task as pending
  	delete      Delete a task
  	details     Show task details
  	add-subtask Add a subtask
  	prioritise  Prioritise a task
  	list        List tasks

Run 'xi <subcommand> --help' for subcommand usage.
`

const detailsUsage = `
    [f8b2cd] Design database schema
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
    Status     :   [✔] done
    Created    :   2026-04-02 00:00:00
    Completed  :   2026-04-03 00:00:00
    Priority   :   High
    # SubTasks :   2
        ⤷ 079d4b     Define tables                  0     [✘] pending
        ⤷ 44b069     Set relations                  0     [✘] pending

Note: # SubTasks column shows nested subtasks under each subtask.
      	0 means the subtask has no further children.
`

func main() {
	// define subcommands & flags
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	doneCmd := flag.NewFlagSet("done", flag.ExitOnError)
	pendingCmd := flag.NewFlagSet("pending", flag.ExitOnError)
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	detailsCmd := flag.NewFlagSet("details", flag.ExitOnError)
	addSubTaskCmd := flag.NewFlagSet("add-subtask", flag.ExitOnError)
	prioritiseCmd := flag.NewFlagSet("prioritise", flag.ExitOnError)
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	listDoneCmd := listCmd.Bool("done", false, "list complete task")
	listPendingCmd := listCmd.Bool("pending", false, "list pending task")
	listCompletedOnCmd := listCmd.String("completed-on", "", "filter by date (YYYY-MM-DD)")

	// subcommand help msg
	addCmd.Usage = func() {
		fmt.Println()
		fmt.Println("Usage: xi add <description>")
		fmt.Println("Example: xi add \"Read Effective Go\"")
		fmt.Println()
	}
	doneCmd.Usage = func() {
		fmt.Println()
		fmt.Println("Usage: xi done <task-id-prefix>")
		fmt.Println("Example: xi done f8b2cd")
		fmt.Println()
	}
	pendingCmd.Usage = func() {
		fmt.Println()
		fmt.Println("Usage: xi pending <task-id-prefix>")
		fmt.Println("Example: xi pending f8b2cd")
		fmt.Println()
	}
	deleteCmd.Usage = func() {
		fmt.Println()
		fmt.Println("Usage: xi delete <task-id-prefix>")
		fmt.Println("Example: xi delete f8b2cd")
		fmt.Println()
	}
	detailsCmd.Usage = func() {
		fmt.Println()
		fmt.Println("Usage: xi details <task-id-prefix>")
		fmt.Println("Example: xi details f8b2cd")
		fmt.Print(detailsUsage)
		fmt.Println()
	}
	addSubTaskCmd.Usage = func() {
		fmt.Println()
		fmt.Println("Usage: xi add-subtask <parent-id-prefix> <description>")
		fmt.Println("Example: xi add-subtask f8b2cd \"Login flow\"")
		fmt.Println()
	}
	prioritiseCmd.Usage = func() {
		fmt.Println()
		fmt.Println("Usage: xi prioritise <task-id-prefix>")
		fmt.Println("Example: xi prioritise f8b2cd")
		fmt.Println()
	}
	listCmd.Usage = func() {
		fmt.Println()
		fmt.Println("Usage: xi list [flags]")
		fmt.Println()
		fmt.Println("Flags:")
		listCmd.PrintDefaults()
		fmt.Println()
		fmt.Println("Examples:")
		fmt.Println("  xi list")
		fmt.Println("  xi list --done")
		fmt.Println("  xi list --pending")
		fmt.Println("  xi list --completed-on 2026-04-14")
		fmt.Println()
	}

	// main help msg
	if len(os.Args) < 2 || os.Args[1] == "help" || os.Args[1] == "--help" || os.Args[1] == "-help" || os.Args[1] == "-h" {
		fmt.Print(usage)
		fmt.Printf("\n")
		os.Exit(0)
	}

	// command execution logics
	switch os.Args[1] {
	case "add":
		addCmd.Parse(os.Args[2:])
		args := addCmd.Args() // positional args after flags
		if len(args) < 1 {
			fmt.Println("error: provide a task description")
			os.Exit(1)
		}
		commands.AddTask(args[0])

	case "done":
		doneCmd.Parse(os.Args[2:])
		args := doneCmd.Args()
		if len(args) < 1 {
			fmt.Println("error: provide a task id")
			os.Exit(1)
		}
		commands.DoneTask(args[0])

	case "pending":
		pendingCmd.Parse(os.Args[2:])
		args := pendingCmd.Args()
		if len(args) < 1 {
			fmt.Println("error: provide a task id")
			os.Exit(1)
		}
		commands.PendingTask(args[0])

	case "delete":
		deleteCmd.Parse(os.Args[2:])
		args := deleteCmd.Args()
		if len(args) < 1 {
			fmt.Println("error: provide a task id")
			os.Exit(1)
		}
		commands.DeleteTask(args[0])

	case "details":
		detailsCmd.Parse(os.Args[2:])
		args := detailsCmd.Args()
		if len(args) < 1 {
			fmt.Println("error: provide a task id")
			os.Exit(1)
		}
		commands.TaskDetail(args[0])

	case "add-subtask":
		addSubTaskCmd.Parse(os.Args[2:])
		args := addSubTaskCmd.Args()
		if len(args) < 2 {
			fmt.Println("error: provide a task id & task description")
			os.Exit(1)
		}
		commands.AddSubTask(args[0], args[1])

	case "prioritise":
		prioritiseCmd.Parse(os.Args[2:])
		args := prioritiseCmd.Args()
		if len(args) < 1 {
			fmt.Println("error: provide a task id")
			os.Exit(1)
		}
		commands.TaskPrioritise(args[0])

	case "list":
		listCmd.Parse(os.Args[2:])

		if *listDoneCmd && *listPendingCmd {
			fmt.Println("error: --done and --pending can't be used together")
			os.Exit(1)
		}

		switch {
		case *listDoneCmd:
			commands.ListDoneTasks()
		case *listPendingCmd:
			commands.ListPendingTasks()
		case *listCompletedOnCmd != "":
			commands.ListCompletedOn(*listCompletedOnCmd)
		default:
			commands.ListAllTasks()
		}

	default:
		fmt.Println("Command not found!")
		os.Exit(1)
	}
}
