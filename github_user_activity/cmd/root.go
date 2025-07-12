package cmd

import (
	"github_user_activity/internal"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "github-activity",
	Short: "GitHub User Activity is a CLI tool for fetching a user's recent activity by providing their username",
	Long: `GitHub User Activity is a CLI tool for fetching a user's recent activity by providing their username
	Example:
	github-activity <github-username>`,
	SilenceErrors: true,
	Args:          cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {

		username := args[0]
		activities, err := internal.FetchUserActivity(username)
		if err != nil {
			return err
		}

		err = internal.DisplayUserActivity(username, activities)
		if err != nil {
			return err
		}
		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}
