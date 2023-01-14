package model

import (
	"EnoseBackend/dao"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Identity string
	Name     string
	Password string
}

func AddUser(user *User) (err error) {
	err = dao.DB.Create(user).Error
	return
}

func UpdateUser(user *User) (err error) {
	err = dao.DB.Save(user).Error
	return
}

func GetUserByName(name string) (user *User, err error) {
	user = new(User)
	err = dao.DB.Debug().Where("name=?", name).First(&user).Error
	if err != nil {
		return nil, err
	}
	return user, err
}

func GetUserById(id uint) (user *User, err error) {
	user = new(User)
	err = dao.DB.Debug().Where("id=?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return
}

func DeleteUser(user *User) {
	dao.DB.Delete(&user)
	return
}
