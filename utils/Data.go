package utils

//
//import (
//	"fmt"
//	"log"
//	"time"
//)
//
//func a() {
//	//设置串口编号
//	c := &serial.Config{Name: "COM7", Baud: 115200}
//
//	//打开串口
//	s, err := serial.OpenPort(c)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	n, err := s.Write([]byte(string("Hello WPx\r\n")))
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Printf("Write %d Bytes\r\n", n)
//	//延时1000
//	time.Sleep(1000)
//	buf := make([]byte, 128)
//	n, err = s.Read(buf)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("Read %d Bytes\r\n", n)
//	for i := 0; i < n; i++ {
//		fmt.Printf("buf[%d]=%c\r\n", i, buf[i])
//	}
//}
