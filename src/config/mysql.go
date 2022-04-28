package config

import (
	"database/sql"
	"embed"
	"fmt"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/pressly/goose/v3"
)

const connMaxLifetime = time.Minute * 3
const maxOpenConnections = 20
const maxIdleConnections = 20

type MySQL struct {
	DB *sql.DB
}

func config(cfg *Config) *mysql.Config {
	addr := fmt.Sprintf(
		"%s:%s",
		cfg.Getenv("MYSQL_HOST", ""),
		cfg.Getenv("MYSQL_PORT", "3306"),
	)

	return &mysql.Config{
		User:                 cfg.Getenv("MYSQL_USER", ""),
		Passwd:               cfg.Getenv("MYSQL_PASSWORD", ""),
		Net:                  "tcp",
		Addr:                 addr,
		DBName:               cfg.Getenv("MYSQL_DBNAME", ""),
		ParseTime:            true,
		AllowNativePasswords: true,
	}
}

func NewMySQLDatabase(cfg *Config) *MySQL {
	db, err := sql.Open("mysql", config(cfg).FormatDSN())
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(connMaxLifetime)
	db.SetMaxOpenConns(maxOpenConnections)
	db.SetMaxIdleConns(maxIdleConnections)

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return &MySQL{DB: db}
}

func (db *MySQL) RunMigration(migrationFs embed.FS) {
	goose.SetBaseFS(migrationFs)

	if err := goose.SetDialect("mysql"); err != nil {
		panic(err)
	}

	if err := goose.Up(db.DB, "migrations"); err != nil {
		panic(err)
	}
}
