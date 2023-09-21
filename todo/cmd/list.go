package cmd

import (
	"log"
	"os"
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
		c := exec.Command("prisma studio")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		if err := c.Run(); err != nil {
			// exec: executable file not found in $PATH
			// https://stackoverflow.com/questions/44786643/exec-executable-file-not-found-in-path
			// Note: this is Unix-only
			c := exec.Command("bash", "-c", "prisma studio")
			c.Stdout = os.Stdout
			c.Stderr = os.Stderr

			if err := c.Run(); err != nil {
				log.Println("Consider running: npm i -g prisma")
				log.Fatal(err)
			}
		}
	},
}
