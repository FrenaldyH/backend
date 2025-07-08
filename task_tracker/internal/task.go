package internal

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"time"
)

const (
	Rc     = "\x1B[0m\n"
	Red    = "\x1B[31m"
	Green  = "\x1B[32m"
	Yellow = "\x1B[33m"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

var (
	ArrT       []Task
	timeLayout = "2006-01-02 15:04 MST"
)

const DB = "tasks.json"

func SearchIndex(idxTarget int) string {
	left, right := 1, ArrT[len(ArrT)-1].ID
	for left <= right {
		mid := (left + right) / 2
		if idxTarget > mid {
			left = mid + 1
		} else if idxTarget < mid {
			right = mid - 1
		} else {
			if ArrT[mid-1].Status == "deleted" {
				return "task has been deleted"
			} else {
				return ""
			}
		}
	}
	return "index not found"
}

func LoadTask() string {
	if _, err := os.Stat(DB); err != nil {
		fmt.Print(Yellow + "\n  Database file not found. Creating file...\n" + Rc)
		os.WriteFile(DB, []byte("[]"), 0644)
	} else if readData, err := os.ReadFile(DB); err != nil {
		return "cannot read the database file"
	} else if err := json.Unmarshal([]byte(readData), &ArrT); err != nil {
		return "cannot unmarshal database"
	}
	return ""
}

func SaveTask() string {
	if newData, err := json.MarshalIndent(ArrT, "", " "); err != nil {
		return "cannot marshal JSON"
	} else if err := os.WriteFile(DB, newData, 0644); err != nil {
		return "cannot write to the database file"
	}
	return ""
}

func UpdateStatusTask(args []string, newStatus string) string {
	if len(args) == 0 {
		return "task ID is required"
	}

	if len(args) > 1 {
		return "mark argument to long"
	} else if idx, err := strconv.Atoi(args[0]); err != nil {
		return "index must be a natural number"
	} else if err := SearchIndex(idx); err != "" {
		return err
	} else {
		ArrT[idx-1].Status = newStatus
		fmt.Print(Green + "\n  Mark task successfully\n" + Rc)
		return ""
	}
}

func UpdateDescription(idx int, newDescription string) {
	ArrT[idx-1].Description = newDescription
	ArrT[idx-1].UpdatedAt = time.Now().Format(timeLayout)
	fmt.Print(Green + "\n Update task successfully\n" + Rc)
}

func AddTask(newTask string) {
	ArrT = append(
		ArrT, Task{
			ID:          len(ArrT) + 1,
			Description: newTask,
			Status:      "todo",
			CreatedAt:   time.Now().Format(timeLayout),
			UpdatedAt:   ""})
}

func DeleteTask(deletedID int) error {
	ArrT[deletedID-1].Status = "deleted"
	return nil
}
