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

	_ "github.com/lib/pq"
)

type postgresParam struct {
	*ExecutorParam
}

func NewPostgresExecutor(opts *ExecutorParam) IExecutor {
	return &postgresParam{opts}
}

func (p *postgresParam) BackupDatabase() {
	mysqlBackupFileName := fmt.Sprintf("%s/%s_%s_%s.sql", p.PathBackup, p.Database, p.DatabaseName, helpers.DateFormatter())

	pgDumpParam := []string{
		fmt.Sprintf("--username=%s", p.Username),
		fmt.Sprintf("--port=%s", p.Port),
		fmt.Sprintf("--host=%s", p.Host),
		fmt.Sprintf("--dbname=%s", p.DatabaseName),
	}

	cmd := exec.Command("pg_dump", pgDumpParam...)
	cmd.Env = append(os.Environ(), fmt.Sprintf("PGPASSWORD=%s", p.Password))

	backupFile, err := os.Create(mysqlBackupFileName)
	if err != nil {
		log.Fatal(model.ErrorCreateBackupFile)
	}
	defer backupFile.Close()

	cmd.Stdout = backupFile

	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	log.Printf("Backup success, saved at %v\n", mysqlBackupFileName)
}

func (p *postgresParam) PingDatabase() error {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.Username, p.Password, p.DatabaseName)

	db, err := sql.Open("postgres", dsn)
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
