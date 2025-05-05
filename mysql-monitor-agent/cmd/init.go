package cmd

import (
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/deploy"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init mysql instance",
	Run: func(cmd *cobra.Command, args []string) {
		deploy.DownloadKernel()
		deploy.RunKernel()
	},
}
