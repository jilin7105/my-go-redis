/*
@Time : 2021/12/28 3:17 下午
@Author : yuyunqing
@File : ConnToWrite
@Software: GoLand
*/
package TcpHelp

import (
	"bytes"
	"encoding/binary"
	"log"
	"net"
)



func Write(s string , conn net.Conn) (err error) {
	log.Println(s)
	req, err := encode(s)
	if err != nil {
		log.Println("发送信息打包失败, encode" ,err.Error())
		return err
	}

	_, err = conn.Write(req)
	return
}

//用于解决 粘包问题
func encode(msg string) ([]byte , error)  {
	len_msg :=int32(len(msg))
	pkg := new(bytes.Buffer)

	err := binary.Write(pkg, binary.LittleEndian, len_msg)
	if err != nil {
		return nil, err
	}

	err = binary.Write(pkg, binary.LittleEndian, []byte(msg))
	if err != nil {
		return nil, err
	}

	return pkg.Bytes(),nil
}