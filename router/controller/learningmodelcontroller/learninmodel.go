package learningmodelcontroller

import (
	"EnoseBackend/model"
	"fmt"
	"github.com/gin-gonic/gin"
)

type devicename struct {
	Devicename string
}
type listLM struct {
	Name       string
	Devicename string
	Expname    string
	FE         string
	FS         []string
}

func Listlearningmodel(c *gin.Context) {
	req := new(devicename)
	c.BindJSON(req)
	fmt.Println(req)
	learningmodel, _ := model.GetLearningmodelByEnoseName(req.Devicename)
	res := []listLM{}
	for i, val := range *learningmodel {
		tmp := listLM{}
		tmp.Name = val.Name
		tmp.Devicename = val.Enose_name
		tmp.Expname = val.Experiment_name
		tmp.FE = val.FE
		tmp.FS = val.FS
		res = append(res, tmp)
		fmt.Println(i, res)
	}
	c.ShouldBind(res)
	c.JSON(200, res)
}
