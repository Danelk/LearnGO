package main

import (
	"awesomeProject2/interface_demo"
	"awesomeProject2/net_demo"
	"fmt"
	"net/http"
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
	//node1 := new(linklist_demo.DListNode)
	//node2 := new(linklist_demo.DListNode)
	//node3 := new(linklist_demo.DListNode)
	//node4 := new(linklist_demo.DListNode)
	//node1.Value = 1
	//node2.Value = 2
	//node3.Value = 3
	//node4.Value = 4
	//// node1为头节点，前指针为空，后指针指向node2
	//node1.Pre = nil
	//node1.Next = node2
	//node2.Pre = node1
	//node2.Next = node3
	//node3.Pre = node2
	//node3.Next = node4
	//// node4位尾节点，前指针指向node3，后指针为空
	//node4.Pre = node3
	//node4.Next = nil
	//linklist_demo.PrintDNode(node1)//打印链表
	//rsList := linklist_demo.DoubleReverse(node1)//翻转链表
	//linklist_demo.PrintDNode(rsList)//打印链表
	//切片
	//sl := []int{1,2,3,4,5,6,7,8,9}
	//s1 := slice_demo.ExtentSlice(sl,2,6)
	//s2 := slice_demo.ExtentSlice(s1, 3, 5) //slice扩展 可向后扩展，不可向前扩展；向后扩展可超过len(sl),不可超过cap(sl)
	//fmt.Println("slice=", sl)
	//fmt.Println("s1=", s1)
	//fmt.Printf("len(s1)= %d cap(s1)= %d", len(s1), cap(s1))
	//fmt.Println("s2=", s2)
	//fmt.Printf("len(s2)= %d cap(s2)= %d", len(s2), cap(s2))
	//fmt.Println()
	//s3 := slice_demo.AppendSlice(s2, 10) // [6,7,8] => [6,7,10]  10替换8
	//s4 := slice_demo.AppendSlice(s3, 11) // [6,7,10,9] => [6,7,10,11]  11替换9
	//s5 := slice_demo.AppendSlice(s4, 12) //s5指向一个新的底层数组  添加元素超过cap，系统会重新分类更大的底层数组
	//fmt.Println("s3=", s3)
	//fmt.Println("s4=", s4)
	//fmt.Println("s5=", s5)
	//fmt.Println("slice=", sl)
	//s6,err := slice_demo.DeleteSlice(s5,1, "f")
	//s7,err := slice_demo.DeleteSlice(s5,len(s5)-1, "b")
	//s7,err := slice_demo.DeleteSlice(s6,1, "d")
	//if err != nil{
	//	fmt.Println(err)
	//}
	//fmt.Printf("s6= %v", s6)
	//fmt.Println()
	//fmt.Printf("s7= %v", s7)
	//fmt.Println(slice_demo.LengthOfNonRepeatingSubStr("aabcbecd"))

	//net http
	//net_demo.CreateServer()
	http.HandleFunc(`/`, net_demo.Index)
	http.HandleFunc(`/f`, net_demo.First)
	http.Handle("/s/", http.StripPrefix("/s/", http.FileServer(http.Dir("./static"))))
	http.ListenAndServe(":8081", nil)
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