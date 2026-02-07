# Task CLI

A simple, robust command-line task manager built in Go.

## Features

- **Add Tasks**: Quickly add new to-dos.
- **List Tasks**: View all tasks with their ID, status, and creation date.
- **Complete Tasks**: Mark tasks as done.
- **Delete Tasks**: Remove tasks permanently.
- **Persistence**: Tasks are saved automatically to a local `tasks.json` file.

## Getting Started

### Prerequisites

- [Go](https://go.dev/dl/) (version 1.16 or higher recommended)

### Installation

1. Clone the repository or download the source code.
2. Navigate to the project directory.
3. Build the executable:

```bash
go build -o task-cli
```

## Usage

### Add a Task
Use the `add` command followed by the task description. If the description contains spaces, wrap it in quotes.

```bash
./task-cli add "Buy groceries"
```

### List Tasks
View all current tasks.

```bash
./task-cli list
```

**Example Output:**
```text
ID   Status   Created      Description
---------------------------------------------------------
1    [ ]      2023-10-27   Buy groceries
```

### Complete a Task
Mark a task as completed using its ID (found in the `list` output).

```bash
./task-cli complete 1
```

### Delete a Task
Remove a task using its ID.

```bash
./task-cli delete 1
```

## Project Structure

- `main.go`: Entry point and CLI command handling.
- `task.go`: Task struct definition and storage logic.
- `tasks.json`: Local data store (created automatically).
