/*
@Time : 2021/12/28 3:17 下午
@Author : yuyunqing
@File : ConnToWrite
@Software: GoLand
*/
package TcpServer

import (
	"log"
	"net"
)



func Write(s string , conn net.Conn) (err error) {
	log.Println("服务端发送信息" ,s)
	_, err = conn.Write([]byte(s))
	return
}