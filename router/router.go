package router

import (
	"EnoseBackend/router/controller/experimentcontroller"
	"EnoseBackend/router/controller/smpcontroller"
	"EnoseBackend/router/controller/usercontroller"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	userGroup := r.Group("/user")
	{
		userGroup.POST("/signIn", usercontroller.UserSignIn)
		userGroup.POST("/signUp", usercontroller.UserSignUp)
		userGroup.POST("/forget", usercontroller.UserForget)
		userGroup.GET("/info", usercontroller.UserInfo)
		userGroup.GET("/list", usercontroller.ListUser)
		userGroup.GET("/logout", usercontroller.UserSignOut)
	}
	smpGroup := r.Group("/smp")
	{
		smpGroup.GET("/list", smpcontroller.ListSmp)
		smpGroup.GET("/select", smpcontroller.SelectSmp)

	}
	expGroup := r.Group("/exp")
	{
		expGroup.GET("/call", experimentcontroller.Callpython)
	}
	return
}
