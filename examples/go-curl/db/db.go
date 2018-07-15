package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var MySQL *sql.DB

func init() {
	var err error
	MySQL, err = sql.Open("mysql", "root:MySQL19930224@tcp(127.0.0.1:3306)/go-server")
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
	MySQL.SetMaxOpenConns(20)
	MySQL.SetMaxIdleConns(1000)
	MySQL.Ping()

}
