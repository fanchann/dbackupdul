# dbackupdul
dbackupdul (database backup scheduler) is a powerful tool that automates database backups, providing precise scheduling and comprehensive support for MySQL and PostgreSQL. Its intuitive interface simplifies setup, enhancing data reliability and minimizing the risk of data loss, making it indispensable for administrators and developers.

## Database Supported
- mysql
- postgresql

## Installation
go 1.2x:
```bash
go install github.com/fanchann/dbackupdul@latest
```

docker:
```bash
docker pull fannchannn/dbackupdul:1.0
```

## Docker
check it out [dbackupdul](https://hub.docker.com/r/fannchannn/dbackupdul)
## Environment
`-db` : specify the database type to connect to. Available options: mysql, postgres\
`-username` : username for database authentication\
`-password`: password for database authentication\
`-host` : host address of the database, for example `127.0.0.1`\
`-port` : port number of the database \
`-dbname` : name of the database to be backup\
`-path_backup` : path to store the backup files\
`-schedule` : backup scheduling

## Scheduling format
`hourly` : backup once a hour\
`daily` : backup once a day\
`midnight` : backup once at midnight\
`weekly` : backup once a week\
Or
you can custom the backup time:\
example :\
`1h30m` : backup every a hour thirty\
`50m` : backup every fifty minutes\
`10s` : backup every ten second
## How to run 
```bash
dbackupdul -db <mysql | postgres> -username <username> -password <password> -host <database_host> -port <database_port> -dbname <database_name> -path_backup <backup_folder> -schedule <schedule_time>
```
