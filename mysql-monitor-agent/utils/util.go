package utils

import (
	"bytes"
	"fmt"
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/logs"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"os"
	"os/exec"
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

func handleErr(stdout, stderr bytes.Buffer, err error) (string, error) {
	errMsg := ""
	if len(stderr.String()) != 0 {
		errMsg = fmt.Sprintf("stderr: %s", stderr.String())
	}
	if len(stdout.String()) != 0 {
		if len(errMsg) != 0 {
			errMsg = fmt.Sprintf("%s; stdout: %s", errMsg, stdout.String())
		} else {
			errMsg = fmt.Sprintf("stdout: %s", stdout.String())
		}
	}
	return errMsg, err
}

func ExecWithTimeOut(cmdStr string, timeout time.Duration) (string, error) {
	cmd := exec.Command("bash", "-c", cmdStr)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Start(); err != nil {
		return "", err
	}
	done := make(chan error, 1)
	go func() {
		done <- cmd.Wait()
	}()
	after := time.After(timeout)
	select {
	case <-after:
		_ = cmd.Process.Kill()
		return "", errors.New("cmd exec timeout")
	case err := <-done:
		if err != nil {
			return handleErr(stdout, stderr, err)
		}
	}
	return stdout.String(), nil
}

func Exec(cmdStr string) (string, error) {
	return ExecWithTimeOut(cmdStr, 600*time.Second)
}

func GetKernelPid() string {
	bys, err := os.ReadFile(KernelPidFile)
	if err != nil {
		logs.GetLogger().Error("read pid error", zap.String("pid_file", KernelPidFile))
		return ""
	}
	return string(bys)
}

func ExecCmd(cmdStr string) error {
	cmd := exec.Command("bash", "-c", cmdStr)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("error : %v, output: %s", err, output)
	}
	return nil
}
