package main

//声明一个结构体--->Node.js的{}
type Messages struct {
	name string
	year int
}

func main() {

	// 赋值

	//声明并且赋值
	var x int = 1
	var tt struct {
		x int
		y int
	}
	//指针间接赋值
	p := &x
	*p = *p + 10

	// 结构体示例
	mes := Messages{"你是猪头", 100}

	// 结构体赋值
	mes.name = "我才不是猪头"

	// 其他语言也有的算术运算快速赋值
	mes.year++
	mes.year *= mes.year
	mes.year--

	// 元祖赋值----一个语句赋值多个变量
	// 在有多个返回值的地方很常用
	// 元祖赋值可以使的琐碎赋值更加紧凑
	var n, m int
	n, m = 100, 111
	// 快速交换值
	n, m = m, n

	// 隐式赋值 实际上对medals[0]也赋值了
	medals := []string{"you", "are", "the one"}

	println(x, p, tt.y, mes.name, mes.year)
	println(n, m, medals[0])
}
