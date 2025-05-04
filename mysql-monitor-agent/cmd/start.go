package cmd

import (
	"fmt"
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/logs"
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/utils"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
)
import "github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/manager"

func init() {
	rootCmd.AddCommand(startCmd)
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start mysql-monitor-agent",
	Long:  "Start mysql-monitor-agent",
	Run: func(cmd *cobra.Command, args []string) {
		err := manager.SingleInstanceCheck()
		if err != nil {
			fmt.Println("Start failed.mysqlMonitorAgent was started in before.")
			logs.GetLogger().Warn("Start failed: mysqlMonitorAgent was started in before.", zap.Int("PID", os.Getpid()))
			return
		}
		//TODO register service here.
		fmt.Println("Mysql-Monitor-Agent runs successfully.")
		logs.GetLogger().Info("mysqlMonitorAgent runs successfully")
		utils.StartDaemon()
	},
}
