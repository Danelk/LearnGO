package main

import (
	"awesomeProject2/interface_demo"
	"awesomeProject2/suger"
	"fmt"
)
//package main
//import "fmt"
//func fibonacci(c, quit chan int) {
//	x, y := 0, 1
//	for {
//		select {
//		case c <- x:
//			x, y = y, x+y
//		case <-quit:
//			fmt.Println("quit")
//			return
//		}
//	}
//}
//func main() {
//	c := make(chan int)
//	quit := make(chan int)
//	//go func() {
//		for i := 0; i < 10; i++ {
//			fmt.Println(<-c)
//		}
//		quit <- 0
//	}()
//	fibonacci(c, quit)
//}
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
	suger.Suger1()
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