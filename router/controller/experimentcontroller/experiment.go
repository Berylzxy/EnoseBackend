package experimentcontroller

import (
	"EnoseBackend/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"sort"
	"strings"
	"time"
)

type ExpInfo struct {
	Name      string
	Device    string
	Class     string
	Sensor    []string
	Pump      string
	Hertz     string
	Duration  string
	SmpNumber string
	SmpStruct string
}
type SmpProgRequestBody struct {
	Filename string
}
type SmpProgResponseBody struct {
	Address string
}
type SaveRequestBody struct {
	Name    string
	Label   string
	Address string
}

func ExpList(c *gin.Context) {
	exp, _ := model.ListExperiment()
	c.JSON(200, gin.H{"res": exp})
}
func StartExp(c *gin.Context) {
	req := new(ExpInfo)
	c.BindJSON(req)
	_, err := model.GetExperimentByName(req.Name)
	sort.Strings(req.Sensor)
	s := "传感器"
	for i := 0; i < len(req.Sensor); i++ {
		req.Sensor[i] = s + req.Sensor[i]
	}
	fmt.Println(req)
	if err != nil {
		exp := new(model.Experiment)
		exp.Name = req.Name
		exp.Pump = req.Pump
		exp.Hertz = req.Hertz
		exp.Sensor = strings.Join(req.Sensor, ",")
		fmt.Println(exp.Sensor)
		exp.State = "start"
		exp.Start_time = time.Now()
		device, err := model.GetEnoseByName(req.Device)
		if err != nil {
			c.JSON(200, gin.H{"success": false, "message": "未找到设备"})
			return
		} else {
			exp.Enose_Name = device.Name
		}
		smp := new(model.Smp)
		expstep := new(model.Exp_step)
		smp.Label = req.Class
		smp.Duration = req.Duration
		smp.Size = req.SmpNumber
		smp.Composition = req.SmpStruct
		expstep.Name = req.Name
		expstep.Step = "setting"
		expstep.Start_Time = time.Now()
		model.AddExperiment(exp)
		model.AddExp_step(expstep)
		model.AddExperiment(exp)
		c.JSON(200, gin.H{"success": true, "message": "创建成功"})
	} else {
		c.JSON(200, gin.H{"success": false, "message": "已存在实验"})
		return
	}

}
func DataCollect(c *gin.Context) {
	req := new(SmpProgRequestBody)
	c.BindJSON(req)
	smp, err := model.GetSmpByName(req.Filename)
	fmt.Println(req)
	if err != nil {
		//fmt.Println("111")
		c.JSON(200, gin.H{"message": "未找到数据"})
		return
	} else {
		res := new(SmpProgResponseBody)
		c.BindJSON(res)
		res.Address = smp.Address
		c.JSON(200, gin.H{"message": res.Address})
	}

}
func SaveRes(c *gin.Context) {
	req := new(SaveRequestBody)
	c.BindJSON(req)
	_, err := model.GetSmpByName(req.Name)
	if err != nil {
		smp := new(model.Smp)
		smp.Label = req.Label
		smp.Address = req.Address
		smp.Name = req.Name
		model.AddSmp(smp)
		c.JSON(200, gin.H{"message": "保存成功"})
	} else {
		c.JSON(200, gin.H{"message": "文件名已存在"})
		return
	}
}
