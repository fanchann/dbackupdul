package executor

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/fanchann/dbackupdul/helpers"
	"github.com/fanchann/dbackupdul/model"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlParam struct {
	*ExecutorParam
}

func NewMysqlExecutor(opts *ExecutorParam) IExecutor {
	return &mysqlParam{opts}
}

func (e *mysqlParam) PingDatabase() error {
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/", e.Username, e.Password, e.Host, e.Port)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return model.ErrorConnectDatabase
	}
	defer db.Close()

	db.SetConnMaxLifetime(5 * time.Second)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil {
		return model.ErrorConnectDatabase
	}

	return nil
}

func (e *mysqlParam) BackupDatabase() {
	mysqlBackupFileName := fmt.Sprintf("%s/%s_%s_%s.sql", e.PathBackup, e.Database, e.DatabaseName, helpers.DateFormatter())

	mysqlDumpParam := []string{
		fmt.Sprintf("-u%s", e.Username),
		fmt.Sprintf("-p%s", e.Password),
		fmt.Sprintf("-h%s", e.Host),
		fmt.Sprintf("-P%s", e.Port),
		e.DatabaseName,
	}

	cmd := exec.Command("mysqldump", mysqlDumpParam...)

	backupFile, err := os.Create(mysqlBackupFileName)
	if err != nil {
		log.Fatal(model.ErrorCreateBackupFile)
	}
	defer backupFile.Close()

	cmd.Stdout = backupFile

	if err := cmd.Run(); err != nil {
		// trying mariadb-dump
		mariaDBDump(e, backupFile)
	}

	log.Printf("Backup success, saved at %v\n", mysqlBackupFileName)
}

func mariaDBDump(e *mysqlParam, fileData *os.File) error {
	mariadbDumpParam := []string{
		fmt.Sprintf("-u%s", e.Username),
		fmt.Sprintf("-p%s", e.Password),
		fmt.Sprintf("-h%s", e.Host),
		fmt.Sprintf("-P%s", e.Port),
		e.DatabaseName,
	}

	mariadbCmd := exec.Command("mariadb-dump", mariadbDumpParam...)

	mariadbCmd.Stdout = fileData

	if err := mariadbCmd.Run(); err != nil {
		return err
	}
	return nil
}
