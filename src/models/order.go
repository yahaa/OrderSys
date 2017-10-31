package models

import (
	db "../dataSource"
	"fmt"
	"github.com/elgs/gosqljson"
	"log"
)

func AllOrdsa(salerId string) (string, error) {
	sq := "select * from ORDERS where SALEID='%d'"
	sq = fmt.Sprintf(sq, salerId)

	data, err := gosqljson.QueryDbToMapJSON(db.RDB, Camel, sq)
	if err != nil {
		return "", err
		log.Fatalln(err.Error())
	}
	return data, err
}

func AllOrdsc(cusId string) (string, error) {
	sq := "select * from ORDERS where USERID='%d'"
	sq = fmt.Sprintf(sq, cusId)

	data, err := gosqljson.QueryDbToMapJSON(db.RDB, Camel, sq)
	if err != nil {
		return "", err
		log.Fatalln(err.Error())
	}
	return data, err
}

func AllOrds() (string, error) {
	sq := "select * from ORDERS"
	data, err := gosqljson.QueryDbToMapJSON(db.RDB, Camel, sq)
	if err != nil {
		return "", err
		log.Fatalln(err.Error())
	}
	return data, err

}
