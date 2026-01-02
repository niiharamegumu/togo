# TOGO

TOGO is a simple, terminal-based TODO list CLI tool written in Go. Focused on simplicity, it allows you to manage tasks efficiently with rich interactive prompts.

## Features

- **Interactive UI**: Powered by `charmbracelet/huh`, providing a modern terminal experience for adding, updating, and deleting tasks.
- **Physical Deletion**: Simplified task management workflow—complete tasks by deleting them permanently.
- **Detailed Table View**: View your tasks in a clean, formatted table with fixed-width columns for readability.
- **Smart Sorting**: 
  - Sort by Title, Priority, Dates, etc.
  - Tasks without a due date are automatically pushed to the bottom of the list when sorting by Due Date.

## Installation

### Using Go

You can install TOGO directly using the Go toolchain:

```shell
go install github.com/niiharamegumu/togo@latest
```

### Using Makefile (for development)

```shell
make install
```

## Configuration

By default, TOGO stores its database at `~/.togo/tasks.db`.

If you wish to change the storage location, set the `TOGO_PROJECT_ROOT_PATH` environment variable:

```shell
export TOGO_PROJECT_ROOT_PATH=/path/to/your/preferred/directory
```

## Usage

### Commands

#### Add a Task (`add`, `a`)

```shell
togo add
```
Opens an interactive form to set the title, priority (0-100), and due date (YYYY-MM-DD).

#### List Tasks (`list`, `l`)

```shell
togo list [flags]
```

**Sorting Options:**
- `--sort`, `-s`: `id(i)`, `title(t)`, `priority(p)`, `created_at(c)`, `updated_at(u)`, `due_date(d)`
- `--sort-direction`, `-d`: `asc`, `desc`

Example:
```shell
togo list --sort priority --sort-direction desc
```

#### Update a Task (`update`, `u`)

```shell
togo update
```
Select a task from the list and update its details (Title, Priority, Due Date) via an interactive form.

#### Delete a Task (`del`, `de`)

```shell
togo del
```
Select one or more tasks using a multi-select list to permanently remove them.

## Development

- `make build`: Build the `togo` binary.
- `make test`: Run project tests.
- `make run`: Run the application directly.
- `make clean`: Remove build artifacts.

## License

MIT License.

