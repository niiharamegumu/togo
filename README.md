# TOGO

TOGO is a simple, terminal-based task management CLI tool written in Go. It allows you to manage your tasks efficiently with commands to add, list, update, and complete tasks directly from your shell.

## Features

- **Interactive Task Addition**: Easily add tasks via an interactive prompt.
- **Task Listing**: Filter tasks by status (Pending, Done, or All).
- **Task Management**: Update, mark as done, or delete tasks by ID.
- **Customizable Sorting**: Sort listed tasks by various attributes (Title, Priority, Due Date, etc.).

## Installation

### Using Go

You can install TOGO directly using the Go toolchain:

```shell
go install github.com/niiharamegumu/togo@latest
```

### Using Makefile (for development)

If you have the source code locally, you can use the provided Makefile:

```shell
make install
```

## Configuration

By default, TOGO stores its database at `~/.togo/tasks.db`. The directory and database file will be created automatically on the first run.

If you wish to change the storage location, you can set the `TOGO_PROJECT_ROOT_PATH` environment variable:

```shell
export TOGO_PROJECT_ROOT_PATH=/path/to/your/preferred/directory
```

## Usage

### Commands

#### Add a Task

```shell
togo add
```
Follow the interactive prompt to enter the task title.

#### List Tasks

```shell
togo list [status]
```

- `pen`: List pending tasks (default).
- `done`: List completed tasks.
- `all`: List all tasks.

Example with flags:
```shell
togo list --status all --sort priority --sort-direction desc
```

#### Update a Task

```shell
togo update [ID]
```
Updates the title of the task with the specified ID.

#### Complete a Task

```shell
togo done [ID]
```
Marks the task with the specified ID as completed.

#### Delete a Task

```shell
togo del [ID]
```
Removes the task with the specified ID.

### Development

The following `make` targets are available:

- `make build`: Build the `togo` binary.
- `make test`: Run project tests.
- `make run`: Run the application directly.
- `make clean`: Remove build artifacts.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details (if applicable).

