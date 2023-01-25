package model

import "EnoseBackend/dao"

type Smp struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Name        string `json:"name"`
	Label       string `json:"label"`
	Address     string `json:"address"`
	Duration    string
	Size        string
	Composition string
}

func AddSmp(smp *Smp) (err error) {
	err = dao.DB.Create(smp).Error
	return
}

func UpdateSmp(smp *Smp) (err error) {
	err = dao.DB.Save(smp).Error
	return
}

func GetSmpById(id uint) (smp *Smp, err error) {
	smp = new(Smp)
	err = dao.DB.Debug().Where("id=?", id).First(smp).Error
	if err != nil {
		return nil, err
	}
	return
}
func GetSmpByLabel(label string) (smp *Smp, err error) {
	smp = new(Smp)
	err = dao.DB.Debug().Where("label=?", label).First(smp).Error
	if err != nil {
		return nil, err
	}
	return
}
func GetSmpByName(name string) (smp *Smp, err error) {
	smp = new(Smp)
	err = dao.DB.Debug().Where("name=?", name).First(smp).Error
	if err != nil {
		return nil, err
	}
	return
}
func GetSmpByAddress(address string) (smp *Smp, err error) {
	smp = new(Smp)
	err = dao.DB.Debug().Where("address=?", address).First(smp).Error
	if err != nil {
		return nil, err
	}
	return
}

func DeleteSmp(smp *Smp) {
	dao.DB.Delete(&smp)
	return
}
