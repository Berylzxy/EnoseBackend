package usercontroller

import (
	"EnoseBackend/model"
	"EnoseBackend/utils"
	"errors"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	_, err = model.GetUserByName(req.Username)
	if err != nil {
		user := new(model.User)
		user.Name = req.Username
		user.Password = utils.PasswordEncrypt(req.Password)
		fmt.Println(req)
		err = model.AddUser(user)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 1, "error": err.Error()})
			return
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"code": 1, "success": false, "message": "用户名已存在"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "success": true, "message": "注册成功"})

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
		c.JSON(http.StatusOK, gin.H{"code": 1, "error": err.Error()})
		return
	}

	if !utils.PasswordVerify(req.Password, user.Password) {
		c.JSON(http.StatusOK, gin.H{"code": 1, "error": "用户名或密码错误"})
		return
	}

	//token := jwt.New(jwt.SigningMethodHS256)
	//claims := token.Claims.(jwt.MapClaims)
	//claims["userId"] = user.ID
	//claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	//
	//tokenString, err := token.SignedString([]byte("your-secret-key"))
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"code": 1, "error": err.Error()})
	//	return
	//}
	//
	//resp := new(UserInfoResponseBody)
	//resp.ID = user.ID
	//resp.Username = user.Name
	//
	//// 将Token写入Cookie
	//c.SetCookie("myCookie", tokenString, 3600, "/", "localhost", false, true)
	//
	//c.JSON(http.StatusOK, gin.H{"code": 0, "success": true, "data": resp})
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
	//c.SetCookie("myCookie", "123456", 3600, "/", "localhost", false, true)
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
	cookie, err := c.Request.Cookie("session") //这个b名卡我一周
	fmt.Println(cookie)
	if err != nil {
		// 未找到 cookie
		// 处理未找到 cookie 的情况
		c.JSON(401, gin.H{"data": "unauthorized"})
	} else {
		// 找到 cookie
		// 使用 cookie.Value 来访问 cookie 值
		value := cookie.Value
		// 处理找到 cookie 的情况
		fmt.Println(value)
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

}

type ListRes struct {
	key      string
	username string
	identity string
	lasttime string
}

func Listuser(c *gin.Context) {
	User, _ := model.ListUser()
	res := []ListRes{}
	for i, val := range *User {
		tmp := ListRes{}
		tmp.key = strconv.Itoa(int(val.ID))
		tmp.username = val.Name
		tmp.identity = val.Identity
		res = append(res, tmp)
		fmt.Println(i, res)
	}
	c.ShouldBind(res)
	c.JSON(200, res)
}

type UserModifyPwd struct {
	Username    string
	Password    string
	Newpassword string
}

func ModifyPwd(c *gin.Context) {
	req := new(UserModifyPwd)
	err := c.ShouldBind(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "error": err.Error()})
		return
	}
	fmt.Println("req", req)
	user, err := model.GetUserByName(req.Username)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"code": 1, "error": err.Error()})
		return
	}

	if !utils.PasswordVerify(req.Password, user.Password) {
		fmt.Println("??????????")
		c.JSON(http.StatusOK, gin.H{"code": 1, "success": false, "message": "密码错误"})
		return
	}

	user.Password = utils.PasswordEncrypt(req.Newpassword)
	model.UpdateUser(user)
	c.JSON(http.StatusOK, gin.H{"code": 0, "success": true, "message": "修改成功"})
}

//func UserForget(c *gin.Context) {
//	req := new(UserForgetRequestBody)
//	err := c.BindJSON(&req)
//	////usercontroller :=new(model.User)
//	//usercontroller:=new(model.User)
//	//fmt.Print(req.Username)
//	user, err := model.GetUserByName(req.Username)
//	//password := user.Password
//	if err != nil {
//		c.JSON(200, gin.H{"message": "用户名不存在"})
//	} else {
//		c.JSON(200, gin.H{"message": "请联系管理员"})
//	}
//	return
//}
