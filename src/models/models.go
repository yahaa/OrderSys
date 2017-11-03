package models

type User struct {
	UserId   int    `json:"userId,string,omitempty"`
	Password string `json:"password,omitempty"`
	Name     string `json:"name,omitempty"`
	Sex      string `json:"sex,omitempty"`
	Phone    string `json:"phone,omitempty"`
	RoleId   int    `json:"roleId,string,omitempty"`
}

type Role struct {
	RoleId   int    `json:"roleId,stirng"`
	RoleName string `json:"roleNae"`
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
	OrdId       int64   `json:"ordId,string"`  //订单的唯一标识 n
	UserId      int     `json:"userId,string"` //该订单所属的用户 n
	SaleId      int     `json:"saleId,string"` //售货员id
	ProId       int     `json:"proId,string"`  //产品编号
	Nums        int     `json:"nums,string"`   //产品数量
	Total       float64 `json:"total,string"`  //订单总价格
	OrderTime   string  `json:"orderTime"`     //下单时间 n
	Marks       string  `json:"marks"`         //备注
	Price       float64 `json:"price,string"`  //产品单价价钱
	OrdCom      string  `json:"ordCom"`        //购买公司或个人
	PlayTime    string  `json:"playTime"`      //支付时间
	SaleType    string  `json:"saleType"`      //销售方式：赠送、销售、样品、返厂、代发等
	DeliverType string  `json:"deliverType"`   //发货形式：邮寄现付、邮寄到付、邮寄垫付、自取
}

type Product struct {
	ProId     int64   `json:"proId,string"`
	ProName   string  `json:"proName"`
	ProSM     string  `json:"proSM"` //产品规格型号
	ProCom    string  `json:"proCom"`
	ProPrice  float64 `json:"proPrice,string"`
	ProCounts int     `json:"proCounts,string"`
	ProStyle  string  `json:"proStye"`   //产品款式
	OrderTime string  `json:"orderTime"` //进货时间
	Marks     string  `json:"marks"`
}

type UserSale struct {
	UserId int    `json:"userId,string"`
	ProId  int    `json:"proId,string"`
	Marks  string `json:"marks"`
}
