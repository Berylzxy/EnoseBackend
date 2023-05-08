package settingcontroller

import (
	"EnoseBackend/dao"
	"EnoseBackend/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

type config struct {
	Path string
}

func Set(c *gin.Context) {

	set := new(model.Setting)
	err := dao.DB.First(&set).Error
	if err != nil {
		return
	}
	fmt.Println(set.Path)
	c.JSON(200, gin.H{"Path": set.Path})
}
