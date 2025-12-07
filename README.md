# Task Tracker

A minimal CLI task manager that stores tasks as JSON. No database, no server, no complexity—just tasks in a file you can version control.

## Why This Exists

Most task managers are bloated. This one does exactly what you need: add tasks, delete them, see what's left. Everything lives in `tasks.json` in your project directory. You can commit it, sync it, or ignore it—it's just a file.

## What It Does

- Creates and manages a local `tasks.json` file
- Adds tasks with priorities (low/med/high)
- Lists all tasks with their IDs and priorities
- Deletes tasks by ID and automatically renumbers remaining ones
- Asks for confirmation before overwriting existing task files

## Installation

```bash
git clone https://github.com/mk-babian/golang-task-tracker.git
cd golang-task-tracker/src
go run main.go
```

## Usage Example

```
Welcome to task tracker!

Provide a command ("help" for help and "exit" to quit) : create
Provide a command ("help" for help and "exit" to quit) : add
Provide a task name: Fix authentication bug
Provide task priority (low/med/high): high

Provide a command ("help" for help and "exit" to quit) : add
Provide a task name: Update documentation
Provide task priority (low/med/high): med

Provide a command ("help" for help and "exit" to quit) : list
		--Tasks--

ID: 1 | Fix authentication bug [HIGH]
ID: 2 | Update documentation [MED]

Provide a command ("help" for help and "exit" to quit) : del
Enter task ID to delete: 1
Task deleted and IDs renumbered.

Provide a command ("help" for help and "exit" to quit) : list
		--Tasks--

ID: 1 | Update documentation [MED]
```

## Available Commands

- `help` - Show command list
- `create` - Initialize or overwrite tasks.json
- `add` - Add a new task
- `list` - Display all tasks
- `del` - Delete a task by ID
- `exit` - Quit

## Honest Limitations

**This is basic by design, but here's what it doesn't do:**

- No task completion status (can't mark tasks as "done")
- No task editing (must delete and re-add)
- No due dates or timestamps
- No categories or tags
- No sorting or filtering beyond what's shown
- No multi-user support or sync
- Tasks renumber after deletion (ID 5 becomes ID 4 if you delete task 3)

If you need these features, you probably want a different tool. If you just need to track what needs doing, this works fine.

## Project Structure

```
.
├── LICENSE
├── README.md
└── src/
    └── main.go
```

## Contributing

1. Fork the repository
2. Create a branch: `git checkout -b feature-name`
3. Make your changes
4. Test them: `go run main.go`
5. Commit: `git commit -m "Add feature-name"`
6. Push: `git push origin feature-name`
7. Open a pull request

Keep changes focused. If you're adding a feature, explain why it's useful. If you're fixing a bug, describe what broke.

## License

MIT License - See `LICENSE` file.

Copyright (c) 2025 Saba
