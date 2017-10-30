package apis

import "github.com/gin-gonic/gin"

var Router = gin.Default()

func test(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
		"host":   "localhost",
	})
}

func init() {
	Router.POST("/login", AuthAdmin.LoginHandler)
	admin := Router.Group("/admin")
	admin.Use(AuthAdmin.MiddlewareFunc())
	admin.GET("/users", GetUsers)

	saler := Router.Group("/saler")
	saler.Use(AuthSaler.MiddlewareFunc())
	saler.GET("/test", test)

	cus := Router.Group("/cus")
	cus.Use(AuthCus.MiddlewareFunc())
	cus.GET("/test")
}
