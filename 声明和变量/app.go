package main

// 导入fmt包
import "fmt"

// 声明一个顶级常量,并且推断类型,此常量当前包都能访问
const topLevel = 11111

// 声明一个可导出的顶级常量,主动声明类型
const TopLevel int = 22222

// 声明一个未初始化的int变量 初始值是0
var defaultValue int

// 声明一个函数
func addNumber(a float64, b float64) float64 {
	return a + b
}
func main() {
	// console.log
	fmt.Println(topLevel)

	fmt.Println(TopLevel)

	fmt.Println(defaultValue)

	// 一次声明多个变量
	var a, b = 1.1, 11.0

	// 默然值为空字符串
	var s string
	fmt.Println(s)

	// 声明多个
	var i,j, k int
	fmt.Println(i,j,k)

	// 函数内部可以省略var使用类型推断赋值
	result := addNumber(a, b)

	fmt.Println(result)

	// 简短变量声明的值多次声明被覆盖
	ins,ons:=1,2
	fmt.Println(ins,ons)

	//但是不能一次覆盖两个同名的值
	insOk,ons:=12,3
	fmt.Println(insOk,ons)

	/* 指针

	一个变量对应保存了对应类型值的内存空间

	只有变量才拥有内存空间,不是所有值都有内存空间比如a.f其实是a的内存空间的值

	一个指针的值实际上是另一个变量的地址如果你知道a变量的地址,那么你可以通过地址获得a.f的值
	*/
}
