package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"todo-hub/db"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(DoneCmd)
}

var DoneCmd = &cobra.Command{
	Use:     "done",
	Short:   "Mark a todo as done",
	Aliases: []string{"x", "check", "finish"},
	Run: func(cmd *cobra.Command, args []string) {
		TodoDone()
	},
}

func TodoDone() {
	fmt.Println("Loop: Marking a todo as done...")
	fmt.Println("Press CTRL-C to exit.")

	// connect to db
	client, disconnect := db.GetClient()
	defer disconnect()

	for {
		eachLoop(client)
	}
}

func eachLoop(client *db.PrismaClient) {
	// the questions to ask
	var qs = []*survey.Question{
		{
			Name:     "id",
			Prompt:   &survey.Input{Message: "* Id:"},
			Validate: survey.Required,
		},
	}
	// the answers will be written to this struct
	answers := struct {
		Id string
	}{}

	// perform the questions
	if err := survey.Ask(qs, &answers); err != nil {
		panic(err)
	}

	ctx := context.Background()

	found, err := client.Todo.FindUnique(
		db.Todo.ID.Equals(answers.Id),
	).Exec(ctx)
	if errors.Is(err, db.ErrNotFound) {
		fmt.Printf("Todo '%s' not found.\n", answers.Id)
		return
	} else if err != nil {
		panic(err)
	} else if found.Status == db.STATUS_DONE {
		fmt.Printf("Todo '%s' has been done before.\n", answers.Id)
		return
	}

	updated, err := client.Todo.FindUnique(
		db.Todo.ID.Equals(answers.Id),
	).Update(
		db.Todo.Status.Set(db.STATUS_DONE),
		db.Todo.LastDoneAt.Set(time.Now()),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}

	// output the json
	result, err := json.MarshalIndent(updated, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Todo marked as done: %s\n", result)
}
