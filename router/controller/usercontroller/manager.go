package usercontroller

import (
	"EnoseBackend/dao"
	"EnoseBackend/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

func ListUser(c *gin.Context) {
	var user []model.User
	if err := dao.DB.Find(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}
}
