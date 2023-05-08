package experimentcontroller

import (
	"EnoseBackend/model"
	"EnoseBackend/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

type ExpInfo struct {
	Name       string
	Enose_Name string
	Class      string
	Sensor     []string
	Modelname  []string
	FE         string
	FS         string
	Selected   []string
	Start_time string
	End_time   string
	State      string
	Address    string
	Duration   string
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
type Slicerequest struct {
	DocAddr    string
	Scale      string //划分验证集的比例
	Train      string // 训练集地址
	Validation string // 验证集地址
}
type GetDetails struct {
	Name string
}

func generateRandomNumber(start int, end int, count int) []int {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}

	//存放结果的slice
	nums := make([]int, 0)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start

		//查重
		exist := false
		for _, v := range nums {
			if v == num {
				exist = true
				break
			}
		}
		if !exist {
			nums = append(nums, num)
		}
	}
	sort.Ints(nums)
	return nums
} //生成随机数，用于验证集划分

//	func ExpList(c *gin.Context) {
//		exp, _ := model.ListExperiment()
//		c.JSON(200, gin.H{"res": exp})
//	}
func StartExp(c *gin.Context) {
	req := new(ExpInfo)
	c.BindJSON(req)
	_, err := model.GetExperimentByName(req.Name)
	sort.Strings(req.Sensor)
	fmt.Println(req)
	if err != nil {
		exp := new(model.Experiment)
		exp.Name = req.Name

		exp.Classifier = req.Class
		exp.Sensor = strings.Join(req.Sensor, ",")
		//fmt.Println(exp.Sensor)
		exp.State = "start"
		exp.Start_time = time.Now()

		device, err := model.GetEnoseByName(req.Enose_Name)
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
		expstep.Name = req.Name
		expstep.Step = "setting"
		expstep.Start_Time = time.Now()
		model.AddExperiment(exp)
		model.AddExp_step(expstep)
		model.AddExperiment(exp)
		model.UpdateEnose(device)
		//fmt.Println(req.Address)
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

func ListExp(c *gin.Context) { //  返回ID 实验名 设备名
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
func ListExpDetail(c *gin.Context) { //接收实验名 返回状态 数据地址 以及实验生成的模型
	type Detail struct {
		State       string
		DataAddress string
		Starttime   time.Time
		Model       []model.Learningmodel
	}
	req := new(GetDetails)
	c.BindJSON(req)
	Exp, _ := model.GetExp_stepByName(req.Name)
	res := new(Detail)
	res.DataAddress = Exp.Data_Address
	res.State = Exp.Step
	res.Starttime = Exp.Start_Time
	Model, _ := model.GetLearningmodelByExpName(req.Name)
	res.Model = *Model
	c.ShouldBind(res)
	c.JSON(200, res)
}
func GetExpInfo(c *gin.Context) {
	req := new(ExpDelete)
	c.BindJSON(req)
	fmt.Println(req)
	c.JSON(200, gin.H{})
}

func Expfinish(c *gin.Context) {
	req := new(ExpInfo)
	c.BindJSON(req)
	fmt.Println(req)
	Exp, _ := model.GetExperimentByName(req.Name)
	for i := range req.Modelname {
		learningmodel, _ := model.GetLearningmodelByName(req.Modelname[i], req.Enose_Name, req.Name)
		learningmodel.FE = req.FE
		learningmodel.FS = req.Selected
		//learningmodel.FS = strings.Join(req.Selected, ",")
		fmt.Println("learning", learningmodel)
		model.UpdateLearningmodel(learningmodel)

	}
	fmt.Println("??")

	Exp.State = "finish"
	Exp.End_time = time.Now()
	model.UpdateExperiment(Exp)
	//fmt.Println(Exp.Name, Exp.Enose_Name)
	c.JSON(200, gin.H{"message": "finish"})
}
func StreamWriterFunc(A string, B string) {
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
	k := 0
	for _, i := range Selectcol {
		fmt.Println(i)
		fmt.Println(string(rune(i) + c))
		s = append(s, string(rune(i)+c-rune(k)))
		k = k + 1
	}
	return s
}
func SliceDataset(c *gin.Context) {
	req := new(Slicerequest)
	c.BindJSON(req)
	fmt.Println(req)
	f, err := excelize.OpenFile(req.DocAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	tr := excelize.NewFile()
	defer func() {
		if err := tr.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 创建一个工作表
	_, err = tr.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	va := excelize.NewFile()
	defer func() {
		if err := va.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// 创建一个工作表
	_, err = va.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	cols, _ := f.GetCols("Sheet1") //获取列内容
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	k := 0 //行数
	for i := range rows {
		if i > 500 {
			fmt.Println(i)
		}
		if err != nil {
			return
		}
		k++
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(rows), func(i, j int) {
		rows[i], rows[j] = rows[j], rows[i]
	})
	sli, _ := strconv.ParseFloat(req.Scale, 32)
	point := float64(k) * sli //验证集个数
	val := generateRandomNumber(0, k, int(point))
	//fmt.Println(val)
	//fmt.Println(k, int(point))
	//获取流式写入器
	streamWritertr, _ := tr.NewStreamWriter("Sheet1")
	streamWriterva, _ := va.NewStreamWriter("Sheet1")
	if err != nil {
		fmt.Println(err)
	}
	//

	//cols, _ := file.GetCols(sheet_name)	//获取列内容
	//fmt.Println("行数rows:  ", len(rows),"列数cols:  ", len(cols))
	//
	////将源文件内容先写入excel
	m := 0
	for rowid, row_pre := range rows {
		fmt.Println(rowid)
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
		//fmt.Println(m)
		if rowid != val[m] {
			fmt.Println(rowid, val[m], rowid+1-m)
			cell_pre, _ := excelize.CoordinatesToCellName(1, rowid+1-m)
			//fmt.Println(rowid)
			if err := streamWritertr.SetRow(cell_pre, row_p); err != nil {
				fmt.Println(err)
			}
		}
		if rowid == val[m] {
			cell_pre, _ := excelize.CoordinatesToCellName(1, m+1)
			if err := streamWriterva.SetRow(cell_pre, row_p); err != nil {
				fmt.Println(err)
			}
			if m+1 < int(point) {
				m = m + 1
			}
		}

	}
	//
	////将新加contents写进流式写入器
	//for rowID := 0; rowID < len(contents); rowID++ {
	//	row := make([]interface{}, len(contents[0]))
	//	for colID := 0; colID < len(contents[0]); colID++ {
	//		row[colID] = contents[rowID][colID]
	//	}
	//	cell, _ := excelize.CoordinatesToCellName(1, rowID+len(rows)+1) //决定写入的位置
	//	if err := streamWriter.SetRow(cell, row); err != nil {
	//		fmt.Println(err)
	//	}
	//}
	//
	////结束流式写入过程
	if err := streamWritertr.Flush(); err != nil {
		fmt.Println(err)
	}
	if err := streamWriterva.Flush(); err != nil {
		fmt.Println(err)
	}
	////保存工作簿
	if err := tr.SaveAs(req.Train); err != nil {
		fmt.Println(err)
	}
	if err := va.SaveAs(req.Validation); err != nil {
		fmt.Println(err)
	}
	//保存至数据库
	smptr, err := model.GetSmpByName(req.Train)
	smpva, err := model.GetSmpByName(req.Train)
	if err != nil {
		smpva := new(model.Smp)
		smpva.Name = req.Validation
		smpva.Address = req.Validation
		model.AddSmp(smpva)
		smptr := new(model.Smp)
		smptr.Name = req.Train
		smptr.Address = req.Train
		model.AddSmp(smptr)
	} else {
		smptr.Address = req.Train
		model.UpdateSmp(smptr)
		smpva.Address = req.Validation
		model.UpdateSmp(smpva)
	}
	c.JSON(200, gin.H{"message": "success"})
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

	//rows, err := fe.GetRows("Sheet1")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	StreamWriterFunc(req.DocAddr, req.SaveAddr)
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
