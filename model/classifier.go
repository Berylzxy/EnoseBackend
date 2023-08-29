package model

import (
	"EnoseBackend/dao"
	_ "github.com/go-sql-driver/mysql"
)

type Classifier struct {
	ID              uint ` gorm:"primary_key"`
	Enose_Name      string
	Classifier_Name string
}

func AddClassifier(classifier *Classifier) (err error) {
	err = dao.DB.Create(classifier).Error
	return
}

func UpdateClassifier(classifier *Classifier) (err error) {
	err = dao.DB.Save(classifier).Error
	return
}

func GetClassifier(id uint) (classifier *Classifier, err error) {
	classifier = new(Classifier)
	err = dao.DB.Debug().Where("id=?", id).First(classifier).Error
	if err != nil {
		return nil, err
	}
	return
}
func GetClassifierByEnoseName(name string) (classifier *[]Classifier, err error) {
	classifier = new([]Classifier)
	err = dao.DB.Debug().Where("enose_name=?", name).Find(classifier).Error
	if err != nil {
		return nil, err
	}
	return
}
func GetClassifierByName(name string, enosename string) (classifier *Classifier, err error) {
	classifier = new(Classifier)
	err = dao.DB.Debug().Where("classifier_name=? AND enose_name=?", name, enosename).First(classifier).Error
	if err != nil {
		return nil, err
	}
	return
}
func DeleteClassifier(classifier *Classifier) {
	dao.DB.Delete(&classifier)
	return
}
func ListClassifier() (classifier *[]Classifier, err error) {
	classifier = new([]Classifier)
	err = dao.DB.Find(&classifier).Error
	if err != nil {
		return nil, err
	}
	return
}
