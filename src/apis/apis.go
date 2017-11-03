package apis

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"time"
)

var Router = gin.Default()

func init() {
	config := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "HEAD", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-type", "Authorization", "Accept"},
		AllowAllOrigins:  true,
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	co := cors.New(config)

	Router.Use(co)
	all := Router.Group("/api")
	all.GET("/products", AllSalePros)

	login := Router.Group("/api/self")
	login.Use(Auth.MiddlewareFunc())
	login.GET("/info", GetSelf)

	Router.POST("/login", AuthAdmin.LoginHandler)
	admin := Router.Group("api/admin")
	admin.Use(AuthAdmin.MiddlewareFunc())
	admin.GET("/users", GetUsers)
	admin.DELETE("/users/:userId", DelUser)
	admin.POST("/user", AddUser)
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
	cus.GET("/orders", GetOrds)            //获取自己的订单
	cus.DELETE("/orders/:ordId", DelOrder) //某个订单详情
	cus.PUT("/orders/:ordId", Play)
	cus.POST("/order", MakOrder) //下订单

}
