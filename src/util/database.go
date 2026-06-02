package util

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init(sqliteFile string) error {
	var err error

	DB, err = sql.Open("sqlite3", sqliteFile)
	if err != nil { LogError(err); return err }

	// create tables
	_, err = DB.Exec("CREATE TABLE IF NOT EXISTS users (id TEXT PRIMARY KEY, name TEXT, password TEXT, config TEXT")
	if err != nil { LogError(err); return err}

	LogInfo("database initialised")
	return DB.Ping()
}