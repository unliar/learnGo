package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"learnGo/examples/go-curl/conf"
	"log"
)

var MySQL *sql.DB

func init() {
	var err error
	MySQL, err = sql.Open("mysql", conf.CurrentConf.MySQL)
	if err != nil {
		log.Fatal(err)
		panic(err.Error())
	}
	MySQL.SetMaxOpenConns(20)
	MySQL.SetMaxIdleConns(1000)
	MySQL.Ping()

}
