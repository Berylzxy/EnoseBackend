package model

import (
	"EnoseBackend/dao"
)

type Enose struct {
	ID         int    `json:"id" gorm:"primary_key"`
	Name       string `json:"name"`
	IP         string
	State      string
	Remark     string `json:"remark"`
	Serial_num int    `json:"serial_num"`
}

func ListEnose() (enose *[]Enose, err error) {
	enose = new([]Enose)
	err = dao.DB.Find(&enose).Error
	if err != nil {
		return nil, err
	}
	return
}
func AddEnose(enose *Enose) (err error) {
	err = dao.DB.Create(enose).Error
	return
}

func UpdateEnose(enose *Enose) (err error) {
	err = dao.DB.Save(enose).Error
	return
}

func GetEnoseById(id uint) (enose *Enose, err error) {
	enose = new(Enose)
	err = dao.DB.Debug().Where("id=?", id).First(enose).Error
	if err != nil {
		return nil, err
	}
	return
}
func GetEnoseByName(name string) (enose *Enose, err error) {
	enose = new(Enose)
	err = dao.DB.Debug().Where("name=?", name).First(enose).Error
	if err != nil {
		return nil, err
	}
	return
}

func DeleteEnose(enose *Enose) {
	dao.DB.Delete(&enose)
	return
}
