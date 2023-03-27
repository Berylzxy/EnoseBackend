package model

import "EnoseBackend/dao"

type Sensor struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	Enose_name  string `json:"enose_name"`
	Sensor_name string `json:"sensor_name"`
}

func AddSensor(sensor *Sensor) (err error) {
	err = dao.DB.Create(sensor).Error
	return
}

func UpdateSensor(sensor *Sensor) (err error) {
	err = dao.DB.Save(sensor).Error
	return
}

func GetSensorById(id uint) (sensor *Sensor, err error) {
	sensor = new(Sensor)
	err = dao.DB.Debug().Where("id=?", id).First(sensor).Error
	if err != nil {
		return nil, err
	}
	return
}
func GetSensorBySensorName(name string) (sensor *Sensor, err error) {
	sensor = new(Sensor)
	err = dao.DB.Debug().Where("sensor_name=?", name).First(sensor).Error
	if err != nil {
		return nil, err
	}
	return
}
func GetSensorByEnoseName(name string) (sensor *[]Sensor, err error) {
	sensor = new([]Sensor)
	err = dao.DB.Debug().Where("enose_name=?", name).Find(sensor).Error
	if err != nil {
		return nil, err
	}
	return
}

func DeleteSensor(sensor *Sensor) {
	dao.DB.Delete(&sensor)
	return
}
