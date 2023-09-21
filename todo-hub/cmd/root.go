package cmd

import (
	"fmt"
	"os"

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
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
