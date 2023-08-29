package model

import "EnoseBackend/dao"

type Pythonfile struct {
	ID      int    `json:"id" gorm:"primary_key"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

func AddPythonfile(pythonfile *Pythonfile) (err error) {
	err = dao.DB.Create(pythonfile).Error
	return
}

func UpdatePythonfile(pythonfile *Pythonfile) (err error) {
	err = dao.DB.Save(pythonfile).Error
	return
}

func GetPythonfileById(id int) (pythonfile *Pythonfile, err error) {
	pythonfile = new(Pythonfile)
	err = dao.DB.Debug().Where("id=?", id).First(pythonfile).Error
	if err != nil {
		return nil, err
	}
	return
}
func GetPythonfileByName(name string) (pythonfile *Pythonfile, err error) {
	pythonfile = new(Pythonfile)
	err = dao.DB.Debug().Where("name=?", name).First(pythonfile).Error
	if err != nil {
		return nil, err
	}
	return
}

func DeletePythonfile(pythonfile *Pythonfile) {
	dao.DB.Delete(&pythonfile)
	return
}
