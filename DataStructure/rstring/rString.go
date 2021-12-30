/*
@Time : 2021/12/28 2:54 下午
@Author : yuyunqing
@File : rString
@Software: GoLand
*/
package rstring

type RString struct {
	Value string
	Time int
	Type string
}

func InitData() *RString {
	return &RString{
		Type: "RString",
	}
}


func  ( r RString) GetDataType () string{
	return r.Type
}

func ( r RString)  ValueToString() (string) {
	return r.Value
}


func (r RString)  GetExpires() (int) {
	return r.Time
}


func (r RString)  GetType() (int) {
	return r.Time
}


