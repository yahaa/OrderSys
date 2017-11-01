package models

import (
	db "../dataSource"
	"../util"
	"fmt"
	"log"
)

const (
	Lower = "lower"
	Upper = "upper"
	Camel = "camel"
)

//检查用户名,密码是否正确，返回用户id
func CheckUser(uId string, password string) (string, error) {
	password = util.Md5(password)
	var userId string
	var err error
	sq := "select USERID from USERS WHERE USERID=%s and PASSWORD='%s' and USE=0"
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

func CheckExisU(uId string) bool {
	var userId string
	sq := "select USERID from USERS WHERE USERID=%s"
	sq = fmt.Sprintf(sq, uId)
	row := db.RDB.QueryRow(sq)
	err := row.Scan(&userId)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	if userId == "" {
		return false
	}
	return true
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
	sq := "select USERID,NAME,SEX,PHONE,ROLEID from USERS WHERE USE=0"
	return ExeSQLR(sq, Camel)
}

func AddUser(u *User) bool {
	sq := "insert into USERS(USERID,PASSWORD,NAME,SEX,PHONE) values(%d,'%s','%s','%s','%s')"
	sq = fmt.Sprintf(sq, u.UserId, u.Password, u.Name, u.Sex, u.Phone)
	return ExeSQL(sq)
}

func DelUser(uId int) bool {
	sq := "update USERS set USE=1 where USERID=%d"
	sq = fmt.Sprintf(sq, uId)
	return ExeSQL(sq)
}

func UpUser(u *User) bool {
	sq := "update USERS set NAME='%s',SEX='%s',PHONE='%s',ROLEID=%d where USERID=%d"
	sq = fmt.Sprintf(sq, u.Name, u.Sex, u.Phone, u.RoleId, u.UserId)
	return ExeSQL(sq)
}

func TT() bool {
	sq := "INSERT INTO TEST(TID,T) VALUES (1,'1509515746')"
	return ExeSQL(sq)
}
