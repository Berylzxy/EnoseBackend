package smpcontroller

import (
	"EnoseBackend/dao"
	"EnoseBackend/model"
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
)

type SelectSmpRequestBody struct {
	Name    string
	Label   string
	Address string
}
type SaveSmpRequestBody struct {
	Name    string
	Label   string
	Folder  string
	Address string
}

type J struct {
	Message string `form:"message" json:"message" binding:"required"`
}

func ListSmp(c *gin.Context) {
	var smp []model.Smp
	if err := dao.DB.Find(&smp).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, smp)
	}
}
func SelectSmp(c *gin.Context) {
	req := new(SelectSmpRequestBody)
	c.BindJSON(&req)
	if len(req.Name) != 0 {
		res, _ := model.GetSmpByName(req.Name)
		c.JSON(200, gin.H{"massage": res})
		return
	} else if len(req.Label) != 0 {
		res, _ := model.GetSmpByLabel(req.Label)
		c.JSON(200, gin.H{"massage": res})
		return
	} else if len(req.Address) != 0 {
		res, _ := model.GetSmpByAddress(req.Address)
		c.JSON(200, gin.H{"massage": res})
		return
	}
	c.JSON(200, gin.H{"massage": "字段为空"})
	return
}
func Savetxt(c *gin.Context) {
	req := new(SaveSmpRequestBody)
	c.BindJSON(&req)
	var js J
	label := req.Label
	name := req.Name
	folder := req.Folder
	address := req.Address
	var smp model.Smp

	if err := c.ShouldBind(&js); err == nil {
		a := js.Message
		count := 0
		f, _ := os.Create(address)
		defer f.Close()
		f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM

		for {
			start := strings.Index(a, "[")
			end := strings.Index(a, "]")
			if end > start {
				part := a[start : end+1]
				a = a[end+2:]
				part = strings.Replace(part, "\\", "", -1)
				part = strings.Replace(part, "r", "", -1)
				part = strings.Replace(part, "n", "", -1)
				part = strings.Replace(part, "[", "", -1)
				part = strings.Replace(part, "]", "", -1)
				io.Copy(f, strings.NewReader(part))
				fmt.Println(count)
				count += 1
			} else {
				break
			}
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	smp.Label = label
	smp.Address = address
	smp.Name = name
	smp.Folder = folder
	dao.DB.Create(&smp)
}
func Savecsv(c *gin.Context) {
	req := new(SaveSmpRequestBody)
	c.BindJSON(&req)
	var js J
	label := req.Label
	name := req.Name
	folder := req.Folder
	address := req.Address

	var smp model.Smp
	if err := c.ShouldBind(&js); err == nil {
		a := js.Message
		count := 0
		f, _ := os.Create(address)
		defer f.Close()
		f.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
		w := csv.NewWriter(f)
		for {
			start := strings.Index(a, "[")
			end := strings.Index(a, "]")
			if end > start {
				part := a[start : end+1]
				a = a[end+2:]
				part = strings.Replace(part, "\\", "", -1)
				part = strings.Replace(part, "r", "", -1)
				part = strings.Replace(part, "n", "", -1)
				part = strings.Replace(part, "[", "", -1)
				part = strings.Replace(part, "]", "", -1)
				parts := strings.Split(part, " ")
				w.Write(parts)
				w.Flush()
				fmt.Println(count)
				count += 1
			} else {
				break
			}
		}

	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	smp.Label = label
	smp.Address = address
	smp.Name = name
	smp.Folder = folder
	dao.DB.Create(&smp)
}
