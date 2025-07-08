# Task Tracker CLI 📝

A simple and efficient command-line tool to manage your daily tasks, built with Go and Cobra. This project is inspired by the "Task Tracker" project from [roadmap.sh](https://roadmap.sh/projects/task-tracker).

---

## ✨ Features

- **Add:** Quickly add new tasks to your list.
- **List:** View all your tasks, or filter them by status (`todo`, `in-progress`, `done`).
- **Update:** Change the status of an existing task.
- **Delete:** Mark a task as "deleted".
- **Local Storage:** All tasks are saved locally in a `tasks.json` file.

---

## ✅ Prerequisites

Before you begin, ensure you have the following installed on your system:
- **Go:** Version 1.18 or higher.
- **Git:** For cloning the repository.

---

## 🚀 Installation & Setup

There are two ways to set up the CLI on your machine.

### Option 1: Using `go install` (Recommended)

This is the easiest way to install the command-line tool so you can run it from anywhere.

```bash
go install [github.com/FrenaldyH/backend/task_tracker@latest](https://github.com/FrenaldyH/backend/task_tracker@latest)
```
After installation, make sure your `$GOPATH/bin` directory is in your system's `PATH`.

### Option 2: From Source (via `git clone`)

Use this method if you want to explore or modify the source code.

**1. Clone the repository:**
```bash
git clone [https://github.com/FrenaldyH/backend.git](https://github.com/FrenaldyH/backend.git)
```

**2. Navigate to the project directory:**
```bash
cd backend/task_tracker
```

**3. Run or Build the application:**
You can either run the program directly or build a permanent executable.

* **To run directly (for testing):**
    ```bash
    go run main.go list
    ```
* **To build an executable file:**
    ```bash
    go build
    ```
    This will create an executable file named `task_tracker`. You can then run it like this:
    ```bash
    ./task_tracker add "My new task from source"
    ```

---

## 📖 Usage

Here are the basic commands to get you started.

### 1. Add a new task
Use the `add` command or its aliases (`a`, `tambah`) followed by your task description in quotes.

```bash
task_tracker add "Finish the Readme for the CLI project"
```

### 2. List all tasks
Use the `list` command or its alias `ls` to view all non-deleted tasks.

```bash
task_tracker list
```

**Filter tasks by status:**
You can also view tasks with a specific status.

```bash
# View tasks that are yet to be done
task_tracker list todo

# View tasks that are currently in progress
task_tracker list in-progress

# View tasks that have been completed
task_tracker list done
```

### 3. Update a task's status
Use the `update` command followed by the task ID and the `--status` flag.

```bash
# Change the status of task with ID 1 to "in-progress"
task_tracker mark-in-progress 1

# Change the status of task with ID 2 to "done"
task_tracker mark-done 2

# Change the status of task with ID 3 to "todo"
tack_tracker mark-todo 3
```

### 4. Delete a task
Use the `delete` command or its aliases (`d`, `rm`, `hapus`) followed by the task ID. This command doesn't actually remove the data from the file but changes its status to "deleted".

```bash
task_tracker delete 3
```