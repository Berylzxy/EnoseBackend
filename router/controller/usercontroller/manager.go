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
	Username string
	Password string
}
type Userdel struct {
	Name string
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

	user, err := model.GetUserByName(req.Username)
	fmt.Println("req", req)
	if err != nil {
		user := new(model.User)
		user.Identity = req.Identity
		user.Name = req.Username
		user.Password = utils.PasswordEncrypt(req.Password)

		model.AddUser(user)
		c.JSON(200, gin.H{"message": "添加成功"})
	} else {

		user.Identity = req.Identity
		user.Name = req.Username
		user.Password = utils.PasswordEncrypt(req.Password)
		model.UpdateUser(user)
		c.JSON(200, gin.H{"message": "修改成功"})
	}
}
func DelUser(c *gin.Context) {
	req := new(Userdel)
	c.BindJSON(&req)
	fmt.Println(req)
	user, _ := model.GetUserByName(req.Name)
	model.DeleteUser(user)
	c.JSON(200, gin.H{"message": "success"})
}
