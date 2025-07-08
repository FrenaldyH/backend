package cmd

import (
	"errors"
	"fmt"
	"os"
	"task_tracker/internal"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "CLI tool for managing your tasks",
	Long:  `CLI tool for managing your tasks. It allows you to add, delete, and list tasks`,
}

func redError(err string) error {
	return errors.New(internal.Red + "\n  " + err + internal.Rc)
}

func Execute() error {
	if err := internal.LoadTask(); err != "" {
		fmt.Fprintln(os.Stderr, redError(err))
		os.Exit(1)
	}

	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

	if err := internal.SaveTask(); err != "" {
		fmt.Fprintln(os.Stderr, redError(err))
		os.Exit(1)
	}
	return nil
}
