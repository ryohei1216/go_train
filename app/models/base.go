package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"main/config"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)


var Db *sql.DB

var err error

const (
	tableNameUser = "users"
	tableNameTodos = "todos"
)

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err!= nil {
		log.Fatalln(err)
	}

	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)
	Db.Exec(cmdU)

	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME )`, tableNameTodos)
	Db.Exec(cmdT)
}

func createUUID() (uuidObj uuid.UUID) {
	uuidObj, err = uuid.NewUUID()
	return uuidObj
}

func Encrypt(planetext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(planetext)))
	return cryptext
}