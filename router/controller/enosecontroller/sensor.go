package enosecontroller

import (
	"EnoseBackend/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// type ListRes struct {
// 	ID string
// 	Enose_name    string
// 	Sensor_name string
// }

type SensorRequestBody struct {
	Enose_name  string
	Sensor_name []string
}
type SensorDelRequestBody struct {
	Enose_name  string
	Sensor_name string
}

// func ListSensor(c *gin.Context) {
// 	Sensor, _ := model.ListSensor()
// 	res := []ListRes{}
// 	for i, val := range *Sensor {
// 		tmp := ListRes{}
// 		tmp.ID = strconv.Itoa(val.ID)
// 		tmp.Enose_name = val.Enose_name
// 		tmp.Sensor_name = val.Sensor_name
// 		res = append(res, tmp)
// 		fmt.Println(i, res)
// 	}
// 	c.ShouldBind(res)
// 	c.JSON(200, res)
// }

func ListSensorByEnoseName(c *gin.Context) {
	req := new(SensorRequestBody)
	err := c.ShouldBind(&req)
	res := new([]model.Sensor)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "error": err.Error()})
		return
	}
	res, err = model.GetSensorByEnoseName(req.Enose_name)
	fmt.Println(res)
	c.ShouldBind(res)
	c.JSON(200, res)
}

func AddSensor(c *gin.Context) {
	req := new(SensorRequestBody)
	c.ShouldBind(&req)
	//sen := new([]model.Sensor)

	fmt.Println(req)
	for _, s := range req.Sensor_name {
		se, err := model.GetSensorBySensorName(s, req.Enose_name)
		fmt.Println(se, s, req.Enose_name)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{"code": 1, "success": false, "message": "传感器已存在"})
			return
		}
		//if err != nil {
		//	c.JSON(http.StatusBadRequest, gin.H{"code": 1, "error": err.Error()})
		//	return
		//}
		sensor := new(model.Sensor)
		sensor.Enose_name = req.Enose_name
		sensor.Sensor_name = s
		fmt.Println(req)
		err = model.AddSensor(sensor)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{"code": 1, "error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"code": 0, "success": true, "message": "传感器添加成功"})

}
func DelSensor(c *gin.Context) {
	req := new(SensorDelRequestBody)
	c.BindJSON(req)
	sensor, _ := model.GetSensorBySensorName(req.Sensor_name, req.Enose_name)
	model.DeleteSensor(sensor)
	c.JSON(http.StatusOK, gin.H{"code": 0, "success": true, "message": "删除成功"})
}
