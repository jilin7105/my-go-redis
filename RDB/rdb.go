/*
@Time : 2022/1/8 4:16 下午
@Author : yuyunqing
@File : rdb
@Software: GoLand
*/
package RDB

import (
	"fmt"
	"go-redis/DataStructure"
	"go-redis/DataStructure/rstring"
	"go-redis/FileAction"
	"go-redis/MemoryManagement"
	"log"
	"strings"
	"time"
)

/**
 * @Author yuyunqing
 * @Description //TODO Rdb初始化功能
 * @Date 5:02 下午 2022/1/26
 **/
func InitRun()  {
	ReadFromRDb()
	CronToRdb()

}


/**
 * @Author yuyunqing
 * @Description //TODO 实现一个定时写入rdb功能
 * @Date 4:58 下午 2022/1/26
 **/
func CronToRdb()  {
	go func() {
		for true {
			time.Sleep(5 * time.Minute)
			SaveToRdb()
		}
	}()

}

/**
 * @Author yuyunqing
 * @Description //TODO 保存当前数据进文件
 * @Date 5:07 下午 2022/1/8
 **/
func SaveToRdb()  {
	BIM := MemoryManagement.GetMainInfo()
	FileAction.Clear("rdb")
	for key, v := range BIM.P {
		str := fmt.Sprintf("%s;%s",key,v.GetRdbStr())
		FileAction.SaveToFile(str,"rdb")
	}
	log.Println("Rdb执行")
}


/**
 * @Author yuyunqing
 * @Description //TODO 从文件中读取
 * @Date 5:07 下午 2022/1/8
 **/
func ReadFromRDb()  {
	log.Println("rdb 数据恢复开始")
	FileAction.ReadForFile("rdb",ActionByReaderLine)
	log.Println("rdb 数据结束开始")
}


/**
 * @Author yuyunqing
 * @Description //TODO 单行rdb数据处理
 * @Date 5:08 下午 2022/1/8
 **/
func ActionByReaderLine(string2 string){
	arr := strings.Split(string2,";")
	data , err := NewData(arr[1])
	if err != nil{
		return
	}
	data.GetStrFromRdb(arr[2])
	MemoryManagement.Set(arr[0],data)
}

//根据第二个值获取数据接头体
func NewData(type_s string) (DataStructure.DataBase,error) {
	list :=map[string]DataStructure.DataBase{
		"RString": rstring.InitData(),
	}

	if _ ,ok := list[type_s] ;ok {
		return list[type_s],nil
	}
	return nil,fmt.Errorf("数据类型错误信息")
}
