package main

//导入fmt包
import "fmt"

//声明一个顶级常量,并且推断类型,此常量当前包都能访问
const topLevel = 11111

//声明一个可导出的顶级常量,主动声明类型
const TopLevel int = 22222

// 声明一个未初始化的int变量 初始值是0
var defaultValue int

//声明一个函数
func addNumber(a float64, b float64) float64 {
	return a + b
}
func main() {

	// console.log
	fmt.Println(topLevel)

	fmt.Println(TopLevel)

	fmt.Println(defaultValue)
	// 一次声明多个变量
	const a, b = 1, 11

	//类型推断赋值
	result := addNumber(a, b)
	fmt.Print(result)
}
