/*
@Time : 2021/12/28 2:14 下午
@Author : yuyunqing
@File : main
@Software: GoLand
*/
package main

import (
	"go-redis/MemoryManagement"
	"go-redis/TcpServer"
)

func init()  {
	MemoryManagement.InitData()
}

func main()  {
	TcpServer.Run()
}
