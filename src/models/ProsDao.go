package models

import (
	"github.com/elgs/gosqljson"
	"log"
	db "../dataSource"
	"fmt"
)

func GetPros() (string, error) {
	sq := "select * from PRODUCT where USE=0"
	data, err := gosqljson.QueryDbToMapJSON(db.RDB, Camel, sq)
	if err != nil {
		log.Fatalln(err.Error())
		return "", err

	}
	return data, err
}

func GetPro(proId string) (string, error) {
	sq := "select * from PRODUCT where PROID='%s' and USE=0"
	sq = fmt.Sprintf(sq, proId)
	data, err := gosqljson.QueryDbToMapJSON(db.RDB, Camel, sq)
	if err != nil {
		log.Fatalln(err.Error())
		return "", err
	}
	return data, err
}

func AddPro(pro *Product) bool {
	sq := "insert into PRODUCT(PROID,PRONAME,PROSM,PROCOM,PROPRICE,PROCOUNTS,PROSTYLE,ORDERTIME,MARKS)" +
		" values(%d,'%s','%s','%s','%f','%d','%s','%s','%s')"
	sq = fmt.Sprintf(sq, pro.ProId, pro.ProName, pro.ProSM,
		pro.ProCom, pro.ProPrice, pro.ProCounts, pro.ProStyle, pro.OrderTime, pro.Marks)

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

func DelPro(proId int) bool {
	sq := "update PRODUCT set USE=1 where PROID='%d'"
	sq = fmt.Sprintf(sq, proId)
	_, err := db.RDB.Exec(sq)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

func UpPro(pro *Product) bool {
	sq := "update PRODUCT set PRONAME='%s',PROSM='%s',PROCOM='%s',PROPRICE='%f'," +
		"PROCOUNTS='%d',PROSTYLE='%s',ORDERTIME='%s',MARKS='%s' where PROID='%d'"
	sq = fmt.Sprintf(sq, pro.ProName, pro.ProSM, pro.ProCom, pro.ProPrice,
		pro.ProCounts, pro.ProStyle, pro.OrderTime, pro.Marks,pro.ProId)
	fmt.Println(sq)
	_, err := db.RDB.Exec(sq)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}
