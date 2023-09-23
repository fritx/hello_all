package lib

import (
	"context"
	"errors"
	"fmt"
	"time"
	"todo-hub/db"
)

// todo: use enum instead
const (
	StatusTodo = ""
	StatusWip  = "wip"
	StatusDone = "done"
)

var StatusTitle = make(map[string]string)

func init() {
	StatusTitle[StatusTodo] = "todo"
	StatusTitle[StatusWip] = "wip"
	StatusTitle[StatusDone] = "done"
}

func UpdateStatus(client *db.PrismaClient, Id string, status string) (updated *db.TodoModel) {

	ctx := context.Background()

	found, err := client.Todo.FindUnique(
		db.Todo.ID.Equals(Id),
	).Exec(ctx)
	if errors.Is(err, db.ErrNotFound) {
		fmt.Printf("Todo '%s' not found.\n", Id)
		return
	} else if err != nil {
		panic(err)
	} else if found.Status == status {
		fmt.Printf("Todo '%s' has been %s before.\n", Id, status)
		return
	}

	setParams := []db.TodoSetParam{
		db.Todo.Status.Set(status),
	}
	if status == StatusDone {
		setParams = append(setParams, db.Todo.LastDoneAt.Set(time.Now()))
	}
	ud, err := client.Todo.FindUnique(
		db.Todo.ID.Equals(Id),
	).Update(
		setParams...,
	).Exec(ctx)
	if err != nil {
		panic(err)
	}
	updated = ud
	return
}
