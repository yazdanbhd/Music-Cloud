package mysqldb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Config struct {
	Username string `koanf:"user_name"`
	Password string `koanf:"password"`
	Port     int    `koanf:"port"`
	Host     string `koanf:"host"`
	DBName   string `koanf:"db_name"`
}

type MySQLDB struct {
	config Config
	db     *sql.DB
}

func New(config Config) (*MySQLDB, error) {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@(%s:%d)/%s",
		config.Username, config.Password, config.Host, config.Port, config.DBName))
	if err != nil {
		return &MySQLDB{}, fmt.Errorf("can't open mysql db: %v", err)
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return &MySQLDB{config: config, db: db}, nil
}
