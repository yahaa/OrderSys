package models

import (
	db "../dataSource"
	"fmt"
	"log"
	"github.com/elgs/gosqljson"
)

const Def = "good job !"

func AddSaler(userId, proId int, marks string) bool {
	if marks == "" {
		marks = Def
	}
	sq := "insert into USERSALER(USERID,PROID,MARKS) values('%d','%d','%s')"
	sq = fmt.Sprintf(sq, userId, proId, marks)
	_, err := db.RDB.Exec(sq)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

func DelSaler(userId, proId int) bool {
	sq := "delete from USERSALER where USERID='%d' and PROID='%d'"
	sq = fmt.Sprintf(sq, userId, proId)
	_, err := db.RDB.Exec(sq)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

func GetSales() (string, error) {
	sq := "select USERID,PROID,MARKS from USERSALER"
	data, err := gosqljson.QueryDbToMapJSON(db.RDB, Camel, sq)
	if err != nil {
		log.Fatalln(err.Error())
		return "", err
	}
	return data, err
}
