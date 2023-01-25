package enosecontroller

import (
	"EnoseBackend/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type ListRes struct {
	Id      string
	Name    string
	Address string
	Status  string
}

func ListEnose(c *gin.Context) {
	Enose, _ := model.ListEnose()
	res := []ListRes{}
	for i, val := range *Enose {
		tmp := ListRes{}
		tmp.Id = strconv.Itoa(val.ID)
		tmp.Status = val.State
		tmp.Name = val.Name
		tmp.Address = val.IP
		res = append(res, tmp)
		fmt.Println(i, res)
	}
	c.ShouldBind(res)
	c.JSON(200, res)
}
