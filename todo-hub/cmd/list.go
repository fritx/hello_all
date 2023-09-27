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
	RunE: func(cmd *cobra.Command, args []string) error {
		c := exec.Command("prisma", "studio")
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		if err := c.Run(); err != nil {
			fmt.Println("Run `npm i -g prisma` first.")
			return err
		}
		return nil
	},
}
