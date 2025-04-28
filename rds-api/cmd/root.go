package cmd

import (
	"github.com/FixedShadow/jammy-cloud-database/rds-api/server"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rds-api",
	Short: "rds-api",
	Run: func(cmd *cobra.Command, args []string) {
		server.Start()
		return
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
