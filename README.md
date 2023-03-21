# Go To Do CLI

### Overview

This is a just for fun project to help me learn go. It's a simple command line tool which allows a user to add, mark complete, list, and remove to dos.

### Usage

Navigate to the todo directory located in `cmd/todo`.

1. Add a task

```zsh

go run main.go -task "Example task"

```

2. List Tasks

```zsh

go run main.go -list

```

3. Mark Task as Complete

```zsh

go run main.go -complete <Task Number>

```
