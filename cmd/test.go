package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Test bool

var testCmd = &cobra.Command{
	Use: "test",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("test")
	},
}

func init() {
	Test = true
	rootCmd.AddCommand(testCmd)
}
