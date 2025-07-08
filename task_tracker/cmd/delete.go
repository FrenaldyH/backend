package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d", "del", "rm", "remove", "hapus"},
	Short:   "Delete a task",
	Long: `Delete a specify task by it's ID
	Example: 
	task-tracker delete 'ID'
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(os.Args) < 3 {
			return redError("task ID is requaired")
		} else if len(os.Args) > 3 {
			return redError("delete argument to long")
		} else if idx, err := strconv.Atoi(args[0]); err != nil {
			return redError("index must be a natural number")
		} else if err := searchIndex(idx); err != "" {
			return redError(err)
		} else {
			return deleteTask(idx)
		}
	},
}

func deleteTask(deletedID int) error {
	arrT[deletedID-1].Status = "deleted"
	fmt.Printf(green+"\n  delete task %d successfully\n"+rc, deletedID)
	return nil
}
