package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(cmdList)
}

var cmdList = &cobra.Command{
	Use:     "list",
	Aliases: []string{"ls"},
	Short:   "List all tasks",
	Long: `command to list all tasks and you can filter it by status
	Example: 
	task-tracker list todo
	task-tracker list in-progress
	task-tracker list done
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		switch len(args) {
		case 0:
			return listTasks("all")
		case 1:
			return listTasks(args[0])
		default:
			return redError("list arguments to long")
		}
	},
}

func listTasks(flag string) error {
	maxWidths := map[string]int{
		"Id":       len("Id"),
		"Task":     len("Task"),
		"CreateAt": len("CreateAt"),
		"UpdateAt": len("UpdateAt"),
	}

	var rows [][]string
	for _, val := range arrT {
		if (flag != "all" && val.Status != flag) || val.Status == "deleted" {
			continue
		}

		indicator := "[ ]"
		switch val.Status {
		case "done":
			indicator = "[✓]"
		case "in-progress":
			indicator = "[○]"
		}

		titleWithIndicator := indicator + val.Description
		rows = append(rows, []string{
			strconv.Itoa(val.ID),
			titleWithIndicator,
			val.CreatedAt,
			val.UpdatedAt,
		})

		if w := len(rows[len(rows)-1][0]); w > maxWidths["Id"] {
			maxWidths["Id"] = w
		}
		if w := len(rows[len(rows)-1][1]); w > maxWidths["Task"] {
			maxWidths["Task"] = w
		}
		if w := len(rows[len(rows)-1][2]); w > maxWidths["CreateAt"] {
			maxWidths["CreateAt"] = w
		}
		if w := len(rows[len(rows)-1][3]); w > maxWidths["UpdateAt"] {
			maxWidths["UpdateAt"] = w
		}
	}

	keys := []string{"Id", "Task", "CreateAt", "UpdateAt"}

	printBorder := func() {
		fmt.Print("+")
		for _, k := range keys {
			fmt.Print(strings.Repeat("-", maxWidths[k]+2) + "+")
		}
		fmt.Println()
	}

	printBorder()

	fmt.Print("|")
	for _, k := range keys {
		fmt.Printf(" %-*s |", maxWidths[k], k)
	}
	fmt.Println()

	printBorder()

	for _, row := range rows {
		fmt.Print("|")
		for i, k := range keys {
			fmt.Printf(" %-*s |", maxWidths[k], row[i])
		}
		fmt.Println()
	}

	printBorder()

	return nil
}
