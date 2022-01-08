/*
@Time : 2021/12/30 6:09 下午
@Author : yuyunqing
@File : RWfile
@Software: GoLand
*/
package FileAction

import (
	"bufio"
	"fmt"
	"go-redis/commonSTR"
	"io"
	"log"
	"os"
	"strings"
)

type FileConn struct {
	Filelist map[string]*os.File
}

//创建内存字典
var Fc *FileConn

//实现单次请求单例模式，多线程状态下会多次实例化该字段
func GetMainFileInfo() *FileConn {
	if Fc == nil {
		Fc = &FileConn{Filelist: map[string]*os.File{}}
		Fc.Init()
	}
	return Fc
}

func (fc *FileConn)Init()  {
	var filelist = map[string]string{
		"aof": "./aof.log",
		"rdb": "./rdb.log",
	}
	var f *os.File
	var err error
	for key, filename := range filelist {
		if !checkFileIsExist(filename) {
			f, err = os.Create(filename)
			if err != nil {
				fmt.Println(err)
			}
			f.Close()
		}

		f, err = os.OpenFile(filename, os.O_APPEND|os.O_RDWR, os.ModeAppend)
		if err != nil {
			fmt.Println(err)
		}

		if err != nil {
			log.Println("打开文件异常",err.Error())
		}
		fc.Filelist[key] = f
	}

}


func InitFile() string {
	GetMainFileInfo()
	//获取基础信息
	return commonSTR.INIT_FILE_SERUCCESS
}


func SaveToFile(s,type_s string)  {

	fc := GetMainFileInfo()

	if _,ok := fc.Filelist[type_s] ;ok {
		_, err := io.WriteString(fc.Filelist[type_s], s+"\n") //写入文件(字符串)
		if err != nil {
			log.Println("写入文件失败",err.Error())
		}
	}


}

func ReadForFile(types string , callback func(s string))  {
	f_l := GetMainFileInfo()

	if _,ok :=f_l.Filelist[types] ;!ok {
		f_l.Init()
	}


	buf := bufio.NewReader(f_l.Filelist[types])
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)

		if err != nil {
			log.Println(err.Error())
			return
		}
		if len(line) >0 {
			callback(line)
		}

	}

}


func checkFileIsExist(filename string) bool {
	var exist = true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		exist = false
	}
	return exist
}


