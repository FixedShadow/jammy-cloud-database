package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start mysql-monitor-agent",
	Long:  "Start mysql-monitor-agent",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
