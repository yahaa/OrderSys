package dataSource

import (
	"database/sql"
	"log"
	_ "github.com/mattn/go-oci8"
)

var RDB *sql.DB

func init() {
	var err error
	RDB, err = sql.Open("oci8", "zihua/123456@127.0.0.1:1521/XE")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = RDB.Ping()

	if err != nil {
		log.Fatal(err.Error())
	}
	log.Println("connect to oracle sucessful !")
}
