/*
@Time : 2021/12/29 4:06 下午
@Author : yuyunqing
@File : client
@Software: GoLand
*/
package main

import (
	"bufio"
	"io"
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
		tcp.Write([]byte(send))

		var p [200]byte
		reder := bufio.NewReader(tcp)
		read, err := reder.Read(p[:])
		if err == io.EOF{
			log.Println("服务端退出")
			return
		}
		if err != nil {
			log.Println("读取失败",err)
			return
		}

		log.Println(string(p[:read]))
	}
}
