package experimentcontroller

import (
	"EnoseBackend/model"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"os/exec"
	"time"
)

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

var Res []byte

type CallPythonRequest struct {
	ExpName   string
	Algorithm string
	Kind      string
	Dataname  string
	Save      string
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
	python := new(model.Pythonfile)
	expstep := new(model.Exp_step)
	data := new(model.Smp)
	data, _ = model.GetSmpByName(req.Dataname)
	python, _ = model.GetPythonfileByName(req.Algorithm)

	cmd := exec.Command("python", python.Address, req.Kind, data.Address)
	if req.Algorithm == "预测" {
		//fmt.Println("进来了")
		learningmodel := new(model.Learningmodel)
		learningmodel, _ = model.GetLearningmodelByName(req.Kind)
		//fmt.Println(learningmodel.Address)
		cmd = exec.Command("python", python.Address, learningmodel.Address, data.Address)
	}
	if req.Algorithm == "训练" {
		learningmodel, err := model.GetLearningmodelByName(req.Kind)
		if err != nil {
			learningmodel := new(model.Learningmodel)
			learningmodel.Address = req.Save
			model.AddLearningmodel(learningmodel)
		} else {
			learningmodel.Address = req.Save
			model.UpdateLearningmodel(learningmodel)
		}

		cmd = exec.Command("python", python.Address, req.Kind, data.Address, req.Save)

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
	c.JSON(200, gin.H{"message": string(Res)})
}
