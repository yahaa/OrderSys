package models

import (
	"fmt"
	"time"
)

func AllOrdsa(salerId string) (string, error) {
	sq := "select * from ORDERS where SALEID='%s' and USE=0"
	sq = fmt.Sprintf(sq, salerId)

	return ExeSQLR(sq, Camel)
}

func AllOrdsc(cusId string) (string, error) {
	sq := "select * from ORDERS where USERID='%s' and USE=0"
	sq = fmt.Sprintf(sq, cusId)
	return ExeSQLR(sq, Camel)
}

func AllOrds() (string, error) {
	sq := "select * from ORDERS"
	return ExeSQLR(sq, Camel)
}

func MakOrder(ord *Orders) bool {
	sq := "insert into ORDERS(PROID,ORDID,USERID,SALEID,NUMS,TOTAL,ORDERTIME,MARKS," +
		"ORDCOM,SALETYPE,DELIVERTYPE,PRICE) values('%d','%d','%d','%d','%d','%f'," +
		"'%s','%s','%s','%s','%s','%f')"
	sq = fmt.Sprintf(sq, ord.ProId, ord.OrdId, ord.UserId, ord.SaleId, ord.Nums, ord.Total, ord.OrderTime, ord.Marks,
		ord.OrdCom, ord.SaleType, ord.DeliverType, ord.Price)
	return ExeSQL(sq)

}

func DelOrd(ordId, uId string) bool {
	sq := "update ORDERS set USE=1 where ORDID='%s' and USERID='%s'"
	sq = fmt.Sprintf(sq, ordId, uId)
	return ExeSQL(sq)
}

func Play(ordId, uId string) bool {
	t := time.Now().UTC()
	sq := "update ORDERS set PLAYTIME='%s' where ORDID='%s' and USERID='%s'"
	sq = fmt.Sprintf(sq, t, ordId, uId)
	return ExeSQL(sq)
}
