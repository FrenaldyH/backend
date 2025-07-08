package cmd

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(&addCmd)
}

var addCmd = cobra.Command{
	Use:     "add",
	Aliases: []string{"a", "insert", "append", "tambah"},
	Short:   "Add a task to the task list",
	Long: `Add a task to the task list
	Example:
	task-tracker add 'description'`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(os.Args) < 3 {
			return redError("task description is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		addTask(strings.Join(args[0:], " "))
	},
}

func addTask(newTask string) {
	arrT = append(
		arrT, Task{
			ID:          len(arrT) + 1,
			Description: newTask,
			Status:      "todo",
			CreatedAt:   time.Now().Format(timeLayout),
			UpdatedAt:   ""})
	fmt.Print(green + "\n  Add task successfully\n" + rc)
}
