/*
@Time : 2021/12/29 2:18 下午
@Author : yuyunqing
@File : entrance
@Software: GoLand
*/
package Cmd

import (
	"fmt"
	"go-redis/FileAction"
	"log"
	"strings"
)

//指令返回方法集
type Cmd_fun struct {
	Funcs func(cmd_arr []string)(string , error) //指定方法
	Func_num int //方法对应的参数数量
	Data_type string //操作数据类型
}

//执行tcp指令入口类似     GET  info  // set info 1
func CmdAction(cmd_input string) (string ,error) {
	//根据空格拆分字符串
	cmd_arr := strings.Split(cmd_input," ")
	if len(cmd_arr) == 0  {
		return "" , fmt.Errorf("未获取到指令，请检查命令是否正确: " +cmd_input)
	}
	log.Println(cmd_arr)
	//获取返回
	FileAction.SaveToFile(strings.Join(cmd_arr,"|"))
	actionFunc, err := GetActionFunc(cmd_arr[0])
	if err != nil {
		return "", err
	}

	if actionFunc.Data_type != "all"{
		//比较最少所需参数数量
		check_len := CheckCmdLen(actionFunc.Func_num , cmd_arr)
		if !check_len {
			return "",fmt.Errorf("参数数量不正确 该命令最少要求参数 ：%d",actionFunc.Func_num)
		}
	}

	res, err := actionFunc.Funcs(cmd_arr)
	if err != nil {
		return "", err
	}
	return res ,nil
}


/**
 * @Author yuyunqing
 * @Description //默认指令第一个为指定方法
 * @Date 2:25 下午 2021/12/29
 **/
func GetActionFunc(cmd_type string) ( Cmd_fun,error) {
	funcs := map[string]Cmd_fun{
		"get" : {Get,1 , "RString" },
		"set" : {Set,2 , "RString" },
		"nomatch" : {NoMatch,0 , "all" },
		"info" : {Info,0 , "all" },
	}
	if _,ok := funcs[cmd_type] ;ok {
		return funcs[cmd_type] ,nil
	}else{
		return funcs["nomatch"] ,nil
	}
}


