/*
@Time : 2021/12/28 2:42 下午
@Author : yuyunqing
@File : ataBase
@Software: GoLand
*/
package DataStructure

type DataBase interface {
	//数据内容转字符串
	ValueToString() (string)
	//获取剩余时间
	GetExpires() (int)
	//返回一个结构体
	GetDataType() string

	GetRdbStr() string

	GetStrFromRdb(string)
}

