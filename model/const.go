package model

import (
	"errors"
)

var (
	EveryHour     string = "@hourly"   // run once ahour
	EveryDay             = "@daily"    // run once a day
	EveryWeek            = "@weekly"   // run once a week
	EveryMidnight        = "@midnight" // run once a day ( midnight )
	CustomTime    string = "@every "   // custom, example 1h30m (every hour thirty), or per 10s (every 10 second)

	ErrorConnectDatabase  = errors.New("error while connected the database")
	ErrorCreateBackupFile = errors.New("error while create backup file")
)

const (
	Mysql    string = "mysql"
	Postgres        = "postgres"
)


