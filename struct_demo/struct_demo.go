package struct_demo

import "fmt"

type Animal struct {
	Color string
}

type Dog struct {
	Animal // 结构体继承
	Id 		int
	Name 	string
	Age 	int
}
type Cat struct {
	Animal // 结构体继承
	Id 		int
	Name 	string
	Age 	int
}

func TestForStruct()  {
	//方式1
	//var dog Dog
	//dog.Id = 1
	//dog.Name = "nara"
	//dog.Age = 3

	//方式2
	//dog := Dog{1,"nara", 2}

	//方式3
	//dog := new(Dog)
	//dog.Id = 1
	//dog.Name = "nara"
	//dog.Age = 3
	//
	//fmt.Println("dog : ", dog)
}

func (d *Dog) Run() string {
	fmt.Println("ID : ", d.Id ," Dog is Run")
	return " Dog is Run"
}

func (d *Dog) Eat() string {
	fmt.Println("eat eat")
	return "eat eat"
}

func (d *Cat) Run() string {
	fmt.Println("ID : ", d.Id ," Cat is Run")
	return " Cat is Run"
}

func (d *Cat) Eat() string {
	fmt.Println("Cat eat eat")
	return "Cat eat eat"
}