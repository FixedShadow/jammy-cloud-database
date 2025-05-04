package utils

import (
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/logs"
	"go.uber.org/zap"
	"os"
	"path/filepath"
	"time"
)

var workingPath string

func GetCurrTSInNano() int64 {
	return time.Now().UnixNano()
}

func GetWorkingPath() string {
	var err error
	if workingPath == "" {
		workingPath, err = filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logs.GetLogger().Error("Get working path path failed", zap.Error(err))
			return ""
		}
	}
	return workingPath
}

func GetMainThreadPidFilePath() string {
	return filepath.Join(GetWorkingPath(), AgentPIDFileName)
}
