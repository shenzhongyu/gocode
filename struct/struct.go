package main

import "fmt"

//结构体声明
type Books struct {
	title, author, subject string
}

type person struct {
	name string
	age  int
}

type rect struct {
	width  int
	height int
}

//"齐天传" 变量类型声明
var (
	Qitianzhuan Books
)

func main() {

	//结构体类型变量赋值
	Qitianzhuan.title = "孙悟空的师父究竟是谁，燃灯吗？"
	Qitianzhuan.author = "楚阳冬"
	Qitianzhuan.subject = "这部小说，让不看西游的人了解西游，让看过西游的人更懂西游，让隐藏着的西游世界展露全貌，让那只看似神通盖世却受尽欺侮的猴子，打破枷锁、真正自由。"
	//打印 Qitianzhuan 信息
	fmt.Printf("《齐天传》： %s\n", Qitianzhuan)
	//打印 Qitianzhuan 结构体成员
	fmt.Printf("《齐天传》: %s\n", Qitianzhuan.subject)

	// 结构体直接赋值
	fmt.Println(person{name: "JoketorSen", age: 17})
	fmt.Println(person{"Sensir", 30})

	//结构体指针示列对比
	fmt.Println(&person{name: "Axecom", age: 40})
	//var person_pointer *person
	person_pointer := &person{"Axecom", 40}
	fmt.Println(person_pointer.age)

	person_pointer.age = 51
	fmt.Println(person_pointer.age)

	r := rect{width: 750, height: 362}
	fmt.Printf("area：%d\n", r.area())
	fmt.Printf("perim：%d\n", r.perim())

	rr := &r
	fmt.Printf("area：%d\n", rr.area())
	fmt.Printf("perim：%d\n", rr.perim())

}

// 结构体指针函数传参
func (r *rect) area() int {
	return r.width * r.height
}

// 结构体函数传参
func (r rect) perim() int {
	return 2*r.width + 2*r.height
}
