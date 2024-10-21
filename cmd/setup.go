package cmd

import (
	"github.com/spf13/cobra"
)

var setupCmd = &cobra.Command{
	Use: "setup",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
	// serveCmd.Flags().BoolVarP(&Prod, "prod", "p", false, "Run in production mode")
}
