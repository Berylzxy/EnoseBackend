package usercontroller

import (
	"EnoseBackend/dao"
	"EnoseBackend/model"
	"EnoseBackend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type UserRequestBody struct {
	Identity string
	Name     string
	Password string
}

func ListUser(c *gin.Context) {
	var user []model.User
	if err := dao.DB.Find(&user).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, user)
	}
}

func DealUser(c *gin.Context) {
	req := new(UserRequestBody)
	c.BindJSON(&req)
	user, err := model.GetUserByName(req.Name)
	user.Identity = req.Identity
	user.Name = req.Name
	user.Password = utils.PasswordEncrypt(req.Password)
	if err != nil {
		model.AddUser(user)
		c.JSON(200, gin.H{"message": "添加成功"})
	} else {
		model.UpdateUser(user)
		c.JSON(200, gin.H{"message": "修改成功"})
	}
}
