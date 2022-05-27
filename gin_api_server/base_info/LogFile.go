package base_info

import (
	"log"
	"os"
)

var logger *log.Logger

func init() {
	//file, err := os.Create("test.log")
	file, err := os.OpenFile("test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		log.Fatalln("fail to create test.log file!")
	}
	logger = log.New(file, "", log.Llongfile)
}
func WriteLog(msg string) {

	// 写入文件log格式：/Users/zhou/go/src/zhouTest/test.go:22: 2.Println log with log.LstdFlags ...
	//logger.Println("2.Println log with log.LstdFlags ...")
	//
	logger.SetFlags(log.LstdFlags) // 设置写入文件的log日志的格式
	//
	//// 写入log文件格式： 2018/07/31 17:28:21 4.Println log without log.LstdFlags ...
	//logger.Println("4.Println log without log.LstdFlags ...")
	//
	//fmt.Println("打印")
	//logger.Fatal("9.Fatal log without log.LstdFlags ...")
	//fmt.Println("Fatal终止了程序，这句不执行！")
	logger.Println(msg)
}
