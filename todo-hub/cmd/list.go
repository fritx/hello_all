package cmd

import (
	"fmt"
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
		c := exec.Command("prisma", "studio")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		if err := c.Run(); err != nil {
			fmt.Println("Run `npm i -g prisma` first.")
			panic(err)
		}
	},
}
