/*
@Time : 2021/12/28 2:14 下午
@Author : yuyunqing
@File : main
@Software: GoLand
*/
package main

import (
	"go-redis/FileAction"
	"go-redis/MemoryManagement"
	"go-redis/RDB"
	"go-redis/TcpServer"
)

func init()  {
	MemoryManagement.InitData()
	FileAction.InitFile()
	//aof 从aof中初始化数据
	//AOF.ReaderByFile()
	RDB.ReadFromRDb()
}

func main()  {
	TcpServer.Run()
}
