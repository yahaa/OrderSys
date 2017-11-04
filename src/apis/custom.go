package apis

import (
	"github.com/gin-gonic/gin"
	"log"
	"../models"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
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

func GetSelf(c *gin.Context) {
	uId := fmt.Sprint(c.Keys["userID"])
	u, err := models.GetSelf(uId)
	if err != nil {
		log.Println(err.Error())
		c.JSON(403, Bad)
		return
	}

	c.JSON(200, u)
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
	uId, err := strconv.Atoi(fmt.Sprint(c.Keys["userID"]))
	if err != nil {
		fmt.Println(789)
		c.JSON(400, Bad)
		return
	}
	ti := time.Now()
	var ords models.Orders
	c.BindJSON(&ords)
	ords.UserId = uId
	ords.OrdId = ti.Unix()
	ords.OrderTime = ti.String()
	ords.Total = ords.Price * float64(ords.Nums)
	if models.MakOrder(&ords) {
		c.JSON(200, Good)
		return
	}
	fmt.Println(32323)
	c.JSON(400, Bad)

}

//退单
func DelOrder(c *gin.Context) {
	uId := fmt.Sprint(c.Keys["userID"])
	ordId := c.Param("ordId")
	if models.DelOrd(ordId, uId) {
		c.JSON(200, Good)
	} else {
		c.JSON(400, Bad)
	}
}

func Play(c *gin.Context) {
	uId := fmt.Sprint(c.Keys["userID"])
	ordId := c.Param("ordId")
	if models.Play(ordId, uId) {
		c.JSON(200, Good)
	} else {
		c.JSON(400, Bad)
	}
}

func Test(c *gin.Context) {
	models.TT()
	c.JSON(200, gin.H{
		"status": "ok",
		"name":   "bad",
	})
}
