package cmd

import (
	"github.com/spf13/cobra"
	"github.com/your-moon/go-fiber-starter/gcli"
)

var setupCmd = &cobra.Command{
	Use: "setup",
	RunE: func(cmd *cobra.Command, args []string) error {
		gcli.Run()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(setupCmd)
	// serveCmd.Flags().BoolVarP(&Prod, "prod", "p", false, "Run in production mode")
}
