/*
@Time : 2022/1/21 6:34 下午
@Author : yuyunqing
@File : ziplist
@Software: GoLand
*/
package ZIplist

import (
	"bytes"
	"encoding/binary"
	"log"
)



const (
	INIT_ZLBYTES = 13
	INIT_ZLTAIL = 0
	INIT_ZLLEN = 0
	INIT_ZLEND = 255

	ZLBYTES_LEN = 4
	ZLTAIL_LEN = 4
	ZLLEN_LEN  = 4
)

func test() {
	pgk :=  CreateZiplist()
	log.Println(pgk)
	AppendValue(&pgk,"inter")
	log.Println(pgk)
	AppendValue(&pgk,"get")
	log.Println(pgk)
	n := Findvalue(&pgk,"get")
	log.Println(n)
	//删除方法有问题
	DeleteEntry(&pgk,n)
	log.Println(pgk)

	log.Println(GetAllMembers(&pgk))
}

func CreateZiplist() []byte {
	pgk := new(bytes.Buffer)
	//写入
	binary.Write(pgk, binary.LittleEndian ,int32(INIT_ZLBYTES)) //zlbytes 4
	binary.Write(pgk, binary.LittleEndian ,int32(INIT_ZLTAIL)) //zltail 4
	binary.Write(pgk, binary.LittleEndian ,int32(INIT_ZLLEN)) //zllen 4
	err :=binary.Write(pgk, binary.LittleEndian ,byte(INIT_ZLEND)) //zlend 1
	if err != nil {
		log.Println(err.Error())
	}
	return pgk.Bytes()
}

func GetZipAllLen(zl *[]byte) int32 {
	t :=LittleEndianDecode((*zl)[0:4])
	return t.(int32)
}

func GetZipTailLen(zl *[]byte) int32 {
	t :=LittleEndianDecode((*zl)[4:8])
	return t.(int32)
}


func GetZipCountLen(zl *[]byte) int32{
	t :=LittleEndianDecode((*zl)[8:12])
	return t.(int32)
}

/**
 * @Author yuyunqing
 * @Description //TODO 向ziplist尾部追加
 * @Date 2:06 下午 2022/1/20
 **/
func AppendValue(zl *[]byte, s string) *[]byte {
	//获取前一个元素长度
	_ ,c , _:= getEntryByIndex(zl, GetZipTailLen(zl) )
	all_len := GetZipAllLen(zl)
	data := []byte(s)
	encoding := int32(len(data))
	p_af := int32(0)
	if c != 0 {
		p_af = c+8
	}
	entry := CreateEntry(p_af,encoding,data)
	index := all_len -1
	*zl = append((*zl)[:index] , append(entry, (*zl)[index:]...)... )
	//改变基础属性
	return UpdateBasicData(zl , encoding + 8 , c+8 ,1 )
}


/**
 * @Author yuyunqing
 * @Description //TODO 从ziplist中查询需要值
 * @Date 2:06 下午 2022/1/20
 **/
func Findvalue(zl *[]byte , string2 string ) int32 {

	fk := []byte(string2)
	count := GetZipCountLen(zl)
	start := GetZipTailLen(zl)
	for i := int32(1) ; i <= count ; i++ {

		p ,c ,e :=getEntryByIndex(zl , start)
		if c == int32(len(fk))  &&  string(e) ==string(fk)   {
			return start
		}
		start = start-p
	}
	return 0
}

/**
 * @Author yuyunqing
 * @Description //TODO 获取ziplist全部数据
 * @Date 2:31 下午 2022/1/20
 **/
func GetAllMembers(zl *[]byte) (res []string) {
	res = []string{}
	count := GetZipCountLen(zl)
	start := GetZipTailLen(zl)
	for i := int32(1) ; i <= count ; i++ {
		p ,_ ,e :=getEntryByIndex(zl , start)
		res = append(res , string(e))
		start = start-p
	}
	return
}


/**
 * @Author yuyunqing
 * @Description //TODO 根据指针下标获取当前 节点的长度
 * @Date 11:27 上午 2022/1/20
 **/
func getEntryByIndex(zl *[]byte, n int32) (int32 , int32 ,[]byte) {
	if n == 0 {
		return 0 ,0 , []byte{}
	}else {
		coding := LittleEndianDecode((*zl)[n+4:n+8]).(int32)
		prevlen := LittleEndianDecode((*zl)[n:n+4]).(int32)
		return prevlen , coding , (*zl)[n+8:n+8+coding]
	}

}


/**
 * @Author yuyunqing
 * @Description //TODO 更新zip头信息
 * @Date 2:05 下午 2022/1/20
 **/
func UpdateBasicData(zl *[]byte , zlbytes,zltail,zllen int32) *[]byte {
	all := GetZipAllLen(zl)
	tail :=GetZipTailLen(zl)
	count :=GetZipCountLen(zl)

	if tail == 0  && zltail>0{
		tail =12
	}else{

		tail = zltail+tail
		if tail ==12 {
			tail =0
		}
	}

	all_new := LittleEndianEncode(all+zlbytes)
	tail_new := LittleEndianEncode(tail)
	count_new := LittleEndianEncode(count+zllen)
	for i := 0; i < 12; i++ {
		switch i/4 {
		case  0 :
			(*zl)[i] = all_new[i%4]
			break
		case  1 :
			(*zl)[i] = tail_new[i%4]
			break
		case  2 :
			(*zl)[i] = count_new[i%4]
			break
		}

	}
	return zl
}

/**
 * @Author yuyunqing
 * @Description //TODO 删除节点
 * @Date 2:38 下午 2022/1/20
 **/
func DeleteEntry( zl *[]byte ,n int32)  {
	_ ,c ,_ :=getEntryByIndex(zl , n)

	count_s := GetZipTailLen(zl)
	//如果是最后一位的话 不更新 下一个节点的  数据
	if count_s != n {
		*zl = append((*zl)[:n+4] , (*zl)[n+c+4:]...)

	}else{
		*zl = append((*zl)[:n] , (*zl)[n+c+8:]...)
	}

	UpdateBasicData(zl , 0-(c+8) , 0-(c+8)  ,-1 )
}





/**
 * @Author yuyunqing
 * @Description //TODO 创建一个节点
 * @Date 10:16 上午 2022/1/20
 **/
func CreateEntry( prevlen , encoding  int32 , data []byte) []byte {
	pgk := new(bytes.Buffer)
	//写入
	binary.Write(pgk, binary.LittleEndian ,prevlen)
	binary.Write(pgk, binary.LittleEndian ,encoding)
	binary.Write(pgk, binary.LittleEndian ,data)

	return pgk.Bytes()
}

/**
 * @Author yuyunqing
 * @Description //TODO 小端模式解析数据
 * @Date 9:46 上午 2022/1/20
 **/
func LittleEndianDecode(i []byte)  interface{} {
	var t int32
	lengthBuff := bytes.NewBuffer(i)
	err := binary.Read(lengthBuff,binary.LittleEndian ,&t )
	if err != nil {
		log.Println(err.Error())
	}
	return t
}


/**
 * @Author yuyunqing
 * @Description //TODO 小端模式序列化 int
 * @Date 9:47 上午 2022/1/20
 **/
func LittleEndianEncode(i int32) []byte {
	pgk := new(bytes.Buffer)
	//写入
	binary.Write(pgk, binary.LittleEndian , i)
	return pgk.Bytes()
}