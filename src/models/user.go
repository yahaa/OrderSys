package models

import (
	db "../dataSource"
	"fmt"
	"log"
	"github.com/elgs/gosqljson"
)

const (
	Lower = "lower"
	Upper = "upper"
	Camel = "camel"
)

//User的方法
//检查用户名,密码是否正确，返回用户id
func CheckUser(uId string, password string) (string, error) {
	var userId string
	var err error
	sq := "select USERID from USERS WHERE USERID=%s and PASSWORD=%s"
	sq = fmt.Sprintf(sq, uId, password)
	rows, err := db.RDB.Query(sq)
	if err != nil {
		log.Fatalln(err.Error())
		return userId, err
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&userId)
		if err == nil {
			break
		}
	}
	return userId, err
}

func CheckRole(userId, role string) bool {
	sq := "select ROLENAME from ROLE where ROLEID in" +
		"(select ROLEID from USERS where USERID=%s)"
	sq = fmt.Sprintf(sq, userId)
	rows, err := db.RDB.Query(sq)
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer rows.Close()
	var roleName string
	for rows.Next() {
		rows.Scan(&roleName)
		if roleName == role {
			return true
		}
	}
	return false
}

func GetUsers() (string, error) {
	sq := "select USERID,NAME,SEX,PHONE,ROLEID from USERS"
	data, err := gosqljson.QueryDbToMapJSON(db.RDB, Camel, sq)
	if err != nil {
		return "", err
		log.Fatalln(err.Error())
	}
	return data, err
}
