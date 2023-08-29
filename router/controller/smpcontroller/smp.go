package smpcontroller

import (
	"EnoseBackend/dao"
	"EnoseBackend/model"
	"bufio"
	"encoding/csv"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"os"
	"strings"
)

type Del struct {
	Id int
}
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
func DelSmp(c *gin.Context) {
	req := new(Del)
	c.BindJSON(req)
	smp, _ := model.GetSmpById(req.Id)
	model.DeleteSmp(smp)
	c.JSON(200, gin.H{"code": "0", "success": "1", "message": "success"})
}
func Detail(c *gin.Context) {
	req := new(Del)
	c.BindJSON(req)
	smp, _ := model.GetSmpById(req.Id)
	c.JSON(200, smp)
}
func SelectSmp(c *gin.Context) {
	req := new(SelectSmpRequestBody)
	c.BindJSON(&req)
	var path string
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
	length := len(path)

	filetype := path[length-3 : length]
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	defer f.Close()
	if filetype == "txt" {
		content, err := readTxt(f)
		if err != nil {
			fmt.Println("err:", err)
			return
		}
		fmt.Println("content:", content)
		c.JSON(200, gin.H{"massage": content})
	} else if filetype == "csv" {
		reader := csv.NewReader(f)
		for {
			content, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println("Error:", err)
			}
			c.JSON(200, gin.H{"line": content})
		}
	}

	return
}
func Savetxt(c *gin.Context) {
	req := new(SaveSmpRequestBody)
	c.BindJSON(&req)
	var js J
	label := req.Label
	name := req.Name

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
	dao.DB.Create(&smp)
}
func Savecsv(c *gin.Context) {
	req := new(SaveSmpRequestBody)
	c.BindJSON(&req)
	var js J
	label := req.Label
	name := req.Name

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
	dao.DB.Create(smp)
}
func readTxt(r io.Reader) ([]string, error) {
	reader := bufio.NewReader(r)
	l := make([]string, 0, 64)
	// 按行读取
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

		l = append(l, strings.Trim(string(line), " "))
	}

	return l, nil
}
