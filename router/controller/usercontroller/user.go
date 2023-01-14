package usercontroller

import (
	"EnoseBackend/model"
	"EnoseBackend/utils"
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserSignUpRequestBody struct {
	Username string
	Password string
}

type UserSignInRequestBody struct {
	Username string
	Password string
}

type UserInfoResponseBody struct {
	ID       uint
	Username string
}
type UserForgetRequestBody struct {
	Username string
}
type UserSignUpResponseBody struct {
	ID       uint
	Username string
}

func (stu *UserSignUpRequestBody) Validate() (err error) {

	if len(stu.Username) < 2 {
		err = errors.New("用户名长度必须大于2")
	}
	if len(stu.Password) < 6 {
		err = errors.New("密码长度不能小于6")
	}
	return
}

func UserSignUp(c *gin.Context) {
	req := new(UserSignUpRequestBody)
	err := c.ShouldBind(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "error": err.Error()})
		return
	}

	err = req.Validate()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "error": err.Error()})
		return
	}

	user := new(model.User)
	user.Name = req.Username
	user.Password = utils.PasswordEncrypt(req.Password)
	fmt.Println(req)
	err = model.AddUser(user)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "success": true})

}

func UserSignIn(c *gin.Context) {
	req := new(UserSignInRequestBody)
	err := c.ShouldBind(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "error": err.Error()})
		return
	}
	fmt.Println("req", req)
	user, err := model.GetUserByName(req.Username)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "error": err.Error()})
		return
	}

	if !utils.PasswordVerify(req.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "error": "用户名或密码错误"})
		return
	}

	session := sessions.Default(c)
	session.Set("userId", user.ID)
	err = session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "error": err.Error()})
		return
	}

	resp := new(UserInfoResponseBody)
	resp.ID = user.ID
	resp.Username = user.Name

	c.JSON(http.StatusOK, gin.H{"code": 0, "success": true, "data": resp})
}
func UserSignOut(c *gin.Context) {

	session := sessions.Default(c)
	session.Clear()
	//session.Delete("userId")
	err := session.Save()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0})
}
func UserInfo(c *gin.Context) {

	session := sessions.Default(c)

	userId := session.Get("userId")

	if userId == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户未登录"})
		return
	}

	user, err := model.GetUserById(userId.(uint))

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "用户未登录"})
		return
	}

	resp := new(UserInfoResponseBody)
	resp.ID = user.ID
	resp.Username = user.Name

	c.JSON(http.StatusOK, gin.H{"data": resp})
}
func UserForget(c *gin.Context) {
	req := new(UserForgetRequestBody)
	err := c.BindJSON(&req)
	////usercontroller :=new(model.User)
	//usercontroller:=new(model.User)
	//fmt.Print(req.Username)
	user, err := model.GetUserByName(req.Username)
	password := user.Password
	if err != nil {
		c.JSON(200, gin.H{"message": password})
	}
	return
}
