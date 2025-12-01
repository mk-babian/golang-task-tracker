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
- Complete task listing
- Task completion marking
- Task deletion

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
add - Add a task to tasks.json
exit - Quit the application
```

### Current Limitations

**Critical Issues:**
1. **No Task List**: Cannot view existing tasks
2. **No Deletion**: Cannot remove tasks
3. **No Completion Tracking**: No way to mark tasks as done

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
1. Implement task listing functionality
2. Add task deletion by ID
3. Add completion status tracking

## License

MIT License - See `LICENSE` file for details.

Copyright (c) 2025 Saba

*The foundation is functional. The path forward requires systematic completion of core features.*
