package main

import (
	"fmt"
	"strings"
)

/*
8.函数

函数声明包含 函数名 参数列表 返回值列表(可省略) 函数体

函数没有返回值可以省略

参数如果类型一致 不必每个参数都写出类型

函数的形参通过值的方式传递,函数的形参是函数的实参的拷贝值,形参的修改不会影响实参.

如果实参包含引用类型 指针 切片 map function channel等类型时 实参可能会被函数修改

如果函数体不存在 则可能此函数不是Go实现的


*/

// 单返回值递归
func catMe(x int) int {
	if x > 100 {
		x--
		v := catMe(x)
		return x + v
	}
	return x

}

/*
多返回值时 必须将多个返回值分配给变量
如果某些值不需要, 可以将他们赋值给_变量 表示改变量省略

如果返回值都有命名 函数可以直接return 不需要显式赋值---->建议不使用
*/

func checkPass(x int) (bool, error) {
	if x > 100 {
		return true, nil
	}
	return false, nil
}

/*
 错误处理

 1. 向上传播错误信息
  当错误的发生是必然的 将发生错误时的上下文传递给调用者

 2. 发生错误时的信息是偶然性的或者不可预知的原因, 可以选择重试并且限制重试次数

 3. 发生错误时,程序将无法运行,输出错误信息并且结束程序---->main程序
   对于库级别的程序,应该只是向上级汇报错误,并且由上级控制程序的行为

 4. 或者只是输出异常信息不需要中断程序

 5. 直接吞掉异常信息
*/

/*
匿名函数 emmmm  就是没名字的函数的意思

*/

/*
 可变参数

 是参数个数可变的函数称为可变参函数 fmt.Printf和类似的函数就是此类
*/

/**
deferred 函数

只有执行完毕deferred 后面的函数 运行流程才会继续下一步
*/
func Sum(v ...int) int {
	total := 0
	for _, val := range v {

		total += val
	}
	return total
}

var his string

func main() {

	x := catMe(1000)
	fmt.Println(x)

	t, _ := checkPass(1000)

	// 下面就是包含了一个匿名函数的栗子
	his = strings.Map(func(r rune) rune {
		return r + 9
	}, "uy")

	fmt.Println(t, his, Sum([]int{1, 3, 4, 5}...))
}
