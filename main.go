package main

import (
	"awesomeProject2/interface_demo"
	"awesomeProject2/linklist_demo"
	"fmt"
)

func main() {
	//makeSlice()
	//makeMap()

	//struct_demo.TestForStruct()
	//dog := new(struct_demo.Dog)
	//dog.Id = 1
	//dog.Name = "nara"
	//dog.Age = 3
	//dog.Color = "red"
	//dog.Run() //作用域 通过首字母大小写判断
	//dog.Eat()

	//测试接口定义变量
	//var b interface_demo.Behavior
	//b := new(struct_demo.Cat)
	//c := new(struct_demo.Dog)
	////b.Run()
	//action(b)
	//action(c)

	//fmt.Printf("cpu num = %d", runtime.NumCPU())
	//runtime.GOMAXPROCS(runtime.NumCPU() -1)

	//go gorotine.Loop()
	//go gorotine.Loop()
	//time.Sleep(time.Second * 5)
	//启动发送数据协程
	//go gorotine.Send()
	//启动接收数据协程
	//go gorotine.Receive()

	//协程同步
	//gorotine.Read()
	//go gorotine.Write()
	//gorotine.WG.Wait()
	//fmt.Println("ALL DONE")
	//time.Sleep(time.Second * 10)

	//指针
	//point_demo.TestPoint()
	//数组指针 指针数组
	//point_demo.TestPointArr()

	//json序列化
	//json_demo.SerializeStruct()
	//json_demo.SerializeMap()
	//json反序列化
	//json_demo.DeSerializeStruct()
	//json_demo.DeSerializeMap()

	//语法糖
	//suger.Suger("A","B","C")
	//suger.Suger1()

	//链表
	//list := new(linklist_demo.ListNode)
	//linklist_demo.CreateNode(list, 5) //创建链表
	node1 := new(linklist_demo.DListNode)
	node2 := new(linklist_demo.DListNode)
	node3 := new(linklist_demo.DListNode)
	node4 := new(linklist_demo.DListNode)
	node1.Value = 1
	node2.Value = 2
	node3.Value = 3
	node4.Value = 4
	// node1为头节点，前指针为空，后指针指向node2
	node1.Pre = nil
	node1.Next = node2
	node2.Pre = node1
	node2.Next = node3
	node3.Pre = node2
	node3.Next = node4
	// node4位尾节点，前指针指向node3，后指针为空
	node4.Pre = node3
	node4.Next = nil
	linklist_demo.PrintDNode(node1)//打印链表
	rsList := linklist_demo.DoubleReverse(node1)//翻转链表
	linklist_demo.PrintDNode(rsList)//打印链表
}

//测试定义接口实现方法
func action(b interface_demo.Behavior) string {
	b.Run()
	b.Eat()
	return ""
}

func makeSlice()  {
	s := make([]string, 3)
	s[0] = "123"
	s[1] = "234"
	s[2] = "345"
	fmt.Println(s)
}

func makeMap()  {
	m := make(map[string]string)
	m["t1"] = "111"
	m["t2"] = "222"
	m["t3"] = "333"
	fmt.Println(m)
}