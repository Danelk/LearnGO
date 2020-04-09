package point_demo

import (
	"fmt"
	"reflect"
)

//go 指针不能进行指针运算

func TestPoint()  {
	var count int = 20
	var CountPoint *int
	var CountPoint1 *int
	CountPoint = &count

	fmt.Printf("count 变量的地址是 ：%x \n", &count)
	if CountPoint != nil{
		fmt.Printf("CountPoint 变量的存储地址是 ：%x \n", CountPoint)
	}
	fmt.Printf("CountPoint 指针指向地址的值 ：%d \n", *CountPoint)
	fmt.Printf("CountPoint1 变量的存储地址是 ：%x \n", CountPoint1)
	fmt.Println("CountPoint1 变量的值:", CountPoint1)
}

func TestPointArr()  {
	//指针数组   数组的每一个元素都是指针
	a,b := 1,2
	pointArr := [...]*int{&a,&b}
	type1 := reflect.TypeOf(pointArr)
	fmt.Println("指针数组 pointArr:",type1)

	//数组指针  指向一个数组的指针
	arr := [...] int{2,3,4}
	arrPoint := &arr
	types := reflect.TypeOf(arrPoint)
	fmt.Println("数组指针 ArrPoint:", types)
}
