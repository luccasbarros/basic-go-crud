package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" // Connection driver of mysql
)

func Connection() (*sql.DB, error) {
	connectionString := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", connectionString)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		return nil, erro
	}

	return db, nil
}
