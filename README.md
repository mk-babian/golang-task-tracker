# Task Tracker

A simple command-line task management application written in Go.

## Current Status

**In Development** - Core functionality partially implemented.

### Implemented Features
- Basic CLI interface with command loop
- Task structure definition (ID, Title, Priority, etc...)
- Single task creation with JSON export
- Help menu framework

### Pending Features
- Task completion marking
- Task deletion
- Other optional features (coloring, personalized settings?)

## Project Structure

```
.
├── LICENSE          # MIT License
├── TODO             # Development roadmap
├── src/
│   └── main.go      # Main application code
└── README.md        # This file
```

## Installation

```bash
# Clone the repository
git clone <repository-url>

# Navigate to source directory
cd task-tracker/src

# Run the application
go run main.go
```

## Usage

The application presents a command menu:

```
help - Display help
create - Create .json file in running directory
list - List the tasks currently in the .json file 
add - Add a task to tasks.json
exit - Quit the application
```

### Current Limitations

**Critical Issues:**
1. **No Deletion**: Cannot remove tasks
2. **No Completion Tracking**: No way to mark tasks as done
3. **No additional features**: Optional features haven't yet been added

### Data Structure

```go
type Task struct {
    ID          int    `json:"id"`
    Title       string `json:"title"`
    Priority    string `json:"priority"`
}
```

### File Output

Tasks are saved to `tasks.json` in the current directory with indented JSON formatting.

## Roadmap

See `TODO` file for detailed development tasks.

**Priority Items:**
1. Build a better task listing functionality
2. Add task deletion by ID
3. Add completion status tracking

## License

MIT License - See `LICENSE` file for details.

Copyright (c) 2025 Saba
