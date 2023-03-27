package model

import (
	"EnoseBackend/dao"
	"gorm.io/gorm"
)

type Learningmodel struct {
	gorm.Model
	Name            string `json:"name"`
	Experiment_name string
	Enose_name      string
	Address         string `json:"address"`
}

func AddLearningmodel(learningmodel *Learningmodel) (err error) {
	err = dao.DB.Create(learningmodel).Error
	return
}

func UpdateLearningmodel(learningmodel *Learningmodel) (err error) {
	err = dao.DB.Save(learningmodel).Error
	return
}

func GetLearningmodelByEnose(Enosename string) (learningmodel *[]Learningmodel, err error) {
	learningmodel = new([]Learningmodel)
	err = dao.DB.Debug().Where("enose_name=?", Enosename).Find(learningmodel).Error
	if err != nil {
		return nil, err
	}
	return
}
func GetLearningmodelByName(name string, Enosename string, Experimentname string) (learningmodel *Learningmodel, err error) {
	learningmodel = new(Learningmodel)
	err = dao.DB.Debug().Where("name=? AND enose_name=? And experiment_name=?", name, Enosename, Experimentname).First(learningmodel).Error
	if err != nil {
		return nil, err
	}
	return
}
func GetLearningmodelByEnoseName(name string) (learningmodel *[]Learningmodel, err error) {
	learningmodel = new([]Learningmodel)
	err = dao.DB.Debug().Where("enose_name=?", name).Find(learningmodel).Error
	if err != nil {
		return nil, err
	}
	return
}
func DeleteLearningmodel(learningmodel *Learningmodel) {
	dao.DB.Delete(&learningmodel)
	return
}
