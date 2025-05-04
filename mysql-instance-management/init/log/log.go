package log

import (
	"fmt"
	"github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/conf"
	"github.com/FixedShadow/jammy-cloud-database/mysql-instance-management/global"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

func Init() {
	l := logrus.New()
	setOutput(l, global.CONF.LogConfig)
	global.LOG = l
	global.LOG.Info("init logger successfully.")
}

func setOutput(logger *logrus.Logger, ocnfig conf.LogConfig) {
	//TODO add log file path.
	fileAndStdoutWriter := io.MultiWriter(os.Stdout)
	logger.SetOutput(fileAndStdoutWriter)
	level, err := logrus.ParseLevel("debug")
	if err != nil {
		panic(err)
	}
	logger.SetLevel(level)
	logger.SetFormatter(new(CustomizedFormatter))
}

type CustomizedFormatter struct {
}

func (c *CustomizedFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	detailInfo := ""
	if entry.Caller != nil {
		function := strings.ReplaceAll(entry.Caller.Function, "github.com/FixedShadow/jammy-cloud-database/mysql-instance-management", "")
		detailInfo = fmt.Sprintf("(%s: %d)", function, entry.Caller.Line)
	}
	if len(entry.Data) == 0 {
		msg := fmt.Sprintf("[%s] [%s] %s %s \n", time.Now().Format(TimeFormat), strings.ToUpper(entry.Level.String()), entry.Message, detailInfo)
		return []byte(msg), nil
	}
	msg := fmt.Sprintf("[%s] [%s] %s %s {%v} \n", time.Now().Format(TimeFormat), strings.ToUpper(entry.Level.String()), entry.Message, detailInfo, entry.Data)
	return []byte(msg), nil
}
