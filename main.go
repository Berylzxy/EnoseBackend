package main

import (
	"EnoseBackend/dao"
	"EnoseBackend/model"
	"EnoseBackend/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	err := dao.InitMySQL()
	if err != nil {
		return
	}
	err = dao.DB.AutoMigrate(&model.User{}, &model.Sensor{}, model.Learningmodel{}, model.Smp{},
		&model.Pythonfile{}, &model.Enose{}, &model.Classifier{}, &model.Exp_step{}, model.Experiment{})
	if err != nil {
		panic("failed to migrate database")
	}
	r := gin.Default()
	r.Use(cors.Default())
	store := cookie.NewStore([]byte("foobar"))
	r.Use(sessions.Sessions("session", store))
	router.Init(r)
	err = r.Run(":8081")
	if err != nil {
		panic("failed to run gin")
	}
}
