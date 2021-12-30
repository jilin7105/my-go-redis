/*
@Time : 2021/12/30 6:09 下午
@Author : yuyunqing
@File : RWfile
@Software: GoLand
*/
package FileAction

import (
	"io/ioutil"
	"os"
)

func SaveToFile(s string)  {
	//数据压缩
	ioutil.WriteFile("./aof.log", DoZlibCompress([]byte(s)), os.ModeAppend)
}

func ReadForFile()  {
	
}