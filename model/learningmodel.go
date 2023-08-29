package model

import (
	"EnoseBackend/dao"
	"database/sql/driver"
	"gorm.io/gorm"
	"strings"
)

type Tag []string

type Learningmodel struct {
	gorm.Model
	Name            string `json:"Name"`
	Experiment_name string
	Enose_name      string
	FE              string `json:"FeatureExtraction"`
	FS              Tag    `json:"FeatureSelected" gorm:"text"`
	Address         string `json:"Address"`
}

func (m *Tag) Scan(val interface{}) error {
	s := val.([]uint8)
	ss := strings.Split(string(s), ",")
	*m = ss
	return nil
}
func (m Tag) Value() (driver.Value, error) {
	str := strings.Join(m, ",")
	return str, nil
}

func AddLearningmodel(learningmodel *Learningmodel) (err error) {
	err = dao.DB.Create(learningmodel).Error
	return
}

func UpdateLearningmodel(learningmodel *Learningmodel) (err error) {
	err = dao.DB.Save(learningmodel).Error
	return
}

func GetLearningmodelByExpName(ExpName string) (learningmodel *[]Learningmodel, err error) {
	learningmodel = new([]Learningmodel)
	err = dao.DB.Debug().Where("experiment_name=?", ExpName).Find(learningmodel).Error
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
