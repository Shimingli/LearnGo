package main
import (
	"fmt"
)
//什么是接口
//
//接⼝是一个或多个方法签名的集合，任何非接口类型只要拥有与之对应的全部方法实现 (包括相同的名称、参数列表以及返回值。)，就表示它"实现" 了该接口，无需显式在该类型上添加接口声明。此种方式又被称作Duck Type。
//
//接口的实例化
//
//接口是可被实例化的类型，而不仅仅是语言上的约束规范。当我们创建接口变量时，将会为其分配内存，并将赋值给它的对象拷贝存储。将对象赋值给接口变量时，会发生值拷贝行为。没错，接口内部存储的是指向这个复制品的指针。而且我们无法修改这个复制品的状态，也无法获取其指针。
//
//接口是对一个对象的各取所需。需要那些特性就定义一个相应特性的接口。
//
//go的实现把不同类型的行为特性区分开来，又从数据本质的层面把它们联系在一起（都是一个人的不同身份），这充分体现了go语言对对象和数据的理解。
func main (){

	var ming People

	ming = &Chinese{Name : "ming", Energy : 10}

	ming.Eat(5)

	ming.Sleep(3)

	ming.Work(8)

}


//这是基类
type People interface {

	Sleep(e int) bool

	Eat(e int) bool

	Work(e int) bool

}
//这是一个中国人 ，
type Chinese struct{

	Name string

	Energy int

}
//中国人  实现 接口中的 sleep的定义
func (s *Chinese) Sleep(e int) bool {
	s.Energy = s.Energy + e
	fmt.Printf("Chinese %s Sleep, Energy = %d\r\n", s.Name, s.Energy)
	return true
}
func (s *Chinese) Eat(e int) bool {
	s.Energy = s.Energy + e
	fmt.Printf("Chinese %s Eat, Energy = %d\r\n", s.Name, s.Energy)
	return true
}

func (s *Chinese) Work(e int) bool {

	if s.Energy > e {

		s.Energy = s.Energy - e

		fmt.Printf("Chinese %s Work, Energy = %d\r\n", s.Name, s.Energy)

		return true

	}

	fmt.Printf("Chinese %s can not Work, Energy = %d\r\n", s.Name, s.Energy)

	return false

}