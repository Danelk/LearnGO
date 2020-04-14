package json_demo

import (
	"encoding/json"
	"fmt"
)
//结构体中如果写了tag类型那么在序列化和反序列化的时候使用tag去寻找key
type Server struct {
	ServerName string `json:"server_name"`
	ServerIp  string `json:"server_ip"`
	ServerPort int `json:"server_port"`
}

//序列化结构体
func SerializeStruct()  {
	server := new(Server)
	server.ServerName = "demo"
	server.ServerIp = "1217.0.0.1"
	server.ServerPort = 80

	b,err := json.Marshal(server)
	if err != nil {
		fmt.Println("json err:",err.Error())
		return
	}
	fmt.Println("json : ", string(b))
}

//序列化Map
func SerializeMap()  {
	server := make(map[string]interface{})
	server["ServerName"] = "demo-map"
	server["ServerIp"] = "1217.0.0.1"
	server["ServerPort"] = 81

	b,err := json.Marshal(server)
	if err != nil {
		fmt.Println("json err:",err.Error())
		return
	}
	fmt.Println("json : ", string(b))
}

//反序列化结构体
func DeSerializeStruct()  {
	jsonString := `{"server_name":"demo","server_ip":"1217.0.0.1","server_port":80}`
	server := new(Server)
	err := json.Unmarshal([]byte(jsonString), &server)
	if err != nil {
		fmt.Println("json decode err:",err.Error())
		return
	}
	fmt.Println("json decode : ", server)
}

//反序列化Map
func DeSerializeMap()  {
	jsonString := `{"server_name":"demo-map","server_ip":"127.0.0.1","server_port":81}`
	server := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonString), &server)
	if err != nil {
		fmt.Println("json decode err:",err.Error())
		return
	}
	fmt.Println("json : ", server)
}