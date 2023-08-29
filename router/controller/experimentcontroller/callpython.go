package experimentcontroller

import (
	"EnoseBackend/model"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"os/exec"
	"strings"
	"time"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

var Res []byte

/* 这里将算法参数放进一个数组中  参数变为 py名 参数数组（大模型的各项参数）输入数据地址 输出数据地址
 */
type CallnewPythonRequest struct {
	ExpName    string
	DeviceName string
	Pyname     string
	Args       []string
	Data       string
	Res        string
}

func Callnewpython(c *gin.Context) { //data是文件夹的名 文件夹下有很多文件夹 每个文件夹的名子对应标签
	req := new(CallnewPythonRequest)
	c.BindJSON(req)
	fmt.Println(req)
	python := new(model.Pythonfile)

	python, _ = model.GetPythonfileByName(req.Pyname)
	args := req.Args
	for i := range args {
		args[i] = strings.TrimSpace(args[i])
	}
	a := append([]string{python.Address}, args...)
	smp, _ := model.GetSmpByName(req.Data)
	
	a = append(a, smp.Address, req.Res)
	fmt.Println(a)
	//cmd := exec.Command("python", "D:\\桌面\\电子鼻\\pythonfile\\Predict.py", "D:\\桌面\\实验路径\\next\\训练-gaussnb.model", "D:\\桌面\\预测路径\\predict.xlsx", "D:\\桌面\\预测路径\\test-next-gaussnb-predict.xlsx")
	cmd := exec.Command("python", a...)

	//创建获取命令输出管道
	fmt.Println(cmd)
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return
	}

	//执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return
	}

	//读取所有输出
	bytes, err := ioutil.ReadAll(stdout)

	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
		return
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return
	}
	fmt.Printf("stdout:\n\n %s", bytes)

	Res = []byte(ConvertByte2String(bytes, GB18030))
	fmt.Println(Res)

	c.JSON(200, gin.H{"message": string(Res)})
}

type CallPythonRequest struct {
	ExpName    string
	DeviceName string
	Algorithm  string
	Kind       string
	Dataname   string
	Ressave    string
	Selected   []int
	Save       string
}

func ConvertByte2String(byte []byte, charset Charset) string {

	var str string
	switch charset {
	case GB18030:
		decodeBytes, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}

func Callpython(c *gin.Context) { //data是文件夹的名 文件夹下有很多文件夹 每个文件夹的名子对应标签
	req := new(CallPythonRequest)
	c.BindJSON(req)
	fmt.Println(req)
	python := new(model.Pythonfile)
	expstep := new(model.Exp_step)
	data := new(model.Smp)
	data, _ = model.GetSmpByName(req.Dataname)
	python, _ = model.GetPythonfileByName(req.Algorithm)
	cmd := exec.Command("python", python.Address, req.Kind, data.Address, req.Ressave)
	if req.Kind == "" {
		c.JSON(200, gin.H{"message": "请选择函数"})
	}
	if req.Dataname == "" {
		c.JSON(200, gin.H{"message": "请选择文件"})
	}
	if req.Algorithm == "预测" || req.Algorithm == "验证" {
		//fmt.Println("进来了")
		learningmodel := new(model.Learningmodel)
		learningmodel, _ = model.GetLearningmodelByName(req.Kind, req.DeviceName, req.ExpName)
		//fmt.Println(learningmodel.Address)
		cmd = exec.Command("python", python.Address, learningmodel.Address, data.Address, req.Ressave)
	}
	if req.Algorithm == "训练" {
		learningmodel, err := model.GetLearningmodelByName(req.Kind, req.DeviceName, req.ExpName)
		if err != nil {
			learningmodel := new(model.Learningmodel)
			learningmodel.Name = req.Kind
			learningmodel.Enose_name = req.DeviceName
			learningmodel.Experiment_name = req.ExpName
			learningmodel.Address = req.Save
			model.AddLearningmodel(learningmodel)
		} else {
			learningmodel.Address = req.Save
			model.UpdateLearningmodel(learningmodel)
		}

		cmd = exec.Command("python", python.Address, req.Kind, data.Address, req.Ressave, req.Save)

	}

	//创建获取命令输出管道
	fmt.Println(cmd.String())
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Printf("Error:can not obtain stdout pipe for command:%s\n", err)
		return
	}

	//执行命令
	if err := cmd.Start(); err != nil {
		fmt.Println("Error:The command is err,", err)
		return
	}

	//for {
	//	tmp := make([]byte, 1024)
	//	_, err := stdout.Read(tmp)
	//	fmt.Print(string(tmp))
	//	if err != nil {
	//		break
	//	}
	//}
	//读取所有输出
	bytes, err := ioutil.ReadAll(stdout)

	if err != nil {
		fmt.Println("ReadAll Stdout:", err.Error())
		return
	}

	if err := cmd.Wait(); err != nil {
		fmt.Println("wait:", err.Error())
		return
	}
	fmt.Printf("stdout:\n\n %s", bytes)
	expstep.Data_Address = data.Address
	expstep.Step = python.Name
	expstep.Start_Time = time.Now()
	expstep.Name = req.ExpName
	model.AddExp_step(expstep)
	Res = []byte(ConvertByte2String(bytes, GB18030))

	smp, err := model.GetSmpByName(req.Ressave)
	if err != nil {
		smp := new(model.Smp)
		smp.Name = req.Ressave
		smp.Address = req.Ressave
		model.AddSmp(smp)
	} else {
		smp.Address = req.Ressave
		model.UpdateSmp(smp)
	}
	fmt.Println(Res)
	c.JSON(200, gin.H{"message": string(Res)})
}
