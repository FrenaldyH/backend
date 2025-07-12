# GitHub User Activity CLI

A command-line interface tool to fetch and display GitHub user activities directly in your terminal. Built with Go and enhanced with elegant styling using Lipgloss.

This project is inspired by the [GitHub User Activity](https://roadmap.sh/projects/github-user-activity) project from roadmap.sh.

## ✨ Features

- Fetch recent GitHub user activities
- Beautiful terminal output with colors and styling
- Support for various GitHub events (push, pull requests, stars, forks, etc.)
- Error handling for network issues and invalid usernames
- Local timezone display

## 🚀 Prerequisites

Before installing and using this tool, make sure you have:

- **Go 1.19+** installed on your system
- **Git** for cloning the repository
- **Internet connection** to fetch GitHub API data

## 📦 Installation

### Option 1: Clone and Build from Source

```bash
# Clone the repository
git clone https://github.com/FrenaldyH/backend.git
cd backend/github_user_activity

# Install dependencies
go mod tidy

# Build the application
go build -o github-activity .
```

### Option 2: Direct Installation

```bash
# Install directly using go install
go install github.com/FrenaldyH/backend/github_user_activity@latest
```

## 🎯 Usage

### Basic Usage

```bash
# Fetch activity for a specific GitHub user
github-activity <username>
```

### Examples

```bash
# Get activities for user "octocat"
github-activity kamranahmedse

# Get activities for user "torvalds"
github-activity FrenaldyH

# Get your own activities
github-activity yourusername
```

### Sample Output

```
         FrenaldyH's Recent Activity
════════════════════════════════════

  │  Pushed 1 commit(s) to FrenaldyH/backend on 2025-07-09 at 21:33
  │  Performed a PublicEvent action in FrenaldyH/snippets on 2025-07-09 at 03:29
  │  Pushed 1 commit(s) to FrenaldyH/backend on 2025-07-08 at 03:53

```

## 🏗️ Project Structure

```
github_user_activity/
├── cmd/
│   └── root.go          # CLI command definitions using Cobra
├── internal/
│   └── activity.go      # Core logic for fetching and displaying activities
├── style/
│   └── style.go         # Terminal styling using Lipgloss
├── main.go              # Application entry point
├── go.mod               # Go module dependencies
├── go.sum               # Go module checksums
└── README.md            # This file
```

## 🔧 Dependencies

This project uses the following Go packages:

- **[Cobra](https://github.com/spf13/cobra)** - For building the CLI interface
- **[Lipgloss](https://github.com/charmbracelet/lipgloss)** - For beautiful terminal styling
- **Standard Library** - For HTTP requests and JSON parsing

## 🎨 Styling

The application features beautiful terminal output with:

- **Title Styling**: Bold text with double borders and custom colors
- **Event Styling**: Left-bordered entries with consistent formatting
- **Error Styling**: Red-colored error messages for better visibility
- **Color Scheme**: Carefully selected colors for optimal readability

## 🐛 Error Handling

The application handles various error scenarios:

- **Network Errors**: Connection issues or API unavailability
- **Invalid Usernames**: Returns user-friendly error for non-existent users
- **API Rate Limits**: Handles GitHub API rate limiting gracefully
- **Invalid JSON**: Manages malformed API responses
- **No Activity**: Notifies when a user has no recent public activity

## 🔍 Supported Activity Types

The CLI recognizes and formats the following GitHub event types:

| Event Type | Description |
|------------|-------------|
| `PushEvent` | Code pushes to repositories |
| `PullRequestEvent` | Pull request creation |
| `WatchEvent` | Repository starring |
| `ForkEvent` | Repository forking |
| `CreateEvent` | Repository creation |
| `ReleaseEvent` | Release publishing |
| `IssuesEvent` | Issue creation |
| *Others* | Generic activity display |



## 🙏 Acknowledgments

- [GitHub API](https://docs.github.com/en/rest) - For providing the user activity data
- [Cobra](https://github.com/spf13/cobra) - For the excellent CLI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - For beautiful terminal styling
- [roadmap.sh](https://roadmap.sh/projects/github-user-activity) - For the project inspiration

