package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"github_user_activity/style"
	"io"
	"net/http"
	"time"
)

type GitHubActivity struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
	CreatedAt string `json:"created_at"`
	Payload   struct {
		Size int64 `json:"size"`
	} `json:"payload"`
}

func FetchUserActivity(username string) ([]GitHubActivity, error) {
	response, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/events", username))
	if err != nil {
		return nil, fmt.Errorf("network error: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode == 404 {
		return nil, errors.New("user not found. Please check the username for typos or errors")
	} else if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error fetching data, status code: %d", response.StatusCode)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	var activities []GitHubActivity
	if err := json.Unmarshal(body, &activities); err != nil {
		return nil, fmt.Errorf("error unmarshaling json: %w", err)
	}

	return activities, nil
}

func DisplayUserActivity(username string, events []GitHubActivity) error {
	if len(events) == 0 {
		return errors.New("no recent activity found for this user")
	}

	fmt.Println(style.Title.Render(fmt.Sprintf("%s's Recent Activity", username)))
	fmt.Println()

	for _, event := range events {
		var tEvent string

		switch event.Type {
		case "PushEvent":
			tEvent = fmt.Sprintf("Pushed %d commit(s) to %s", event.Payload.Size, event.Repo.Name)
		case "PullRequestEvent":
			tEvent = fmt.Sprintf("Opened a pull request in %s", event.Repo.Name)
		case "WatchEvent":
			tEvent = fmt.Sprintf("Starred %s", event.Repo.Name)
		case "ForkEvent":
			tEvent = fmt.Sprintf("Forked %s", event.Repo.Name)
		case "CreateEvent":
			tEvent = fmt.Sprintf("Created a repositoty %s", event.Repo.Name)
		case "ReleaseEvent":
			tEvent = fmt.Sprintf("Published a release in %s", event.Repo.Name)
		case "IssuesEvent":
			tEvent = fmt.Sprintf("Opened an issue in %s", event.Repo.Name)
		default:
			tEvent = fmt.Sprintf("Performed a %s action in %s", event.Type, event.Repo.Name)
		}

		parsedTime, err := time.Parse(time.RFC3339, event.CreatedAt)
		var timeString string
		if err != nil {
			timeString = "(invalid time)"
		} else {
			timeString = parsedTime.Local().Format("on 2006-01-02 at 15:04")
		}

		fmt.Println(style.Event.Render(fmt.Sprintf("%s %s", tEvent, timeString)))
	}
	return nil
}
