package apis

import (
	"time"
	"github.com/gin-gonic/gin"
	"github.com/appleboy/gin-jwt"
	"../models"
	"log"
)

const (
	Key      = "OrderSysTokenKey"
	Real     = "orderSys"
	Admin    = "Admin"
	Saler    = "Saler"
	Customer = "Customer"
)

//用户登录时候认证用
func CheckLogin(userId string, password string, c *gin.Context) (string, bool) {
	uId, err := models.CheckUser(userId, password)
	if err != nil {
		log.Fatalln(err.Error())
		return "", false
	}

	if uId == "" {
		return "", false
	}
	return uId, true
}

//认证是管理员的权限
func authAdmin(userId string, c *gin.Context) bool {
	if models.CheckRole(userId, Admin) {
		return true
	}
	return false
}

//认证是销售员的权限
func authSaler(userId string, c *gin.Context) bool {
	if models.CheckRole(userId, Saler) {
		return true
	}
	return false
}

//认证是客户的权限
func authCus(userId string, c *gin.Context) bool {
	if models.CheckRole(userId, Customer) {
		return true
	}
	return false
}

func auth(userId string, c *gin.Context) bool {
	if models.CheckExisU(userId) {
		return true
	}
	return false
}

func UnAuth(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}

var AuthAdmin = &jwt.GinJWTMiddleware{
	Realm:         Admin,
	Key:           []byte(Key),
	Timeout:       time.Hour,
	MaxRefresh:    time.Hour,
	Authenticator: CheckLogin,
	Authorizator:  authAdmin,
	Unauthorized:  UnAuth,
	TokenLookup:   "header:Authorization",
	TokenHeadName: "Bearer",
	TimeFunc:      time.Now,
}

var AuthSaler = &jwt.GinJWTMiddleware{
	Realm:         Saler,
	Key:           []byte(Key),
	Timeout:       time.Hour,
	MaxRefresh:    time.Hour,
	Authenticator: CheckLogin,
	Authorizator:  authSaler,
	Unauthorized:  UnAuth,
	TokenLookup:   "header:Authorization",
	TokenHeadName: "Bearer",
	TimeFunc:      time.Now,
}

var AuthCus = &jwt.GinJWTMiddleware{
	Realm:         Customer,
	Key:           []byte(Key),
	Timeout:       time.Hour,
	MaxRefresh:    time.Hour,
	Authenticator: CheckLogin,
	Authorizator:  authCus,
	Unauthorized:  UnAuth,
	TokenLookup:   "header:Authorization",
	TokenHeadName: "Bearer",
	TimeFunc:      time.Now,
}

var Auth = &jwt.GinJWTMiddleware{
	Realm:         Real,
	Key:           []byte(Key),
	Timeout:       time.Hour,
	MaxRefresh:    time.Hour,
	Authenticator: CheckLogin,
	Authorizator:  auth,
	Unauthorized:  UnAuth,
	TokenLookup:   "header:Authorization",
	TokenHeadName: "Bearer",
	TimeFunc:      time.Now,
}
