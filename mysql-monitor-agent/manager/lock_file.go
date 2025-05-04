package manager

import (
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/logs"
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/utils"
	"github.com/nightlyone/lockfile"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func SingleInstanceCheck() (err error) {
	lock, err := lockfile.New(utils.GetMainThreadPidFilePath())
	if err != nil {
		return errors.Wrap(err, "Cannot init pid file")
	}
	err = lock.TryLock()
	if err != nil {
		p, _ := lock.GetOwner()
		logs.GetLogger().Debug("GetOwner successfully and pid in file ", zap.Any("pid", p.Pid))
		return errors.Wrap(err, "Cannot lock")
	}
	return nil
}
