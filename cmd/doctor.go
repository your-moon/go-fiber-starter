package cmd

import (
	"github.com/spf13/cobra"
)

var doctorCmd = &cobra.Command{
	Use: "doctor",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(doctorCmd)
}
