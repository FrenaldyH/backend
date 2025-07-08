package cmd

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

var (
	arrT       []Task
	timeLayout = "2006-01-02 15:04 MST"
)

const DB = "tasks.json"

func searchIndex(idxTarget int) string {
	left, right := 1, arrT[len(arrT)-1].ID
	for left <= right {
		mid := (left + right) / 2
		if idxTarget > mid {
			left = mid + 1
		} else if idxTarget < mid {
			right = mid - 1
		} else {
			if arrT[mid-1].Status == "deleted" {
				return "task has been deleted"
			} else {
				return ""
			}
		}
	}
	return "index not found"
}

func loadTask() string {
	if _, err := os.Stat(DB); err != nil {
		fmt.Print(yellow + "\n  File database does not exist. Creating file...\n" + rc)
		os.WriteFile(DB, []byte("[]"), 0644)
	} else if readData, err := os.ReadFile(DB); err != nil {
		return "cannot read the data base file"
	} else if err := json.Unmarshal([]byte(readData), &arrT); err != nil {
		return "cannot unmarshaling data base"
	}
	return ""
}

func saveTask() string {
	if newData, err := json.MarshalIndent(arrT, "", " "); err != nil {
		return "cannot marshaling JSON"
	} else if err := os.WriteFile(DB, newData, 0644); err != nil {
		return "cannot writing the data base file"
	}
	return ""
}
