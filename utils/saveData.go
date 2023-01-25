package utils

import (
	"EnoseBackend/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ExpIdRequestBody struct {
	ID uint
}

type ExpSaveResponseBody struct {
	Addr string
}

func SaveCsv(c *gin.Context) {
	req := new(ExpIdRequestBody)
	err := c.ShouldBind(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 2, "error": err.Error()})
		return
	}
	fmt.Println("req", req)
	fmt.Println("req", req.ID)

	DocAddr, err := model.GetExp_stepById(req.ID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "error": err.Error()})
		return
	}

	/*resp := new(ExpSaveResponseBody)
	resp.Addr = DocAddr.Result_Address
	c.JSON(http.StatusOK, gin.H{"code": 0, "success": true, "data": resp})*/
	//返回地址

	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"实验数据")
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(DocAddr.Result_Address)
}

//model----------------------------------------------------------------
