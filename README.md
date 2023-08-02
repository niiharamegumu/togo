# TOGO

TOGO is a simple CLI tool for task management that runs on the terminal. You can use the following commands to perform tasks such as adding, listing, updating, and completing tasks.

## Installation

To install TOGO, follow these steps:

1. Make sure you have Go language installed. If it's not installed, download the installer from the Go official website and install it. After that, create a new project:

```
go mod init xxx
```

2. Use the `go get` command to install TOGO:

```shell
go get github.com/niiharamegumu/togo
```

3. Specify the project root path in the environment variable of the project root:

```
export TOGO_PROJECT_ROOT_PATH=
```

```
source ~/.zshrc
```

4. The `togo` command will now be available. Execute the command below to confirm it's working:

```
togo
```

## Command List

### Add Task

To add a new task, use the following command:

```
togo add
```

interactive mode

```
Enter the new task title : write new tittle
```

### List Tasks

To display the list of current tasks, use the following command with the option [pen | done | all]:

```
togo list ["pen" | "done" | "all" | ""]
```

### Update Task

既存のタスクのタイトルを更新するには、以下のコマンドを使用します。

```
togo update [ID]
```

interactive mode

```
Enter the new task title : write new tittle
```

### Complete Task

To mark a task as completed, use the following command:

```
togo done [ID]
```

### Delete Task

To delete a task, use the following command:

```
togo del [ID]
```
