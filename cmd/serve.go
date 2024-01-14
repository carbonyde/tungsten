package cmd

import (
	"github.com/labstack/echo/v4"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run a local server for your static generated page",
	Run: func(cmd *cobra.Command, args []string) {
		e := echo.New()

		e.Static("/", "dist")

		e.Logger.Fatal(e.Start("0.0.0.0:8080"))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
}
