package cmd

import (
	"strconv"
	"strings"
	"task_tracker/internal"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(markTodoCmd)
	rootCmd.AddCommand(markInProgressCmd)
	rootCmd.AddCommand(markDoneCmd)
	rootCmd.AddCommand(updateCmd)
}

var (
	updateCmd = &cobra.Command{
		Use:     "update",
		Aliases: []string{"u"},
		Short:   "Update a task",
		Long: `Update a task by providing the task ID and the new status
		Example:
		task-tracker update 'ID' 'new description'
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return redError("task id is required")
			} else if len(args) < 2 {
				return redError("new description is requaired")
			} else if idx, err := strconv.Atoi(args[0]); err != nil {
				return redError("index must be a natural number")
			} else if err := internal.SearchIndex(idx); err != "" {
				return redError(err)
			} else {
				internal.UpdateDescription(idx, strings.Join(args[1:], " "))
				return nil
			}

		},
	}
	markTodoCmd = &cobra.Command{
		Use:     "mark-todo",
		Aliases: []string{"mt", "todo"},
		Short:   "Mark a task as todo",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := internal.UpdateStatusTask(args, "todo"); err != "" {
				return redError(err)
			}
			return nil
		},
	}
	markInProgressCmd = &cobra.Command{
		Use:     "mark-in-progress",
		Aliases: []string{"mip", "in-progress", "progress"},
		Short:   "Mark a task as in-progress",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := internal.UpdateStatusTask(args, "in-progress"); err != "" {
				return redError(err)
			}
			return nil
		},
	}
	markDoneCmd = &cobra.Command{
		Use:     "mark-done",
		Aliases: []string{"md", "done"},
		Short:   "Mark a task as done",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := internal.UpdateStatusTask(args, "done"); err != "" {
				return redError(err)
			}
			return nil
		},
	}
)
