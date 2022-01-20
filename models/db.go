package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func GetDb() *sql.DB {
	db, err := sql.Open("mysql", "root:123@(10.0.2.2:3306)/db1")
	if err != nil {
		log.Panic(err)
	}
	return db
}
