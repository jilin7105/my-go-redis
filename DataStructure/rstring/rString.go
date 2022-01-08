/*
@Time : 2021/12/28 2:54 下午
@Author : yuyunqing
@File : rString
@Software: GoLand
*/
package rstring

import (
	"fmt"
	"strconv"
	"strings"
)

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


func  (r RString) GetDataType () string{
	return r.Type
}

func (r RString)  ValueToString() (string) {
	return r.Value
}


func (r RString)  GetExpires() (int) {
	return r.Time
}


func (r RString)  GetType() (int) {
	return r.Time
}

func (r RString)  GetRdbStr() string{
	return fmt.Sprintf("RString;%d,%s",r.Time,r.Value)
}

func (r *RString)  GetStrFromRdb(s string) {
	arr := strings.Split(s,",")
	if len(arr) ==2 {
		r.Time, _ = strconv.Atoi(arr[0])
		r.Value= arr[1]
	}

}

