package router

import (
	"EnoseBackend/router/controller/enosecontroller"
	"EnoseBackend/router/controller/experimentcontroller"
	"EnoseBackend/router/controller/learningmodelcontroller"
	"EnoseBackend/router/controller/pythoncontroller"
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
		userGroup.POST("/modifyPwd", usercontroller.ModifyPwd)
	}
	deviceGroup := r.Group("/device")
	{

		deviceGroup.GET("/list", enosecontroller.ListEnose)
		deviceGroup.POST("/listmodel", learningmodelcontroller.Listlearningmodel)
		deviceGroup.POST("/update", enosecontroller.UpdateEnose)
		deviceGroup.POST("/add", enosecontroller.AddEnose)
		deviceGroup.POST("/addSensor", enosecontroller.AddSensor)
		deviceGroup.POST("/delSensor", enosecontroller.DelSensor)
		deviceGroup.POST("/listSensor", enosecontroller.ListSensorByEnoseName)
		deviceGroup.POST("/addClassifier", enosecontroller.AddClassifier)
		deviceGroup.POST("/delClassifier", enosecontroller.DelClassifier)
		deviceGroup.POST("/listClassifier", enosecontroller.ListClassifierByEnoseName)
		deviceGroup.POST("/delete", enosecontroller.DeleteEnose)
	}
	smpGroup := r.Group("/smp")

	{
		smpGroup.GET("/list", smpcontroller.ListSmp)
		smpGroup.POST("/del", smpcontroller.DelSmp)
		smpGroup.POST("/detail", smpcontroller.Detail)
		smpGroup.GET("/select", smpcontroller.SelectSmp)

	}
	pythonGroup := r.Group("/python")
	{
		pythonGroup.GET("/list", pythoncontroller.ListPy)
		pythonGroup.POST("/del", pythoncontroller.Delpy)
	}
	expGroup := r.Group("/exp")
	{
		expGroup.GET("/list", experimentcontroller.ListExp)
		expGroup.POST("/expdel", experimentcontroller.ExpDel)
		expGroup.POST("/start", experimentcontroller.StartExp)
		expGroup.POST("/call", experimentcontroller.Callpython)
		expGroup.POST("/callnew", experimentcontroller.Callnewpython)
		expGroup.POST("/datacollect", experimentcontroller.DataCollect)
		expGroup.POST("/setexp", experimentcontroller.SetExp)
		expGroup.GET("/saveCsv", utils.SaveCsv)
		expGroup.POST("/xlsxToJson", utils.XlsxToJson)
		expGroup.POST("/FS", experimentcontroller.FS)
		expGroup.POST("/Slice", experimentcontroller.SliceDataset)
		expGroup.POST("/ExpFinish", experimentcontroller.Expfinish)
		expGroup.POST("/Detail", experimentcontroller.ListExpDetail)
	}
	r.POST("/test", experimentcontroller.Callnewpython)
	r.GET("/set", settingcontroller.Set)
	r.GET("/ws", service.Ws)
	return
}
