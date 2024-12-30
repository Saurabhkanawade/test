package main

import (
	"flag"
	"fmt"
	"github.com/saurabhkanawade/todocli/service"
	"os"
)

var (
	binName      = "todo"
	todoFileName = ".todo.json"
)

func main() {
	defer service.HandlePanic()

	task := flag.String("task", "", "task to be included in the todo list")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")

	flag.Parse()

	l := &service.List{}

	if err := l.Get(todoFileName); err != nil {
		_, err := fmt.Fprintf(os.Stderr, err.Error())
		if err != nil {
			return
		}
		os.Exit(1)
	}

	switch {
	case *list:
		for _, item := range *l {
			if !item.Done {
				fmt.Println(item.Task)
			}
		}
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			_, err := fmt.Fprintln(os.Stderr, err)
			if err != nil {
				return
			}
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			_, err := fmt.Fprintln(os.Stderr, err)
			if err != nil {
				return
			}
			os.Exit(1)
		}
	case *task != "":
		l.Add(*task)
		if err := l.Save(todoFileName); err != nil {
			_, err := fmt.Fprintln(os.Stderr, err)
			if err != nil {
				return
			}
			os.Exit(1)
		}
	default:
		_, err := fmt.Fprintln(os.Stderr, "Invalid Option")
		if err != nil {
			return
		}
		os.Exit(1)
	}
}
