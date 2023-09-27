package cmd

import (
	"fmt"
	"log"

	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "todo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and easy todo-list CLI and admin UI in Go.
		Complete documentation is available at https://github.com/fritx/hello_all/todo-hub`,
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		// alias to `todo add`
		// AddCmd.Execute() // fatal error: stack overflow
		return add()
	},
}

func handleError(err interface{}) {
	if err == terminal.InterruptErr {
		fmt.Println("Canceled.")
		return
	}
	log.Fatal(err) // log without stack
}

func Execute() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("root recover: %v\n", r)
			handleError(r)
		}
	}()

	if err := rootCmd.Execute(); err != nil {
		// fmt.Printf("root execute: %v\n", err)
		handleError(err)
	}
}
