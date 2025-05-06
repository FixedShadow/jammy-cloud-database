package deploy

import (
	"context"
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/logs"
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/monitor/config"
	"github.com/FixedShadow/jammy-cloud-database/mysql-monitor-agent/utils"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
)

var (
	OSSClient   *minio.Client
	StorageConf *config.MonitorConfig
)

func init() {
	StorageConf = config.GetMonitorConfig()
	var err error
	OSSClient, err = minio.New(StorageConf.Storage.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(StorageConf.Storage.AccessKeyId, StorageConf.Storage.SecretAccessKey, ""),
		Secure: false,
	})
	if err != nil {
		logs.GetLogger().Error("init oss client error", zap.Error(err))
		panic(err)
	}
}

// DownloadKernel download the tar from oss
func DownloadKernel() {
	err := OSSClient.FGetObject(context.Background(), StorageConf.Storage.KernelBucket,
		StorageConf.Storage.KernelTarFile, StorageConf.Storage.DownloadPath, minio.GetObjectOptions{})
	if err != nil {
		logs.GetLogger().Error("cannot download the kernel packages", zap.Error(err))
		panic(err)
	}
}

// RunKernel start the mysql kernel.
func RunKernel() {
	err := utils.ExecCmdAsync("bash " + utils.InstallKernelShellPath)
	if err != nil {
		logs.GetLogger().Error("run mysql kernel error", zap.Error(err))
		panic(err)
	}
}

// TODO upgrade agent
func UpgradeAgent() {

}
