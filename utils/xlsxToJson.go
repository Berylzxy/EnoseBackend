package utils

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

type rowdata []interface{}

type meta struct {
	id string
}

type XlsxToJsonReqBody struct {
	DocAddr string
}

func XlsxToJson(c *gin.Context) {
	req := new(XlsxToJsonReqBody)
	err := c.BindJSON(&req)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 1, "error": err.Error()})
		return
	}
	fmt.Println("req", req)
	fmt.Println("req", req.DocAddr)

	//解析Xlsx 注：源文件内只有一张表Sheet1----------------------------------------------------------------
	fmt.Println("\n\n\n\n", req.DocAddr)
	xlsx, err := excelize.OpenFile(req.DocAddr)
	//xlsx, err := excelize.OpenFile("D:\\桌面\\电子鼻\\data\\钙果（七成）\\1.xlsx")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 2, "error": err.Error()})
		panic(err.Error())
	}

	rows, err := xlsx.GetRows("sheet1")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 3, "error": err.Error()})
		panic(err.Error())
	}
	/*if len(rows) < 13 {
		c.JSON(http.StatusBadRequest, gin.H{"code": 4, "error": err.Error()})
		return
	}*/

	colNum := len(rows[0])
	fmt.Println("colNum:", colNum)

	dataList := make([]rowdata, 0, len(rows))
	for _, row := range rows {
		data := make(rowdata, colNum)
		for k := 0; k < colNum; k++ {
			if k < len(row) {
				data[k] = row[k]
			}
		}
		dataList = append(dataList, data)
	}

	metaList := make([]meta, 12, colNum)
	metaList[0].id = "传感器1"
	metaList[1].id = "传感器2"
	metaList[2].id = "传感器3"
	metaList[3].id = "传感器4"
	metaList[4].id = "传感器5"
	metaList[5].id = "传感器6"
	metaList[6].id = "传感器7"
	metaList[7].id = "传感器8"
	metaList[8].id = "传感器9"
	metaList[9].id = "传感器10"
	metaList[10].id = "传感器11"
	metaList[11].id = "传感器12"

	xlsx.Close()

	//生成json字符串（时间=行数方案）注：该方案源xlsx文件有第12列作为时间------------------------------------------------------

	ret := "["
	for rowCount, row := range dataList {
		//ret += "{"
		//fmt.Println(rowCount)
		//ret += fmt.Sprintf("{%cLine%d%c:[", '"', rowCount+1, '"')
		for idx, meta := range metaList {
			//ret += fmt.Sprintf("%c%s%c:", '"', meta.id, '"')
			//fmt.Println(idx)
			//if row[11] == nil || row[11] == "" {
			//ret += fmt.Sprintf("{%c时间%c:%c%s%c,", '"', '"', '"', "NULL", '"')
			//} else {
			ret += fmt.Sprintf("{%ctime%c:%c%d%c,", '"', '"', '"', rowCount+1, '"')
			//}
			ret += fmt.Sprintf("%csensor%c:%c%s%c,", '"', '"', '"', meta.id, '"')
			ret += fmt.Sprintf("%cvalue%c:", '"', '"')
			if row[idx] == nil || row[idx] == "" {
				ret += "0}"
			} else {
				ret += fmt.Sprintf("%s}", row[idx])
			}
			ret += ","
		}
		ret = ret[:len(ret)-1]
		//ret += "]},"
		if colNum == 1 {
			ret += "}"
		} else {
			ret += ","
		}
	}
	ret = ret[:len(ret)-1]
	ret += "]"

	//生成json字符串（行列数另列方案）注：该方案源xlsx文件须有第13列作为时间,时间是纯文本不是xlsx的时间格式
	/*	ret := "["
		for rowCount, row := range dataList {
			//ret += "{"
			//fmt.Println(rowCount)
			//ret += fmt.Sprintf("{%cLine%d%c:[", '"', rowCount+1, '"')
			for idx, meta := range metaList {
				//ret += fmt.Sprintf("%c%s%c:", '"', meta.id, '"')
				//fmt.Println(idx)
				ret += fmt.Sprintf("{%c行%c:%c%d%c,", '"', '"', '"', rowCount+1, '"')
				ret += fmt.Sprintf("%c列%c:%c%d%c,", '"', '"', '"', idx+1, '"')
				if row[12] == nil || row[12] == "" {
					ret += fmt.Sprintf("%c时间%c:%c%s%c,", '"', '"', '"', "NULL", '"')
				} else {
					ret += fmt.Sprintf("%c时间%c:%c%s%c,", '"', '"', '"', row[12], '"')
				}
				ret += fmt.Sprintf("%c传感器%c:%c%s%c,", '"', '"', '"', meta.id, '"')
				ret += fmt.Sprintf("%cvalue%c:", '"', '"')
				if row[idx] == nil || row[idx] == "" {
					ret += "0}"
				} else {
					ret += fmt.Sprintf("%s}", row[idx])
				}
				ret += ","
			}
			ret = ret[:len(ret)-1]
			//ret += "]},"
			if colNum == 1 {
				ret += "}"
			} else {
				ret += ","
			}
		}
		ret = ret[:len(ret)-1]
		ret += "]"
	*/
	//fmt.Println(ret)

	//输出文件并返回地址
	/*fileSource := filepath.Dir(req.DocAddr)
	fileName := filepath.Base(req.DocAddr)
	suf := filepath.Ext(req.DocAddr)
	jsonName := fileName[0 : (len(fileName))-(len(suf))]
	f, err := os.OpenFile(fileSource+"/"+jsonName+".json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 5, "error": err.Error()})
		panic(err.Error())
	}
	defer f.Close()

	_, err = f.WriteString(ret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 6, "error": err.Error()})
		panic(err.Error())
	}
	//存文件
	resp := new(XlsxToJsonReqBody)
	resp.DocAddr = fileSource + "/" + jsonName + ".json"*/

	type Response struct {
		Code    int
		Data    interface{}
		Success bool
	}

	bytes := json.RawMessage(ret)
	resp := &Response{
		Code:    0,
		Data:    bytes,
		Success: true,
	}
	c.JSON(http.StatusOK, resp)
	fmt.Println(resp)
}
