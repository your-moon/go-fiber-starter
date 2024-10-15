package cmd

import (
	"github.com/spf13/cobra"
)

var Prod bool

var serveCmd = &cobra.Command{
	Use: "serve",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().BoolVarP(&Prod, "prod", "p", false, "Run in production mode")
}
