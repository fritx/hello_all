package lib

import (
	"encoding/json"
	"fmt"
	"todo-hub/db"

	"github.com/AlecAivazis/survey/v2"
)

func StatusPromptLoop(status string) {
	fmt.Printf("Loop: Marking a todo as '%s'...\n", StatusTitle[status])
	fmt.Println("Press CTRL-C to exit.")

	// connect to db
	client, disconnect := db.GetClient()
	defer disconnect()

	for {
		statusPromptEach(client, status)
	}
}

func statusPromptEach(client *db.PrismaClient, status string) {
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
	// update status
	updated := UpdateStatus(client, answers.Id, status)
	if updated == nil {
		return // skip
	}

	// output the json
	result, err := json.MarshalIndent(updated, "", "  ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Todo marked as '%s': %s\n", StatusTitle[status], result)
}
