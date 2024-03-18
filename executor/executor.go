package executor


type ExecutorParam struct {
	Database     string
	Host         string
	Port         string
	Username     string
	Password     string
	DatabaseName string
	PathBackup   string
}

type IExecutor interface {
	BackupDatabase()
	PingDatabase() error
}
