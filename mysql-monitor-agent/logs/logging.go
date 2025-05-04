package logs

import (
	"encoding/json"
	"fmt"
	error2 "github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/error"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var (
	logger              *zap.Logger
	logLevel            string
	logDir              string
	singleFileMaxSizeMB int
	maxBackups          int
	logsKeepDay         int
	loggerSync          sync.Once
)

type Config struct {
	LogLevel            string
	LogDir              string
	SingleFileMaxSizeMB int
	MaxBackups          int
	LogsKeepDay         int
}

func initLogDir(dir string) error {
	info, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return os.MkdirAll(dir, 0755)
		}
		return err
	}
	if info.IsDir() {
		return nil
	}
	return error2.Errors.DirPathError
}

func initLogger() {
	pwd := GetCurrentDirectory()
	file, err := os.Open(pwd + "/conf_logs.json")
	if err != nil {
		fmt.Fprint(os.Stderr, "Open conf_logs.json configuration file", err)
		return
	}
	decoder := json.NewDecoder(file)
	logsConfig := Config{}
	err = decoder.Decode(&logsConfig)
	if err != nil {
		fmt.Fprint(os.Stderr, "Parse conf_logs.json failed", err)
		return
	}
	logLevel = logsConfig.LogLevel
	logDir = logsConfig.LogDir
	singleFileMaxSizeMB = logsConfig.SingleFileMaxSizeMB
	maxBackups = logsConfig.MaxBackups
	logsKeepDay = logsConfig.LogsKeepDay

	var level zapcore.Level
	var encoder zapcore.Encoder
	var appName string
	if !filepath.IsAbs(logDir) {
		logDir, err = filepath.Abs(logDir)
		if err != nil {
			panic(err)
		}
	}
	executableName, err := os.Executable()
	if err != nil {
		appName = "unknown"
	} else {
		appName = filepath.Base(executableName)
	}
	initLogDir(logDir)
	level.Set(logLevel)
	levelFunc := zap.LevelEnablerFunc(func(l zapcore.Level) bool {
		if l >= level {
			return true
		}
		return false
	})
	encoder = zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filepath.Join(logDir, appName+".log"),
		MaxSize:    singleFileMaxSizeMB,
		MaxAge:     logsKeepDay,
		MaxBackups: maxBackups,
	})
	c := zapcore.NewCore(encoder, w, levelFunc)
	logger = zap.New(c)
	if err != nil {
		panic(err)
	}
}

func GetLogger() *zap.Logger {
	loggerSync.Do(initLogger)
	return logger
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fmt.Fprint(os.Stderr, err)
	}
	return strings.Replace(dir, "\\", "/", -1)
}
