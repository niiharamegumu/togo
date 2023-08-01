package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var helpCmd = &cobra.Command{
	Use:   "help",
	Short: "Display available commands",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Available commands:")
		for _, c := range cmd.Root().Commands() {
			if c.Name() != "help" {
				fmt.Printf("- %s: %s\n", c.Name(), c.Short)
			}
		}
	},
}
