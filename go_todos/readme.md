# Todos app

A simple todo cli app

## Running the app

```bash
$ go run cmd/todos.go
```

## How to use

Passing help to the args will show a list of commands you can use

```bash
$ go run cmd/todos.go help
```

To call a command, simple use

```bash
$ go run cmd/todos.go <command> (args)
```

### Available commands:

- add (value): add a todo item
- delete (index): display this message
- list: show all the todos
- complete (index): toggle completion of todos
- clear: remove all todos
- help: display this message
