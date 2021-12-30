/*
@Time : 2021/12/29 2:37 下午
@Author : yuyunqing
@File : Funcs
@Software: GoLand
*/
package Cmd

import (
	"fmt"
	"go-redis/DataStructure/RString"
	"go-redis/MemoryManagement"
	"log"
	"strconv"
)

type DataAction struct {
	CmdArr []string
	DataType string
}

//获取rstring 类型数据
func Get(cmd_arr []string)  (string , error) {
	get, err := MemoryManagement.Get(cmd_arr[1])
	if err != nil {
		return "", err
	}
	log.Println("查询结果",get)
	if get.GetDataType() != "RString" {
		return "",fmt.Errorf("类型不正确")
	}
	return get.ValueToString() ,nil

}


//写入rstring类型数据
func Set(cmd_arr []string)  (string , error) {

	d :=rstring.InitData()

	d.Value = cmd_arr[2]

	if len(cmd_arr) >= 4 {
		atoi, err := strconv.Atoi(cmd_arr[3])
		if err != nil {
			return "", fmt.Errorf("参数转换int 失败 " +cmd_arr[3])
		}
		d.Time = atoi
	}
	MemoryManagement.Set(cmd_arr[1],d)
	log.Println(cmd_arr,d)
	return "ok" ,nil
}


func NoMatch(cmd_arr []string) (string , error) {
	return "",fmt.Errorf("没有找到该命令NoMatch" + cmd_arr[0])
}


//查询当前状态
func Info(cmd_arr []string) (string , error) {
	info, err := MemoryManagement.Info()
	if err != nil {
		return "", err
	}
	return info ,nil
}


/**
 * @Author yuyunqing
 * @Description //简单验证所需参数
 * @Date 2:50 下午 2021/12/29
 **/
func CheckCmdLen(n int,cmd_arr []string ) bool {
	return n <= len(cmd_arr) -1
}