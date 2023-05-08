package model

import (
	"EnoseBackend/dao"
	"gorm.io/gorm"
)

type Setting struct {
	gorm.Model
	Path string
}

func UpdateSetting(user *Setting) (err error) {
	err = dao.DB.Save(user).Error
	return
}
