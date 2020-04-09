package suger

import "fmt"
//语法糖
//"..."  可变长度的参数 类型需一致
func Suger(values...string)  {
	for _,v := range values{
		fmt.Println("v:", v)
	}
}

// ":" 推断数据类型 赋值 声明
func Suger1()  {
	value := "A"
	fmt.Println("value:", value)
}
