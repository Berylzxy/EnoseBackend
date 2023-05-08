package router

import (
	"EnoseBackend/router/controller/enosecontroller"
	"EnoseBackend/router/controller/experimentcontroller"
	"EnoseBackend/router/controller/learningmodelcontroller"
	"EnoseBackend/router/controller/settingcontroller"
	"EnoseBackend/router/controller/smpcontroller"
	"EnoseBackend/router/controller/usercontroller"
	"EnoseBackend/service"
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
		userGroup.POST("/deal", usercontroller.DealUser)
		userGroup.POST("/delete", usercontroller.DelUser)
	}
	deviceGroup := r.Group("/device")
	{

		deviceGroup.GET("/list", enosecontroller.ListEnose)
		deviceGroup.POST("/listmodel", learningmodelcontroller.Listlearningmodel)
		deviceGroup.POST("/update", enosecontroller.UpdateEnose)
		deviceGroup.POST("/add", enosecontroller.AddEnose)

		deviceGroup.POST("/delete", enosecontroller.DeleteEnose)
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
		expGroup.POST("/Slice", experimentcontroller.SliceDataset)
		expGroup.POST("/ExpFinish", experimentcontroller.Expfinish)
		expGroup.POST("/Detail", experimentcontroller.ListExpDetail)
	}
	r.GET("/set", settingcontroller.Set)
	r.GET("/ws", service.Ws)
	return
}
