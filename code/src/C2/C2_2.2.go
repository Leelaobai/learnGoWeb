package C2

import "fmt"
import "reflect"

func HelloC2() {
	fmt.Println("Hello C2!")
}

//多种定义变量的方式
func DefineVar() {
	var v1 int
	v1 = 1
	var v2, v3, v4 int
	v2 = 2
	v3 = 3
	v4 = 4
	var v5 int = 5
	var v6, v7, v8 int = 6, 7, 8
	var v9, v10, v11 = 9, 10, 11
	v12, v13, v14 := 12, 13, 14

	fmt.Println(v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14)
}

//定义常量
//所谓常量，也就是在程序编译阶段就确定下来的值，
//而程序在运行时则无法改变该值。
//在Go程序中，常量可定义为数值、布尔值或字符串等类型。
//语法：	const constantName (type) = value
func Define_const() {
	const PI float32 = 3.1415926

	const i = 10000

	const MaxThread = 10

	const prefix = "astaxie_"

	fmt.Println(PI, i, MaxThread, prefix)
}

//GO内置基础类型

//boolean：布尔类型

func Test_boolean() {
	var isActive bool                   //一般声明
	var enabled, disabled = true, false //忽略类型的生命
	var available bool                  //一般声明
	valid := false                      //简短声明
	available = true                    //赋值操作

	fmt.Println(isActive, enabled, disabled, available, valid)
}

//数值类型

func Test_number() {
	var c complex64 = 5 + 5i
	fmt.Printf("%v", c)
}

//字符串类型

func Test_string() {
	var frenchHello string
	var emptyString string = ""
	no, yes, maybe := "no", "yes", "maybe"
	japaneseHello := "Ohaiou"
	frenchHello = "Bonjour"

	s := "hello"
	c := []byte(s) //将字符串转换为[]byte类型
	c[0] = 'c'
	s2 := string(c)

	ss := "hello"
	hh := " world!"
	sh := ss + hh

	sss := "hahaha"
	sss1 := sss[1:]
	sss2 := "c" + sss1

	m := `hello		
			world`

	fmt.Println(frenchHello, emptyString, japaneseHello, s, c, s2, no, yes, maybe, sh, sss, sss2, reflect.TypeOf(sss1))
	fmt.Println(m)
}

//error类型

//iota枚举

func Test_iota() {
	const (
		a = iota
		b = iota
		_
		d
	)
	fmt.Println(a, b, d)
}

//数组

func Test_array() {
	var arr [10]int
	arr[0] = 42
	arr[1] = 13
	a := [3]int{1, 2, 3}
	b := [10]int{1, 2, 3}
	c := [...]int{4, 5, 6}
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}

	esayArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

	fmt.Println(arr, a, b, c, doubleArray, esayArray)
}

//切片

/*
func Test_slice(){
	var fslice []int
	slice := []byte{'a','b','c','d'}

	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	var a,b []byte

	a = ar[2:5]
	b = ar[3:5]
}
*/

//条件判断语句

func Test_ifelse() {
	if x := 10; x < 10 {
		fmt.Println("这数小于10！")
	} else {
		fmt.Println("这数大于等于10！")
	}
}

//map的遍历

func Test_range_map() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	for k, v := range m {
		fmt.Println("key：", k, "value：", v)
	}
}

//函数作为值；类型
//在go语言中函数也是一种变量，它的类型就是拥有相同参数相同返回值的一种类型。
func IsOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func IsEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

//结构体
type Person struct {
	Name string
	Age  int
}

func Older(p1, p2 Person) (Person, int) {
	if p1.Age > p2.Age {
		return p1, p1.Age - p2.Age
	}
	return p2, p2.Age - p1.Age
}

//方法的继承

type Human struct {
	name  string
	age   int
	phone string
}
type Student struct {
	Human  //匿名字段
	school string
}
type Employee struct {
	Human   //匿名字段
	company string
}

//Human定义method
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Employee的method重写Human的method
func (e *Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone) //Yes you can split into 2 lines here.
}
func TestMethod() {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}
	mark.SayHi()
	sam.SayHi()
	sam.Human.SayHi()
}
