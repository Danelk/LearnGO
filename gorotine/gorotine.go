package gorotine

import (
	"fmt"
	"sync"
	"time"
)

var chanInt = make(chan int, 10)
var timeout = make(chan bool)

var WG  = sync.WaitGroup{}

func Loop()  {
	for i:=1; i < 11; i++ {
		time.Sleep(time.Second * 1)
		fmt.Printf("%d,", i)
	}
}
//发送数据
func Send()  {
	time.Sleep(time.Second * 1)
	chanInt <- 1
	time.Sleep(time.Second * 1)
	chanInt <- 2
	time.Sleep(time.Second * 1)
	chanInt <- 3
	time.Sleep(time.Second * 2)
	timeout <- true
}
//接收数据
func Receive()  {
	//num := <- chanInt
	//fmt.Printf("num is %d,", num)
	//num1 := <- chanInt
	//fmt.Printf("num1 is %d,", num1)
	//num2 := <- chanInt
	//fmt.Printf("num2 is %d,", num2)
	for{
		select {
		case num := <- chanInt:
			fmt.Printf("num is %d,", num)
		case <-timeout:
			fmt.Println("timeout...")

		}
	}
}
//读取数据
func Read()  {
	for i:=0; i<3 ;i++  {
		WG.Add(1)
	}
}
//写入数据
func Write()  {
	for i:=0; i<3 ; i++  {
		time.Sleep(time.Second * 1)
		fmt.Println("DONE->", i)
		WG.Done()
	}
}