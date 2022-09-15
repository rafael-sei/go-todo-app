package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"todo_app/config"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

var err error

// tablename作成
const (
	tableNameUser    = "users"
	tableNameTodo    = "todos"
	tableNameSession = "sessions"
)

// table作成
func init() {
	Db, err = sql.Open(config.Config.SQLDriver, config.Config.DbName)
	if err != nil {
		log.Fatal(err)
	}

	// #users tabel作成コマンド
	cmdU := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STIRNG,
		email STRING,
		password STRING,
		created_at DATETIME)`, tableNameUser)
	// cmdU実行
	Db.Exec(cmdU)

	// #todos tabel作成コマンド
	cmdT := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		content TEXT,
		user_id INTEGER,
		created_at DATETIME)`, tableNameTodo)
	// cmdU実行
	Db.Exec(cmdT)

	// #sessoins tabel作成コマンド
	cmdS := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STIRNG,
		user_id INTEGER,
		created_at DATETIME)`, tableNameSession)
	// cmdU実行
	Db.Exec(cmdS)
}

// UUID作成(uuidパッケージを使う)
func createUUID() (uuidobj uuid.UUID) {
	// uuid生成
	uuidobj, err = uuid.NewUUID()
	if err != nil {
		log.Fatalln(err)
	}
	return uuidobj
}

// passwordをハッシュ値に変える
func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
