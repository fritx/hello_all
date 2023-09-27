package lib

import (
	"encoding/json"
	"fmt"
	"todo-hub/db"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
)

func StatusPromptLoop(status string) error {
	fmt.Printf("Loop: Marking a todo as '%s'...\n", StatusTitle[status])
	fmt.Println("Press CTRL-C to exit.")

	// connect to db
	client, disconnect, err := db.GetClient()
	if err != nil {
		return err
	}
	defer disconnect()

	for {
		// infinite loop without error handling
		err := statusPromptEach(client, status)

		if err == terminal.InterruptErr {
			return err // quit loop
		}
		if err == ErrSkipRound {
			continue // skip round
		}
		continue
	}
}

func statusPromptEach(client *db.PrismaClient, status string) error {
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
		return err
	}
	// update status
	updated, err := UpdateStatus(client, answers.Id, status)
	if err != nil {
		return err
	}

	// output the json
	result, err := json.MarshalIndent(updated, "", "  ")
	if err != nil {
		return err
	}
	fmt.Printf("Todo marked as '%s': %s\n", StatusTitle[status], result)
	return nil
}
