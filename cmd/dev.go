package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var devCmd = &cobra.Command{
	Use:   "dev",
	Short: "Start a dev server with hot reload",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("dev called")
	},
}

func init() {
	rootCmd.AddCommand(devCmd)
}
