/*
@Time : 2022/1/8 3:35 下午
@Author : yuyunqing
@File : aof
@Software: GoLand
*/
package AOF

import (
	"go-redis/Cmd"
	"go-redis/FileAction"
	"log"
	"strings"
)

/**
 * @Author yuyunqing
 * @Description //TODO 读取aof内容
 * @Date 3:48 下午 2022/1/8
 **/
func ReaderByFile()  {
	log.Println("aof")
	FileAction.ReadForFile("aof" ,ActionByReaderLine)

}

/**
 * @Author yuyunqing
 * @Description //TODO 逐行处理方法
 * @Date 3:49 下午 2022/1/8
 **/
func ActionByReaderLine(string2 string){
	cmd := strings.Replace(string2 , "|" ," " ,-1)
	Cmd.CmdAction(cmd,false)
}