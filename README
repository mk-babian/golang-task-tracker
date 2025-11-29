# Task Tracker

A simple command-line task management application written in Go.

## Current Status

**In Development** - Core functionality partially implemented.

### Implemented Features
- Basic CLI interface with command loop
- Task structure definition (ID, Title)
- Single task creation with JSON export
- Help menu framework

### Pending Features
- Complete task listing
- Task completion marking
- Task deletion
- Proper JSON array handling (currently overwrites)
- Persistent data management

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
0  - Display help
1  - Add task
2  - [Not implemented]
-1 - Exit
```

### Current Limitations

**Critical Issues:**
1. **Data Persistence**: Each `addTask()` call overwrites `out.json` instead of appending
2. **Hard-coded Values**: Task ID and title are static
3. **No Task List**: Cannot view existing tasks
4. **No Deletion**: Cannot remove tasks
5. **No Completion Tracking**: No way to mark tasks as done

## Technical Details

### Data Structure

```go
type Task struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
}
```

**Observation**: The structure lacks a `Completed` field for tracking task status.

### File Output

Tasks are saved to `out.json` in the current directory with indented JSON formatting.

## Roadmap

See `TODO` file for detailed development tasks.

**Priority Items:**
1. Implement task array management (read/append/write)
2. Add user input for task creation
3. Implement task listing functionality
4. Add task deletion by ID
5. Add completion status tracking

## License

MIT License - See `LICENSE` file for details.

Copyright (c) 2025 Saba

## Development Notes

**Architectural Suggestions:**

1. **Centralize Data Management**: Create a `TaskManager` struct to handle all CRUD operations
2. **Use Slice Storage**: Store tasks as `[]Task` instead of single Task objects
3. **Add Validation**: Implement ID uniqueness checks and input validation
4. **Error Handling**: Replace `panic()` with graceful error messages
5. **Configuration**: Consider using a constant for the JSON filename

**Example Improvement:**
```go
type TaskManager struct {
    Tasks    []Task
    FilePath string
    NextID   int
}
```

This would provide a cleaner abstraction for task operations and state management.

---

*The foundation is functional. The path forward requires systematic completion of core features.*
