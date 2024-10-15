package cmd

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var Prod bool

var serveCmd = &cobra.Command{
	Use: "serve",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("CMD", args)
		if Prod && Test {
			color.RedString("Cannot run in both prod and test mode")
			return errors.New("Cannot run in both prod and test mode")
		}

		fmt.Println("serve")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().BoolVarP(&Prod, "prod", "p", false, "Run in production mode")
}
