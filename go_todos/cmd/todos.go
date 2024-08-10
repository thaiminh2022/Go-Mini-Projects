package main

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"

	"github.com/thaiminh2022/go_todos/internal"
)

func main() {
	// Setup db
	db := &internal.Database{}
	if err := db.LoadDatabase(); err != nil {
		err = db.NewDatabase()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	// args -> path command_name value
	args := os.Args
	if len(args) <= 1 {
		fmt.Println("bad commands, call the program again with help")
		return
	}

	command := args[1]

	switch command {
	case "add":
		if len(args) < 3 {
			badArgs()
			return
		}

		value := args[2]
		todo := internal.TodoData{Value: value}
		db.AddTodo(todo)
	case "list":
		w := tabwriter.NewWriter(
			os.Stdout,
			0,
			0,
			5,
			' ',
			tabwriter.AlignRight|tabwriter.Debug)

		fmt.Fprintln(w, "Index\tTodo\tCompleted\t")
		for i, e := range db.Todos {
			fmt.Fprintf(w, "%v\t%v\t%v\t\n", i, e.Value, e.Complete)
		}
		w.Flush()
	case "delete":
		if len(args) < 3 {
			badArgs()
			return
		}

		index, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println(err)
			badArgs()
		}
		db.DeleteTodoAt(index)
		fmt.Println(db.Todos)
	case "complete":
		if len(args) < 3 {
			badArgs()
			return
		}

		index, err := strconv.Atoi(args[2])
		if err != nil {
			fmt.Println(err)
			badArgs()
		}
		db.ToggleTodoComplete(index)
	case "clear":
		db.ClearTodo()

	case "help":
		fmt.Println("Available commands:")
		fmt.Println("\nadd (value): add a todo item")
		fmt.Println("delete (index): display this message")
		fmt.Println("list: show all the todos")
		fmt.Println("complete (index): toggle completion of todos")
		fmt.Println("clear: remove all todos")
		fmt.Println("help: display this message")
	default:
		badArgs()
	}
	err := db.SaveDatabase()
	if err != nil {
		fmt.Println(err)
	}

}

func badArgs() {

	fmt.Println("bad commands, call the program again with help")

}
