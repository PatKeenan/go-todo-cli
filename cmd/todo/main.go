package main 

import (
    "flag"
    "fmt"
    "os"

    todo "github.com/Patkeenan/go-todo-cli"
)


var todoFileName = "../../todoList.json"

func main() {
 task := flag.String("task", "", "Task to be included in list")
 list := flag.Bool("list", false, "List all tasks")
 complete := flag.Int("complete", 0, "Item to be completed")

 flag.Parse()

 l:= &todo.List{}

 if err := l.Get(todoFileName); err != nil {
     fmt.Fprintln(os.Stderr, err)

     os.Exit(1)
 }
     switch {
    case *list:
        // list current to do items
        for _, item := range *l {
            if !item.Done {
                fmt.Println(item.Task)
            }
        }
    // to verify if complete flag is set with value more than 0 (default)
    case *complete > 0:
        if err := l.Complete(*complete); err != nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(1)
        }
        // save the new list
        if err := l.Save(todoFileName); err != nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(1)
        }
    // verify if task flag is set with different than empty string
    case *task != "":
        l.Add(*task)
        if err := l.Save(todoFileName); err != nil {
            fmt.Fprintln(os.Stderr, err)
            os.Exit(1)
        }
    default:
        // print an error msg
        fmt.Fprintln(os.Stderr, "Invalid option")
        os.Exit(1)
    }

}
