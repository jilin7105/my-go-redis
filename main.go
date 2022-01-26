/*
@Time : 2021/12/28 2:14 下午
@Author : yuyunqing
@File : main
@Software: GoLand
*/
package main

import (
	"go-redis/AOF"
	"go-redis/FileAction"
	"go-redis/MemoryManagement"
	"go-redis/RDB"
	"go-redis/TcpServer"
)

func init()  {
	MemoryManagement.InitData()
	FileAction.InitFile()


	//rdb 和aof 共治模式 ，rdb操作成功后，同时清除aof缓存
	RDB.InitRun()
	AOF.ReaderByFile()

}

func main()  {
	TcpServer.Run()
}
