package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"todo-hub/db"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(AddCmd)
}

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a todo",
	Run: func(cmd *cobra.Command, args []string) {
		TodoAdd()
	},
}

// the questions to ask
var qs = []*survey.Question{
	{
		Name:     "title",
		Prompt:   &survey.Input{Message: "* Title:"},
		Validate: survey.Required,
		// Transform: survey.Title,
	},
	{
		Name:   "desc",
		Prompt: &survey.Input{Message: "Description:"},
	},
}

func TodoAdd() {
	fmt.Println("Adding a todo...")
	// the answers will be written to this struct
	answers := struct {
		Title string // survey will match the question and field names
		Desc  string // if the types don't match, survey will convert it
	}{}
	// perform the questions
	if err := survey.Ask(qs, &answers); err != nil {
		if err == terminal.InterruptErr {
			fmt.Println("Canceled.")
			return // skip
		}
		panic(err)
	}

	// connect to db
	client := db.NewClient()
	if err := client.Prisma.Connect(); err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Prisma.Disconnect(); err != nil {
			log.Fatal(err) // log without stack
		}
	}()

	// insert a record
	ctx := context.Background()
	created, err := client.Todo.CreateOne(
		db.Todo.Title.Set(answers.Title),
		db.Todo.Desc.Set(answers.Desc),
	).Exec(ctx)
	if err != nil {
		panic(err)
	}

	// output the json
	result, err := json.MarshalIndent(created, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Todo added: %s\n", result)
}