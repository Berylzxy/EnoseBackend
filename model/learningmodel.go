package model

import (
	"EnoseBackend/dao"
	"gorm.io/gorm"
)

type Learningmodel struct {
	gorm.Model
	Name    string `json:"name" gorm:""`
	Address string `json:"address"`
}

func AddLearningmodel(learningmodel *Learningmodel) (err error) {
	err = dao.DB.Create(learningmodel).Error
	return
}

func UpdateLearningmodel(learningmodel *Learningmodel) (err error) {
	err = dao.DB.Save(learningmodel).Error
	return
}

func GetLearningmodelById(id uint) (learningmodel *Learningmodel, err error) {
	learningmodel = new(Learningmodel)
	err = dao.DB.Debug().Where("id=?", id).First(learningmodel).Error
	if err != nil {
		return nil, err
	}
	return
}
func GetLearningmodelByName(name string) (learningmodel *Learningmodel, err error) {
	learningmodel = new(Learningmodel)
	err = dao.DB.Debug().Where("name=?", name).First(learningmodel).Error
	if err != nil {
		return nil, err
	}
	return
}

func DeleteLearningmodel(learningmodel *Learningmodel) {
	dao.DB.Delete(&learningmodel)
	return
}
