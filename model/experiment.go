package model

import (
	"EnoseBackend/dao"
	"time"
)

type Experiment struct {
	ID         int `json:"id" gorm:"primary_key"`
	Name       string
	Enose_Name string
	Classifier string
	Sensor     string    `gorm:"type:text"`
	Start_time time.Time `json:"start_time"`
	End_time   time.Time `json:"end_time"`
	State      string    `json:"state"`
	Pump       string
	Hertz      string
	Address    string
}

func AddExperiment(experiment *Experiment) (err error) {
	err = dao.DB.Create(experiment).Error
	return
}

func UpdateExperiment(experiment *Experiment) (err error) {
	err = dao.DB.Save(experiment).Error
	return
}

func GetExperimentById(id uint) (experiment *Experiment, err error) {
	experiment = new(Experiment)
	err = dao.DB.Debug().Where("id=?", id).First(experiment).Error
	if err != nil {
		return nil, err
	}
	return
}
func GetExperimentByName(name string) (experiment *Experiment, err error) {
	experiment = new(Experiment)
	err = dao.DB.Debug().Where("name=?", name).First(experiment).Error
	if err != nil {
		return nil, err
	}
	return
}

func DeleteExperiment(experiment *Experiment) {
	dao.DB.Delete(&experiment)
	return
}
func ListExperiment() (experiment *[]Experiment, err error) {
	experiment = new([]Experiment)
	err = dao.DB.Find(&experiment).Error
	if err != nil {
		return nil, err
	}
	return
}
