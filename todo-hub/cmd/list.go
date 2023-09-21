package cmd

import (
	"log"
	"os/exec"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Launch the db admin",
	Run: func(cmd *cobra.Command, args []string) {
		shell := exec.Command("prisma studio")

		if err := shell.Run(); err != nil {
			// exec: executable file not found in $PATH
			// https://stackoverflow.com/questions/44786643/exec-executable-file-not-found-in-path
			// Note: this is Unix-only
			shell := exec.Command("bash", "-c", "prisma studio")

			if err := shell.Run(); err != nil {
				log.Println("Consider running: npm i -g prisma")
				log.Fatal(err)
			}
		}
	},
}
