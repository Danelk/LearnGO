package net_demo

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
)

func CreateServer()  {
	listen, _ := net.Listen("tcp", `:86`)
	//等待连接连接并建立连接
	conn, _ := listen.Accept()
	//http半双工需要先读取里面的数据
	buf := make([]byte, 1024)
	l, _ := conn.Read(buf)
	fmt.Println(string(buf[:l]))
	//通过连接发送数据
	//http 协议格式
	conn.Write([]byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain;charset=UTF-8\r\n\r\n数据开始：net模拟http"))
	conn.Close()
}

func Index(w http.ResponseWriter, r *http.Request)  {
	w.Write([]byte(`hello Index`))
}

func First(w http.ResponseWriter, r *http.Request)  {
	f, _ := os.Open(`./README.md`)
	buf,_ := ioutil.ReadAll(f)
	w.Write(buf)
}
