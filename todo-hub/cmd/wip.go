package cmd

import (
	"fmt"
	"todo-hub/lib"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(wipCmd)
}

var wipCmd = &cobra.Command{
	Use:     "wip",
	Short:   fmt.Sprintf("Mark a todo as '%s'", lib.StatusTitle[lib.StatusWip]),
	Aliases: []string{"w", "doing"},
	Run: func(cmd *cobra.Command, args []string) {
		lib.StatusPromptLoop(lib.StatusWip)
	},
}
