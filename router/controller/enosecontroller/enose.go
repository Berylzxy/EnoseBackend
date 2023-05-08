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

type update struct {
	Id         string
	Name       string
	Address    string
	Status     string
	Classifier []string
	Sensor     []string
}
type delete struct {
	Id   string
	Name string
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
func AddEnose(c *gin.Context) {
	req := new(update)
	c.BindJSON(req)
	enose, err := model.GetEnoseByName(req.Name)
	if err == nil {

		c.JSON(200, gin.H{"message": "已存在"})
		return
	}
	enose = new(model.Enose)
	enose.ID, _ = strconv.Atoi(req.Id)
	enose.Name = req.Name
	enose.IP = req.Address
	enose.State = req.Status
	enose.Serial_num = len(req.Sensor)
	sensor := req.Sensor
	for i := range sensor {
		s := new(model.Sensor)
		s.Enose_name = req.Name
		s.Sensor_name = sensor[i]
		model.AddSensor(s)
	}
	class := req.Classifier
	for i := range class {
		c := new(model.Classifier)
		c.Enose_Name = req.Name
		c.Classifier_Name = class[i]
		model.AddClassifier(c)
	}
	model.AddEnose(enose)
	c.JSON(200, gin.H{"message": "finished"})
}
func UpdateEnose(c *gin.Context) {
	req := new(update)
	c.BindJSON(req)
	fmt.Println(req)
	enose, _ := model.GetEnoseByName(req.Name)
	enose.ID, _ = strconv.Atoi(req.Id)
	enose.Name = req.Name
	enose.IP = req.Address
	sensor1, _ := model.GetSensorByEnoseName(req.Name)
	for _, elem2 := range req.Sensor {
		found := false
		for _, elem1 := range *sensor1 {
			if elem1.Sensor_name == elem2 {
				found = true
				break
			}
		}
		if !found {
			s := new(model.Sensor)
			s.Enose_name = req.Name
			s.Sensor_name = elem2
			model.AddSensor(s)
		}
	}
	classify1, _ := model.GetClassifierByEnoseName(req.Name)
	for _, elem2 := range req.Classifier {
		found := false
		for _, elem1 := range *classify1 {
			if elem1.Classifier_Name == elem2 {
				found = true
				break
			}
		}
		if !found {
			c := new(model.Classifier)
			c.Enose_Name = req.Name
			c.Classifier_Name = elem2
			model.AddClassifier(c)
		}
	}
	sensorc, _ := model.GetSensorByEnoseName(req.Name)
	enose.Serial_num = len(*sensorc)
	model.UpdateEnose(enose)
	c.JSON(200, gin.H{"message": "success"})
}

func DeleteEnose(c *gin.Context) {
	req := new(delete)
	c.BindJSON(req)
	enose, err := model.GetEnoseByName(req.Name)
	if err != nil {
		c.JSON(200, gin.H{"message": "不存在该设备"})
		return
	}
	model.DeleteEnose(enose)
	sensor, _ := model.GetSensorByEnoseName(req.Name)
	for _, i := range *sensor {
		s, _ := model.GetSensorById(i.ID)
		model.DeleteSensor(s)
	}
	class, _ := model.GetClassifierByEnoseName(req.Name)
	for _, i := range *class {
		c, _ := model.GetClassifier(i.ID)
		model.DeleteClassifier(c)
	}
	c.JSON(200, gin.H{"message": "success"})
}
