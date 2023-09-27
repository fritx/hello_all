package cmd

import (
	"fmt"
	"todo-hub/lib"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doneCmd)
}

var doneCmd = &cobra.Command{
	Use:     "done",
	Aliases: []string{"x", "check", "finish"},
	Short:   fmt.Sprintf("Mark a todo as '%s'", lib.StatusTitle[lib.StatusDone]),
	RunE: func(cmd *cobra.Command, args []string) error {
		return lib.StatusPromptLoop(lib.StatusDone)
	},
}
