package models

import (
	"github.com/golang/protobuf/ptypes/timestamp"
)

type Users struct {
	UserId   int
	password string
	Name     string
	Sex      string
	Phone    string
	RoleId   int
}

type Role struct {
	RoleId   int
	RoleName string
}

type Permit struct {
	PerId   int                    `权限id`
	PerName string                 `权限名,名字和opt对用的函数指针名一样`
	opt     *func(arg interface{}) `该权限可以做的事情`
}

type PermitRole struct {
	RoleId int
	PerId  int
}

type Orders struct {
	OrdId       int                 `订单的唯一标识`
	UserId      int                 `该订单所属的用户`
	SaleId      int                 `售货员id`
	ProId       int                 `产品编号`
	N           int                 `产品数量`
	Total       float64             `订单总价格`
	OrderTime   timestamp.Timestamp `下单时间`
	Marks       string              `备注`
	OrdCom      string              `购买公司或个人`
	PlayTime    timestamp.Timestamp `支付时间`
	SaleType    string              `销售方式：赠送、销售、样品、返厂、代发等`
	DeliverType string              `发货形式：邮寄现付、邮寄到付、邮寄垫付、自取`
}

type Product struct {
	ProId     int
	ProName   string
	ProSM     string              `产品规格型号`
	ProCom    string
	ProPrice  float64
	ProCounts int
	ProStyle  string              `产品款式`
	OrderTime timestamp.Timestamp `进货时间`
	Marks     string
}

type UserSale struct {
	UserId int
	ProId  int
	Marks  string `备注`
}