/*
@Time : 2022/1/26 3:43 下午
@Author : yuyunqing
@File : quicklist
@Software: GoLand
*/
package Quicklist

import (
	"fmt"
	"go-redis/DataStructure/ZIplist"
	"log"
)

const (
	ziplist_max = 2048
)


/*
	整体逻辑 Quicklist 是一个节点为 ziplist 的双向链表

	每个节点的ziplist的最大值为 2k ,
	压入数据时，每超过ziplist 的最大值， 生成新的节点 node

	数据查找和数据 append  push  都是基于  双向链表和 ziplist 实现
 */

type Quicklist struct {
	head *QuicklistNode
	tail *QuicklistNode
	node_count  int32
	members_count  int32
}

type QuicklistNode struct {
	prv *QuicklistNode
	next *QuicklistNode
	zl *[]byte
	count int32
	members_count int32

}




func init()  {
	a:=CreateQuicklist()
	for i := 1; i <= 200; i++ {
		if i%2 ==0 {
			a.Append(fmt.Sprintf("%d" ,i))
		}else {
			a.Push(fmt.Sprintf("%d" ,i))
		}
	}
	res := a.getAllMembers()
	log.Println(res)


	for i := 1; i <= 200; i++ {

		if i%2 ==0 {
			log.Println(a.LPop())
		}else {
			log.Println(a.RPop())
		}

	}
	res = a.getAllMembers()
	log.Println(res)
}


func  CreateQuicklist()  *Quicklist  {
	ql :=  Quicklist{
		head: nil,
		tail: nil,
		members_count: int32(0),
		node_count: int32(0),
	}
	return &ql
}

/**
 * @Author yuyunqing
 * @Description //TODO 头部压入
 * @Date 2:34 下午 2022/1/26
 **/
func (ql *Quicklist) Push(msg string){
	if ql.head == nil {
		node := CreateQuicklistNode()
		ql.head = node
		ql.tail = node
		ql.node_count +=1
	}

	if ql.head.count + int32(len([]byte(msg))) > ziplist_max {
		new_head := CreateQuicklistNode()
		new_head.next = ql.head
		ql.head.prv = new_head
		ql.head = new_head
		ql.node_count +=1
	}

	ZIplist.PushEntry(ql.head.zl , msg)
	ql.head.count =  ZIplist.GetZipAllLen(ql.head.zl)
	ql.head.members_count =  ZIplist.GetZipCountLen(ql.head.zl)
	ql.members_count +=  1
	log.Println(ql.head)
}


/**
 * @Author yuyunqing
 * @Description //TODO 尾部追加
 * @Date 2:34 下午 2022/1/26
 **/
func (ql *Quicklist) Append(msg string){
	if ql.tail == nil {
		node := CreateQuicklistNode()
		ql.head = node
		ql.tail = node
		ql.node_count +=1
	}

	if ql.tail.count + int32(len([]byte(msg))) > ziplist_max {
		new_tail := CreateQuicklistNode()
		new_tail.prv = ql.tail
		ql.tail.next = new_tail
		ql.tail = new_tail
		ql.node_count +=1
	}

	ZIplist.AppendValue(ql.tail.zl , msg)
	ql.tail.count =  ZIplist.GetZipAllLen(ql.tail.zl)
	ql.tail.members_count =  ZIplist.GetZipCountLen(ql.tail.zl)
	ql.members_count +=  1

}

/**
 * @Author yuyunqing
 * @Description //TODO 创建一个节点
 * @Date 2:41 下午 2022/1/26
 **/
func CreateQuicklistNode() *QuicklistNode {
	zl := ZIplist.CreateZiplist()
	qln := QuicklistNode{
		prv :nil,
		next: nil,
		zl : &zl,
		count: ZIplist.GetZipAllLen(&zl),
	}
	return &qln
}

/**
 * @Author yuyunqing
 * @Description //TODO 获取所有成员
 * @Date 2:18 下午 2022/1/26
 **/
func (ql *Quicklist) getAllMembers() []string {
	next := ql.head
	res := []string{}
	for next != nil {
		res = append(res, ZIplist.GetAllMembers(next.zl)...)
		next = next.next
	}
	return res
}


/**
 * @Author yuyunqing
 * @Description //TODO 从头部弹出
 * @Date 2:49 下午 2022/1/26
 **/
func (ql *Quicklist) LPop() string {
	res := ""
	if ql.head.members_count >0 {
		_, p ,d   := ZIplist.GetEntryByIndex(ql.head.zl , 12)
		if p >0  {
			res = string(d)
		}
		ZIplist.DeleteEntry(ql.head.zl ,12)
		ql.head.members_count -= 1
	}

	ql.members_count = ql.members_count-1

	if ql.head.members_count == 0 {
		//将指针指向下一个元素
		if ql.node_count != 1 {
			ql.head = ql.head.next
			ql.head.prv = nil
		}else{
			ql.head = nil
			ql.tail = nil
		}

		ql.node_count -=1
	}

	return res
}




/**
 * @Author yuyunqing
 * @Description //TODO 从尾部弹出
 * @Date 2:49 下午 2022/1/26
 **/
func (ql *Quicklist) RPop() string {
	res := ""
	if ql.tail.members_count >0 {
		zl_tail :=  ZIplist.GetZipTailLen(ql.tail.zl)
		_, p ,d   :=  ZIplist.GetEntryByIndex(ql.tail.zl , zl_tail)
		if p >0  {
			res = string(d)
		}
		ZIplist.DeleteEntry(ql.tail.zl ,zl_tail)
		ql.tail.members_count -= 1
	}

	ql.members_count = ql.members_count-1

	if ql.tail.members_count == 0 {
		//将指针指向下一个元素
		if ql.node_count != 1 {
			ql.tail = ql.tail.prv
			ql.tail.next = nil
		}else{
			ql.head = nil
			ql.tail = nil
		}

		ql.node_count -=1
	}

	return res
}


