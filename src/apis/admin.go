package apis

import (
	"github.com/gin-gonic/gin"
	"../models"
	"encoding/json"
)

func GetUsers(c *gin.Context) {
	var us []models.Users
	data, err := models.GetUsers()
	if err != nil {
		c.JSON(403, gin.H{
			"code":    403,
			"massage": "服务器响应错误",
		})
	}

	err = json.Unmarshal([]byte(data), &us)
	c.JSON(200, us)

}
