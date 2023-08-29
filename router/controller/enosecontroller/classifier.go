package enosecontroller

import (
	"EnoseBackend/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ClassifierRequestBody struct {
	Enose_Name      string
	Classifier_Name []string
}
type ClassifierDelRequestBody struct {
	Enose_Name      string
	Classifier_Name string
}

func ListClassifierByEnoseName(c *gin.Context) {
	req := new(ClassifierRequestBody)
	err := c.ShouldBind(&req)
	res := new([]model.Classifier)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "error": err.Error()})
		return
	}
	res, err = model.GetClassifierByEnoseName(req.Enose_Name)
	fmt.Println(res)
	c.ShouldBind(res)
	c.JSON(200, res)
}

func AddClassifier(c *gin.Context) {
	req := new(ClassifierRequestBody)
	c.ShouldBind(&req)
	//sen := new([]model.Classifier)

	fmt.Println(req)
	for _, s := range req.Classifier_Name {
		se, err := model.GetClassifierByName(s, req.Enose_Name)
		fmt.Println(se, s, req.Enose_Name)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"code": 1, "success": false, "message": "分类已存在"})
			return
		}
		//if err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{"code": 1, "error": err.Error()})
		//	return
		//}
		classisier := new(model.Classifier)
		classisier.Enose_Name = req.Enose_Name
		classisier.Classifier_Name = s
		fmt.Println(req)
		err = model.AddClassifier(classisier)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 1, "error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "success": true, "message": "分类添加成功"})

}
func DelClassifier(c *gin.Context) {
	req := new(ClassifierDelRequestBody)
	c.BindJSON(req)
	fmt.Println(req)
	classisier, _ := model.GetClassifierByName(req.Classifier_Name, req.Enose_Name)
	fmt.Println(classisier)
	model.DeleteClassifier(classisier)
	c.JSON(http.StatusOK, gin.H{"code": 0, "success": true, "message": "删除成功"})
}
