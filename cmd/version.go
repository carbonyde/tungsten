package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Ensure you are using the latest version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Tungsten v0.1.0")
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
