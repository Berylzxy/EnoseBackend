package pythoncontroller

import (
	"EnoseBackend/dao"
	"EnoseBackend/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Del struct {
	Id int
}

func ListPy(c *gin.Context) {
	var pythonfile []model.Pythonfile
	if err := dao.DB.Find(&pythonfile).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, pythonfile)
	}
}
func Delpy(c *gin.Context) {
	req := new(Del)
	c.BindJSON(req)
	pythonfile, _ := model.GetPythonfileById(req.Id)
	model.DeletePythonfile(pythonfile)
	c.JSON(200, gin.H{"code": "0", "success": "1", "message": "success"})
}
