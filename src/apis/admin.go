package apis

import (
	"github.com/gin-gonic/gin"
	"../models"
	"encoding/json"
	"../util"
	"log"
)

var Good = gin.H{
	"code":    "200",
	"message": "success",
}

var Bad = gin.H{
	"code":    "400",
	"message": "param is bad!",
}

func GetUsers(c *gin.Context) {
	var us []models.User
	data, err := models.GetUsers()
	if err != nil {
		c.JSON(403, Bad)
	}
	err = json.Unmarshal([]byte(data), &us)
	c.JSON(200, us)
}

func AddUser(c *gin.Context) {
	var u models.User
	c.BindJSON(&u)
	u.Password = util.Md5(u.Password)
	if models.AddUser(&u) {
		c.JSON(200, Good)
	} else {
		c.JSON(400, Bad)
	}
}

func DelUser(c *gin.Context) {
	var u models.User
	c.BindJSON(&u)
	if models.DelUser(u.UserId) {
		c.JSON(200, Good)
	} else {
		c.JSON(400, Bad)
	}

}

func UpUser(c *gin.Context) {
	var u models.User
	c.BindJSON(&u)
	if models.UpUser(&u) {
		c.JSON(200, Good)
	} else {
		c.JSON(400, Bad)
	}

}

func GetPros(c *gin.Context) {
	var pros []models.Product
	data, err := models.GetPros()
	if err != nil {
		log.Fatalln(err.Error())
		c.JSON(403, Bad)
	}
	err = json.Unmarshal([]byte(data), &pros)
	c.JSON(200, pros)

}

func AddPro(c *gin.Context) {
	var pro models.Product
	c.BindJSON(&pro)
	if models.AddPro(&pro) {
		c.JSON(200, Good)
	} else {
		c.JSON(400, Bad)
	}
}

func GetPro(c *gin.Context) {
	proId := c.Query("proId")
	var pros []models.Product
	data, err := models.GetPro(proId)
	if err != nil {
		log.Fatalln(err.Error())
		c.JSON(403, Bad)
		return
	}
	err = json.Unmarshal([]byte(data), &pros)
	if len(pros) <= 0 {
		c.JSON(403, Bad)
		return
	}
	c.JSON(200, pros[0])
}

func DelPro(c *gin.Context) {
	var pro models.Product
	c.BindJSON(&pro)
	if models.DelPro(pro.ProId) {
		c.JSON(200, Good)
	} else {
		c.JSON(400, Bad)
	}
}

func UPPro(c *gin.Context) {
	var pro models.Product
	c.BindJSON(&pro)
	if models.UpPro(&pro) {
		c.JSON(200, Good)
	} else {
		c.JSON(400, Bad)
	}

}

func AddSaler(c *gin.Context) {
	var sal models.UserSale
	c.BindJSON(&sal)
	if models.AddSaler(sal.UserId, sal.ProId, sal.Marks) {
		c.JSON(200, Good)
	} else {
		c.JSON(400, Bad)
	}
}

func DelSaler(c *gin.Context) {
	var sal models.UserSale
	c.BindJSON(&sal)
	if models.DelSaler(sal.UserId, sal.ProId) {
		c.JSON(200, Good)
	} else {
		c.JSON(400, Bad)
	}
}
func GetSals(c *gin.Context) {
	var us []models.UserSale
	data, err := models.GetSales()
	if err != nil {
		log.Fatalln(err.Error())
		c.JSON(403, Bad)
	}
	err = json.Unmarshal([]byte(data), &us)
	c.JSON(200, us)
}
