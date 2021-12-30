/*
@Time : 2021/12/28 3:00 下午
@Author : yuyunqing
@File : server
@Software: GoLand
*/
package TcpServer

import (
	"bufio"
	"go-redis/Cmd"
	"go-redis/commonSTR"
	"io"
	"log"
	"net"
)

var conn_num = 0
var max_conn_count = 100

/**
 * @Author yuyunqing
 * @Description //处理tcp链接
 * @Date 3:06 下午 2021/12/28
 **/
func process(conn net.Conn)  {
	defer colse(conn)

	ip := conn.RemoteAddr()
	reader := bufio.NewReader(conn)
	//判断当前链接数
	if conn_num >= max_conn_count {
		err := Write(commonSTR.CONN_MORE_THAN_MAX, conn)
		if err != nil {
			log.Println("写数据异常" , err)
			return
		}
		log.Println("超过最大连接数")
		return
	}else{
		conn_num += 1
		log.Println("服务端现有连接数 ",conn_num)
	}

	//保持tcp链接
	for  {

		str, err := Read(ip.String(), reader)
		if err == io.EOF {
			return
		}

		if err != nil {
			log.Println("获取输入异常 " ,err)
			Write(commonSTR.NO_REQ, conn)
			continue
		}

		//处理请求指令逻辑
		res, err := Cmd.CmdAction(str)
		if err != nil {
			log.Println("执行指令异常",err.Error())
			Write(err.Error(), conn)
			continue
		}

		Write(res, conn)
	}
}

/**
 * @Author yuyunqing
 * @Description //关闭链接
 * @Date 3:07 下午 2021/12/28
 **/
func colse(conn net.Conn)  {

	log.Println("关闭tcp链接 ",conn.RemoteAddr().String())
	conn.Close()

	conn_num -= 1
}


func Run ()  {

	listen, err := net.Listen("tcp", "127.0.0.1:20001")
	if err != nil {
		log.Println("监听端口失败 127.0.0.1:20001 ", err.Error())
		return
	}

	for  {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(" 获取conn异常 ：", err.Error())
			return
		}
		log.Println("tcp链接" ,conn.RemoteAddr())
		go process(conn)
	}
}