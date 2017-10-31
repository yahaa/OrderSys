package apis

import (
	"github.com/gin-gonic/gin"
	"../models"
	"fmt"
	"log"
	"encoding/json"
)

//查看自己销售的任务
func AllTasks(c *gin.Context) {
	uId := fmt.Sprint(c.Keys["userID"])
	data, err := models.AllPro(uId)
	var pros []models.Product
	if err != nil {
		log.Fatalln(err.Error())
		c.JSON(403, Bad)
		return
	}
	err = json.Unmarshal([]byte(data), &pros)
	c.JSON(200, pros)

}

//查看自己经办的所有订单
func AllOrds(c *gin.Context) {
	uId := fmt.Sprint(c.Keys["userID"])
	data, err := models.AllOrdsa(uId)
	var ords [] models.Orders
	if err != nil {
		log.Fatalln(err.Error())
		c.JSON(403, Bad)
		return
	}
	err = json.Unmarshal([]byte(data), &ords)
	c.JSON(200, ords)

}
