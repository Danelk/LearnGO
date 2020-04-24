package net_demo

import (
	"fmt"
	"net"
)

func Request()  {
	conn,_ := net.Dial("tcp",`:86`)
	buf := make([]byte,1024)
	l,_ := conn.Read(buf)
	fmt.Println(string(buf[:l]))
}