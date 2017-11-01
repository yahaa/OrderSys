package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

var Router = gin.Default()



func init() {

	Router.Use(cors.Default())

	all := Router.Group("/api")
	all.Use(Auth.MiddlewareFunc())
	all.GET("/test", Test)

	Router.POST("/login", AuthAdmin.LoginHandler)
	admin := Router.Group("api/admin")
	admin.Use(AuthAdmin.MiddlewareFunc())
	admin.GET("/users", GetUsers)
	admin.POST("/user", AddUser)
	admin.DELETE("/user", DelUser)
	admin.PUT("/user", UpUser)
	admin.GET("/products", GetPros)
	admin.POST("/product", AddPro)
	admin.GET("/product", GetPro)
	admin.DELETE("/product", DelPro)
	admin.PUT("/product", UPPro)
	admin.POST("/saler", AddSaler)
	admin.DELETE("/saler", DelSaler)
	admin.GET("/salers", GetSals) //获取系统中所有销售员以及其销售情

	saler := Router.Group("api/saler")
	saler.Use(AuthSaler.MiddlewareFunc())
	saler.GET("/tasks", AllTasks) //获取自己的销售的所有商品
	saler.GET("/orders", AllOrds) //获取自己经办的所有订单 **未测试

	cus := Router.Group("api/cus")
	cus.Use(AuthCus.MiddlewareFunc())
	cus.GET("/products", AllSalePros) //获取所有销售员在销售的商品
	cus.GET("/orders", GetOrds)    //获取自己的订单
	cus.POST("/order", MakOrder)   //下订单
	cus.GET("/order", DelOrder)    //某个订单详情
	cus.PUT("/order",Play)
}
