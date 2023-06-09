package models

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"todo/config"
)

var Db *sql.DB
var err error

const (
    tableNameUser    = "users"
    tableNameTodo    = "todos"
    tableNameSession = "sessions"
)

func init() {
    Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
    if err != nil {
        log.Fatalln(err)
    }

    createUser := `CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`

    cmdU := fmt.Sprintf(createUser, tableNameUser)
    Db.Exec(cmdU)

    createTodo := `CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME)`
    cmdT := fmt.Sprintf(createTodo, tableNameTodo)
    Db.Exec(cmdT)

    createSession := `CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid INTEGER NOT NULL UNIQUE,
		email STRING,
		user_id INTEGER,
		created_at DATETIME)`

    cmdS := fmt.Sprintf(createSession, tableNameSession)
    Db.Exec(cmdS)

}
