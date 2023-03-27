package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateDateDir(basePath string) (dirPath, dataString string) {
	fmt.Println(basePath)
	folderName := ""
	folderPath := filepath.Join(basePath, folderName)
	if _, err := os.Stat(folderPath); os.IsNotExist(err) {
		// 必须分成两步
		// 先创建文件夹
		os.Mkdir(folderPath, 0777)
		// 再修改权限
		os.Chmod(folderPath, 0777)
	}
	fmt.Println(folderPath, folderName)
	return folderPath, folderName
}
