package main

import (
	"flag"
	"log"
	"os"

	"github.com/robfig/cron/v3"

	"github.com/fanchann/dbackupdul/executor"
	"github.com/fanchann/dbackupdul/helpers"
	"github.com/fanchann/dbackupdul/model"
)

var (
	db         string
	username   string
	password   string
	host       string
	port       string
	dbName     string
	pathBackup string
	schedule   string
)

func init() {
	flag.StringVar(&db, "db", os.Getenv("DB"), "Specify the database type to connect to. Available options: mysql, postgres")

	flag.StringVar(&username, "username", os.Getenv("DB_USERNAME"), "Username for database authentication")

	flag.StringVar(&password, "password", os.Getenv("DB_PASSWORD"), "Password for database authentication")

	flag.StringVar(&host, "host", os.Getenv("DB_HOST"), "Database host address. Example: 127.0.0.1")

	flag.StringVar(&port, "port", os.Getenv("DB_PORT"), "Database port number")

	flag.StringVar(&dbName, "dbname", os.Getenv("DB_NAME"), "Name of the database to backup")

	flag.StringVar(&pathBackup, "path_backup", os.Getenv("PATH_BACKUP"), "Path to save the backup files")

	flag.StringVar(&schedule, "schedule", os.Getenv("SCHEDULE"), "Backup scheduler")

	flag.Parse()

}

func main() {
	c := cron.New()

	opts := &executor.ExecutorParam{
		Database:     db,
		Username:     username,
		Password:     password,
		Host:         host,
		Port:         port,
		DatabaseName: dbName,
		PathBackup:   pathBackup,
	}

	switch db {
	case model.Mysql:
		mysql := executor.NewMysqlExecutor(opts)

		errPingDatabase := mysql.PingDatabase()
		if errPingDatabase != nil {
			log.Fatalf("%v \n", errPingDatabase)

		}

		c.AddFunc(helpers.Scheduler(schedule), func() {
			mysql.BackupDatabase()
		})

	case model.Postgres:

		postgres := executor.NewPostgresExecutor(opts)

		errPingDatabase := postgres.PingDatabase()
		if errPingDatabase != nil {
			log.Fatalf("%v \n", errPingDatabase)
		}

		c.AddFunc(helpers.Scheduler(schedule), func() {
			postgres.BackupDatabase()
		})

	default:
		log.Printf("environment not found!\nuse -h for help")
		os.Exit(1)
	}

	c.Run()
}
