package cmd

import (
	"github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/server"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "mysql-instance-management",
	Short: "mysql-instance-management",
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
