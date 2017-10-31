package apis

import (
	"github.com/gin-gonic/gin"
	"log"
	"../models"
	"encoding/json"
	"fmt"
)

type SPU struct {
	SalerId   string `json:"salerId"`
	SalerName string `json:"salerName"`
	models.Product
	models.User
}

//获取所有在销售的商品
func AllSalePros(c *gin.Context) {
	var pros []SPU
	data, err := models.AllSalePro()
	if err != nil {
		log.Fatalln(err.Error())
		c.JSON(403, Bad)
	}
	err = json.Unmarshal([]byte(data), &pros)
	c.JSON(200, pros)

}

//获取所有订单
func GetOrds(c *gin.Context) {
	uId := fmt.Sprint(c.Keys["userID"])
	data, err := models.AllOrdsc(uId)
	var ords [] models.Orders
	if err != nil {
		log.Fatalln(err.Error())
		c.JSON(403, Bad)
		return
	}
	err = json.Unmarshal([]byte(data), &ords)
	c.JSON(200, ords)

}

//下订单
func MakOrder(c *gin.Context) {

}

//点单详情
func DetOrder(c *gin.Context) {

}
