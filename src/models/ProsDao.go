package models

import (
	"fmt"
)

func GetPros() (string, error) {
	sq := "select * from PRODUCT where USE=0"
	return ExeSQLR(sq, Camel)
}

func GetPro(proId string) (string, error) {
	sq := "select * from PRODUCT where PROID='%s' and USE=0"
	sq = fmt.Sprintf(sq, proId)
	return ExeSQLR(sq, Camel)
}

func AddPro(pro *Product) bool {
	sq := "insert into PRODUCT(PROID,PRONAME,PROSM,PROCOM,PROPRICE,PROCOUNTS,PROSTYLE,ORDERTIME,MARKS)" +
		" values(%d,'%s','%s','%s','%f','%d','%s','%s','%s')"
	sq = fmt.Sprintf(sq, pro.ProId, pro.ProName, pro.ProSM,
		pro.ProCom, pro.ProPrice, pro.ProCounts, pro.ProStyle, pro.OrderTime, pro.Marks)

	return ExeSQL(sq)
}

func DelPro(proId int64) bool {
	sq := "update PRODUCT set USE=1 where PROID='%d'"
	sq = fmt.Sprintf(sq, proId)
	return ExeSQL(sq)
}

func UpPro(pro *Product) bool {
	sq := "update PRODUCT set PRONAME='%s',PROSM='%s',PROCOM='%s',PROPRICE='%f'," +
		"PROCOUNTS='%d',PROSTYLE='%s',ORDERTIME='%s',MARKS='%s' where PROID='%d'"
	sq = fmt.Sprintf(sq, pro.ProName, pro.ProSM, pro.ProCom, pro.ProPrice,
		pro.ProCounts, pro.ProStyle, pro.OrderTime, pro.Marks, pro.ProId)
	return ExeSQL(sq)
}

func AllPro(userId string) (string, error) {
	sq := "select * from PRODUCT where PROID in (select PROID from USERSALER where USERID='%s')"
	sq = fmt.Sprintf(sq, userId)
	return ExeSQLR(sq, Camel)
}

func AllSalePro() (string, error) {
	sq := "select s.USERID salerId,p.*,u.NAME salerName, u.phone from " +
		"PRODUCT p,USERSALER s,USERS u where p.PROID=s.PROID and u.USERID=s.USERID and s.USE=0"
	return ExeSQLR(sq, Camel)
}
