package cmd

import (
	"fmt"
	"os"
	"strconv"
	"task_tracker/internal"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d", "del", "rm", "remove", "hapus"},
	Short:   "Delete a task",
	Long: `Delete a specific task by its ID 
	Example: 
	task-tracker delete 'ID'
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(os.Args) < 3 {
			return redError("a task ID is required")
		} else if len(os.Args) > 3 {
			return redError("too many arguments for delete command")
		} else if idx, err := strconv.Atoi(args[0]); err != nil {
			return redError("ID must be a positive number")
		} else if err := internal.SearchIndex(idx); err != "" {
			return redError(err)
		} else {
			fmt.Printf(internal.Green+"\n  Task %d deleted successfully"+internal.Rc, idx)
			return internal.DeleteTask(idx)
		}
	},
}
