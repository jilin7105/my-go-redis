/*
@Time : 2021/12/28 3:38 下午
@Author : yuyunqing
@File : ConnToRead
@Software: GoLand
*/
package TcpServer

import (
	"bufio"
	"fmt"
	"io"
	"log"
)


/**
 * @Author yuyunqing
 * @Description //统一tcp协议数据读取入口
 * @Date 3:42 下午 2021/12/28
 **/
func Read(ip string ,reader *bufio.Reader) (string ,error) {

	var p [126]byte
	read, err := reader.Read(p[:])
	if err == io.EOF {
		return "", fmt.Errorf("ip 链接已退出 ",ip)
	}
	if err != nil {

		return "", err
	}
	log.Println("读取输入",string(p[:read]))
	return string(p[:read]) ,nil

}