/*
@Time : 2021/12/28 3:38 下午
@Author : yuyunqing
@File : ConnToRead
@Software: GoLand
*/
package TcpHelp

import (
	"bufio"
	"bytes"
	"encoding/binary"
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

	s, err := decode(reader)
	if err != nil {
		return "", err
	}
	log.Println("读取输入",s)
	return s ,nil

}


//用于解决 粘包问题
func decode(reader *bufio.Reader) (string , error)  {
	len_msg,_ := reader.Peek(4)

	lenbuf := bytes.NewBuffer(len_msg)
	var length int32
	err := binary.Read(lenbuf, binary.LittleEndian, &length)
	if err != nil {
		return "", err
	}

	if int32(reader.Buffered()) < 4+length {
		return "", fmt.Errorf("客户端输入读取异常")
	}
	// 读取真正的消息数据
	pack := make([]byte, int(4+length))
	_, err = reader.Read(pack)
	if err == io.EOF {
		return "", fmt.Errorf("客户端退出")
	}
	if err != nil {
		return "", err
	}
	return string(pack[4:]), nil
}