package experimentcontroller

import (
	"EnoseBackend/model"
	"EnoseBackend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"sort"
	"strconv"
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
	Address   string
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
type GetDeviceName struct {
	Devicename string
}
type SetExpInfo struct {
	Devicename string
	Classifier []string
	Sensor     []string
	Learnmodel []string
}
type ExpDelete struct {
	Name string
}
type ListRes struct {
	Id         string
	Name       string
	Devicename string
}
type ListDetail struct {
	ID         int `json:"id" gorm:"primary_key"`
	Name       string
	Enose_Name string
	Classifier string
	Sensor     string `gorm:"type:text"`
	State      string `json:"state"`
	Pump       string
	Hertz      string
	Start_time time.Time `json:"start_time"`
	End_time   time.Time `json:"end_time"`
}
type FSrequest struct {
	DocAddr   string
	Selectcol []int
	SaveAddr  string
}

//	func ExpList(c *gin.Context) {
//		exp, _ := model.ListExperiment()
//		c.JSON(200, gin.H{"res": exp})
//	}
func StartExp(c *gin.Context) {
	req := new(ExpInfo)
	c.BindJSON(req)
	_, err := model.GetExperimentByName(req.Name)
	sort.Strings(req.Sensor)
	//fmt.Println(req)
	if err != nil {
		exp := new(model.Experiment)
		exp.Name = req.Name
		exp.Pump = req.Pump
		exp.Hertz = req.Hertz
		exp.Classifier = req.Class
		exp.Sensor = strings.Join(req.Sensor, ",")
		//fmt.Println(exp.Sensor)
		exp.State = "start"
		exp.Start_time = time.Now()
		exp.Address = req.Address
		device, err := model.GetEnoseByName(req.Device)
		if err != nil {
			c.JSON(200, gin.H{"success": false, "message": "未找到设备"})
			return
		} else {
			exp.Enose_Name = device.Name
		}
		device.State = "在线"
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
		model.UpdateEnose(device)
		fmt.Println(req.Address)
		utils.CreateDateDir(req.Address + "\\" + req.Name)
		c.JSON(200, gin.H{"success": true, "message": "创建成功"})
	} else {
		c.JSON(200, gin.H{"success": false, "message": "已存在实验"})
		return
	}

}
func ExpDel(c *gin.Context) {
	req := new(ExpDelete)
	c.BindJSON(req)
	Exp, _ := model.GetExperimentByName(req.Name)
	model.DeleteExperiment(Exp)
	c.JSON(200, gin.H{"message": "实验删除成功"})
}
func SetExp(c *gin.Context) {
	req := new(GetDeviceName)
	c.BindJSON(req)
	res := new(SetExpInfo)
	res.Devicename = req.Devicename
	//enose, _ := model.GetEnoseByName(res.Devicename)
	sensor, _ := model.GetSensorByEnoseName(req.Devicename)
	learningmodel, _ := model.GetLearningmodelByEnoseName(req.Devicename)
	classifier, _ := model.GetClassifierByEnoseName(req.Devicename)
	for _, val := range *sensor {
		tmp := val.Sensor_name
		res.Sensor = append(res.Sensor, tmp)

	}
	for _, val := range *learningmodel {
		tmp := val.Name
		res.Learnmodel = append(res.Learnmodel, tmp)

	}
	for _, val := range *classifier {
		tmp := val.Classifier_Name
		res.Classifier = append(res.Classifier, tmp)

	}
	fmt.Println(res)
	c.ShouldBind(res)
	c.JSON(200, res)
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

func ListExp(c *gin.Context) {
	Exp, _ := model.ListExperiment()
	res := []ListRes{}
	for i, val := range *Exp {
		tmp := ListRes{}
		tmp.Id = strconv.Itoa(val.ID)
		tmp.Name = val.Name
		tmp.Devicename = val.Name
		res = append(res, tmp)
		fmt.Println(i, res)
	}
	c.ShouldBind(res)
	c.JSON(200, res)
}
func StreamWriterFunc(contents [][]string, A string, B string) {
	//打开工作簿
	file, err := excelize.OpenFile(A)
	if err != nil {
		return
	}
	sheet_name := "Sheet1"
	//获取流式写入器
	streamWriter, _ := file.NewStreamWriter(sheet_name)
	if err != nil {
		fmt.Println(err)
	}

	rows, _ := file.GetRows(sheet_name) //获取行内容
	cols, _ := file.GetCols(sheet_name) //获取列内容
	fmt.Println("行数rows:  ", len(rows), "列数cols:  ", len(cols))

	//将源文件内容先写入excel
	for rowid, row_pre := range rows {
		row_p := make([]interface{}, len(cols))
		for colID_p := 0; colID_p < len(cols); colID_p++ {
			//fmt.Println(row_pre)
			//fmt.Println(colID_p)
			if row_pre == nil {
				row_p[colID_p] = nil
			} else {
				row_p[colID_p] = row_pre[colID_p]
			}
		}
		cell_pre, _ := excelize.CoordinatesToCellName(1, rowid+1)
		if err := streamWriter.SetRow(cell_pre, row_p); err != nil {
			fmt.Println(err)
		}
	}

	//将新加contents写进流式写入器
	for rowID := 0; rowID < len(contents); rowID++ {
		row := make([]interface{}, len(contents[0]))
		for colID := 0; colID < len(contents[0]); colID++ {
			row[colID] = contents[rowID][colID]
		}
		cell, _ := excelize.CoordinatesToCellName(1, rowID+len(rows)+1) //决定写入的位置
		if err := streamWriter.SetRow(cell, row); err != nil {
			fmt.Println(err)
		}
	}

	//结束流式写入过程
	if err := streamWriter.Flush(); err != nil {
		fmt.Println(err)
	}
	//保存工作簿
	if err := file.SaveAs(B); err != nil {
		fmt.Println(err)
	}
}
func convertnumbertocol(Selectcol []int) (s []string) {
	var c rune = 'A'
	for i := range Selectcol {
		fmt.Println(string(rune(i) + c))
		s = append(s, string(rune(i)+c))
	}
	return s
}
func FS(c *gin.Context) {
	req := new(FSrequest)
	c.BindJSON(req)
	fmt.Println(req)
	fe, err := excelize.OpenFile(req.DocAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := fe.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	//rows, err := fe.GetRows("Sheet1")

	rows, err := fe.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	StreamWriterFunc(rows, req.DocAddr, req.SaveAddr)
	fs, err := excelize.OpenFile(req.SaveAddr)
	s := convertnumbertocol(req.Selectcol)
	for i := range s {
		fmt.Println(s[i])
		err := fs.RemoveCol("Sheet1", s[i])
		if err != nil {
			return
		}
	}
	err = fs.Save()
	if err != nil {
		return
	}
	smp, err := model.GetSmpByName(req.SaveAddr)
	if err != nil {
		smp := new(model.Smp)
		smp.Name = req.SaveAddr
		smp.Address = req.SaveAddr
		model.AddSmp(smp)
	} else {
		smp.Address = req.SaveAddr
		model.UpdateSmp(smp)
	}
	c.JSON(200, gin.H{"message": "finish"})
}

//	func ListDetail(c *gin.Context)  {
//		Exp, _ := model.ListExperiment()
//		res := new(model.Experiment)
//		for i, val := range *Exp {
//			tmp := model.Experiment{}
//			tmp.ID = strconv.Itoa(val.ID)
//			tmp.Name = val.Name
//			tmp.Devicename = val.Name
//			res = append(res, tmp)
//			fmt.Println(i, res)
//		}
//		c.ShouldBind(res)
//		c.JSON(200, res)
//	}
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
