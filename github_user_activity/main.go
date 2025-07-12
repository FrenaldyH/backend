package main

import (
	"fmt"
	"github_user_activity/cmd"
	"github_user_activity/style"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(style.Error.Render("Error: " + err.Error()))
	}
}
