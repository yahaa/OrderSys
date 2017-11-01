package models

import (
	"fmt"
)

const Def = "good job !"

func AddSaler(userId, proId int, marks string) bool {
	if marks == "" {
		marks = Def
	}
	sq := "insert into USERSALER(USERID,PROID,MARKS) values('%d','%d','%s')"
	sq = fmt.Sprintf(sq, userId, proId, marks)
	return ExeSQL(sq)
}

func DelSaler(userId, proId int) bool {
	sq := "delete from USERSALER where USERID='%d' and PROID='%d'"
	sq = fmt.Sprintf(sq, userId, proId)
	return ExeSQL(sq)
}

func GetSales() (string, error) {
	sq := "select USERID,PROID,MARKS from USERSALER"
	return ExeSQLR(sq, Camel)
}
