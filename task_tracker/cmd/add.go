package cmd

import (
	"fmt"
	"os"
	"strings"
	"task_tracker/internal"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a", "insert", "append", "tambah"},
	Short:   "Add a task to the task list",
	Long: `Add a task to the task list
	Example:
	task-tracker add 'description'`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(os.Args) < 3 {
			return redError("a task description is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		internal.AddTask(strings.Join(args[0:], " "))
		fmt.Print(internal.Green + "\n  Task added successfully\n" + internal.Rc)
	},
}
