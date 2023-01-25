package model

import (
	"EnoseBackend/dao"
	"time"
)

type Exp_step struct {
	ID             int `json:"id" gorm:"primary_key"`
	Name           string
	Step           string    `json:"step"`
	Data_Address   string    `json:"data_address"`
	Result_Address string    `json:"result_address"`
	Start_Time     time.Time `json:"start_time"`
	End_Time       time.Time `json:"end_time"`
}

func ListExp_step() (expstep *[]Exp_step, err error) {
	expstep = new([]Exp_step)
	err = dao.DB.Find(&expstep).Error
	if err != nil {
		return nil, err
	}
	return
}
func AddExp_step(expstep *Exp_step) (err error) {
	err = dao.DB.Create(expstep).Error
	return
}

func UpdateExp_step(expstep *Exp_step) (err error) {
	err = dao.DB.Save(expstep).Error
	return
}

func GetExp_stepById(id uint) (expstep *Exp_step, err error) {
	expstep = new(Exp_step)
	err = dao.DB.Debug().Where("id=?", id).First(expstep).Error
	if err != nil {
		return nil, err
	}
	return
}
func GetExp_stepByName(name string) (expstep *Exp_step, err error) {
	expstep = new(Exp_step)
	err = dao.DB.Debug().Where("name=?", name).First(expstep).Error
	if err != nil {
		return nil, err
	}
	return
}

func DeleteExp_step(expstep *Exp_step) {
	dao.DB.Delete(&expstep)
	return
}
