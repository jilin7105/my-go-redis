/*
@Time : 2021/12/29 4:06 下午
@Author : yuyunqing
@File : client
@Software: GoLand
*/
package main

import (
	"bufio"
	"go-redis/TcpHelp"
	"log"
	"net"
	"os"
	"strings"
)

func main()  {
	tcp, err := net.Dial("tcp","127.0.0.1:20001")
	if err != nil {
		log.Println("链接失败")
		return
	}
	input := bufio.NewReader(os.Stdin)
	for true {
		msg, err := input.ReadString('\n')
		if err != nil {
			log.Println("读取失败")
			return
		}

		send := strings.Trim(msg, "\r\n")
		//发送命令
		TcpHelp.Write(send, tcp)

		reader := bufio.NewReader(tcp)
		read, err := TcpHelp.Read("127.0.0.1:20001", reader)
		if err != nil {
			return
		}
		log.Println("来自服务端消息127.0.0.1:20001>", read)
	}
}
