package config

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

const connMaxLifetime = time.Minute * 3

var MySQLConfig *mysql.Config

func getMySQLConfig(cfg *Config) *mysql.Config {
	addr := fmt.Sprintf(
		"%s:%s",
		cfg.Getenv("MYSQL_HOST", ""),
		cfg.Getenv("MYSQL_PORT", "3306"),
	)

	MySQLConfig = &mysql.Config{
		User:   cfg.Getenv("MYSQL_USER", ""),
		Passwd: cfg.Getenv("MYSQL_PASSWORD", ""),
		Net:    "tcp",
		Addr:   addr,
		DBName: cfg.Getenv("MYSQL_DBNAME", ""),
	}

	return MySQLConfig
}

func NewMySQLDatabase(cfg *Config) *sql.DB {
	db, err := sql.Open("mysql", getMySQLConfig(cfg).FormatDSN())
	if err != nil {
		panic(err)
	}

	db.SetConnMaxLifetime(connMaxLifetime)

	var version string
	err2 := db.QueryRow("SELECT VERSION()").Scan(&version)

	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Println(version)
	fmt.Println("oe")

	return db
}
