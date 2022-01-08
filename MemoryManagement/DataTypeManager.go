/*
@Time : 2021/12/29 2:06 下午
@Author : yuyunqing
@File : DataTypeManager
@Software: GoLand
*/
package MemoryManagement

import (
	"fmt"
	"go-redis/DataStructure"
	"go-redis/commonSTR"
)

type BaseDataInMemory struct {
	P map[string]DataStructure.DataBase
}

//创建内存字典
var BIM *BaseDataInMemory

//实现单次请求单例模式，多线程状态下会多次实例化该字段
func GetMainInfo() *BaseDataInMemory {
	if BIM == nil {
		BIM = &BaseDataInMemory{P: map[string]DataStructure.DataBase{}}

	}
	return BIM
}



func InitData() string {
	GetMainInfo()
	//获取基础信息
	return commonSTR.INIT_DATA_SERUCCESS
}


/**
 * @Author yuyunqing
 * @Description //返回基础信息
 * @Date 2:11 下午 2021/12/29
 **/
func Info() (string,error) {
	str := fmt.Sprintf("当前内存含有数据 %d" , len(BIM.P) )
	return str,nil
}

func Set(key string, base DataStructure.DataBase)  {
	BIM = GetMainInfo()
	BIM.P[key] = base
}

func Get(key string) (DataStructure.DataBase,error) {
	BIM = GetMainInfo()
	if _,ok := BIM.P[key] ;ok{
		return BIM.P[key],nil
	}
	return nil ,fmt.Errorf(commonSTR.GET_ERROR)
}