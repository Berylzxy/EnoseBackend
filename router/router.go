package router

import (
	"EnoseBackend/router/controller/enosecontroller"
	"EnoseBackend/router/controller/experimentcontroller"
	"EnoseBackend/router/controller/smpcontroller"
	"EnoseBackend/router/controller/usercontroller"
	"EnoseBackend/utils"
	"github.com/gin-gonic/gin"
)

func Init(r *gin.Engine) {

	userGroup := r.Group("/user")
	{
		userGroup.POST("/signIn", usercontroller.UserSignIn)
		userGroup.POST("/signUp", usercontroller.UserSignUp)
		//userGroup.POST("/forget", usercontroller.UserForget)
		userGroup.GET("/info", usercontroller.UserInfo)
		userGroup.GET("/list", usercontroller.ListUser)
		userGroup.GET("/logout", usercontroller.UserSignOut)
	}
	deviceGroup := r.Group("/device")
	{
		deviceGroup.GET("/list", enosecontroller.ListEnose)
	}
	smpGroup := r.Group("/smp")

	{
		smpGroup.GET("/list", smpcontroller.ListSmp)
		smpGroup.GET("/select", smpcontroller.SelectSmp)

	}
	expGroup := r.Group("/exp")
	{
		expGroup.GET("/list", experimentcontroller.ListExp)
		expGroup.POST("/expdel", experimentcontroller.ExpDel)
		expGroup.POST("/start", experimentcontroller.StartExp)
		expGroup.POST("/call", experimentcontroller.Callpython)
		expGroup.POST("/datacollect", experimentcontroller.DataCollect)
		expGroup.POST("/setexp", experimentcontroller.SetExp)
		expGroup.GET("/saveCsv", utils.SaveCsv)
		expGroup.POST("/xlsxToJson", utils.XlsxToJson)
		expGroup.POST("/FS", experimentcontroller.FS)
	}
	return
}
