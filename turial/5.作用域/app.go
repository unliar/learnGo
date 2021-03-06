package main

import "fmt"

/*
生命周期 VS 作用域

生命周期是指程序在运行时变量存在的有效时间范围,是一个运行时的概念.

作用域对应的时一个源代码的文本区域,是编译时候的一个属性.

*/

/*
语法块

语法块是指花括弧{}所包含的一系列语句.

在语法块内声明的变量无法被外部语法块访问.

语法块决定了内部声明的名字的作用域范围.

声明语句的词法域决定了作用域的范围大小.

内置的类型 函数 常量 作用域是全局的

包级别的声明可以在这个包的任何区域使用

编译器遇到一个名字的引用时,是从最内层的作用域向全局作用域查找,查找失败则报错未定义的声明之类的.内部的声明会覆盖外层的声明.
在内部寻找到了声明的定义,则中断了继续寻找的行为[作用域链....]


*/

// 包级别的声明
var a string = "你是猪"

func main() {
	b := "我是小可爱"

	//打印了包级别的a的值
	fmt.Println(a)

	// 打印当前a的指针
	fmt.Println(&a)

	// 打印块级声明b
	fmt.Println(b)

	// 给包级别的a赋值
	a = "大猪头"

	// 打印当前a的指针
	fmt.Println(&a)

	//声明一个函数级别的a [推断赋值]
	a := "嘻嘻嘻 打我啊"

	// 打印当前a的指针-->不和之前的指针相同 所以已经是新的声明了
	fmt.Println(&a)

	fmt.Println(a)

	x := "is life always this hard? "
	for i := 0; i < len(x); i++ {
		fmt.Printf("%c", x[i])
	}

	for _, y := range x {
		fmt.Println(string(y))
	}
}
