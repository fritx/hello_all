package cmd

import (
	"fmt"
	"todo-hub/lib"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(undoCmd)
}

var undoCmd = &cobra.Command{
	Use:     "undo",
	Short:   fmt.Sprintf("Mark a todo as '%s'", lib.StatusTitle[lib.StatusTodo]),
	Aliases: []string{"o", "todo", "uncheck"},
	Run: func(cmd *cobra.Command, args []string) {
		lib.StatusPromptLoop(lib.StatusTodo)
	},
}
