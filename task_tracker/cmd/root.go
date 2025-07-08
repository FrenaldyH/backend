package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

const (
	rc     = "\x1B[0m\n"
	red    = "\x1B[31m"
	green  = "\x1B[32m"
	yellow = "\x1B[33m"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "CLI tool for managing your tasks",
	Long:  `CLI tool for managing your tasks. It allows you to add, delete, and list tasks`,
}

func redError(err string) error {
	return errors.New(red + "\n  " + err + rc)
}

func Execute() error {
	if err := loadTask(); err != "" {
		return redError(err)
	} else if err := rootCmd.Execute(); err != nil {
		return err
	} else if err := saveTask(); err != "" {
		return redError(err)
	}
	return nil
}
