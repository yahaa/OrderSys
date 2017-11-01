package models

import (
	db "../dataSource"
	"github.com/elgs/gosqljson"
	"log"
)

func ExeSQL(sq string) bool {
	res, err := db.RDB.Exec(sq)
	if err != nil {
		log.Println(err.Error())
		return false
	}

	if n, _ := res.LastInsertId(); n <= 0 {
		return false
	}
	return true
}

func ExeSQLR(sq, m string) (string, error) {
	data, err := gosqljson.QueryDbToMapJSON(db.RDB, m, sq)
	if err != nil {
		log.Fatalln(err.Error())
		return "", err
	}
	return data, err
}
