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
	Run: func(cmd *cobra.Command, args []string) {
		// alias to `todo add`
		// AddCmd.Execute() // fatal error: stack overflow
		TodoAdd()
	},
}

func Execute() {
	defer func() {
		if r := recover(); r != nil {
			if r == terminal.InterruptErr {
				fmt.Println("Canceled.")
				return
			}
			// panic(r)
			log.Fatal(r) // log without stack
		}
	}()

	if err := rootCmd.Execute(); err != nil {
		// panic(err)
		log.Fatal(err) // log without stack
	}
}
