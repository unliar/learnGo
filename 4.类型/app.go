package main

import "fmt"

// 类型----->在底层类型之上包装的一层

// type 类型名字 底层类型

// 许多类型都会定义一个String方法,用于fmt打印
type IsCode bool

var hah IsCode = true

type Celsius float64

type Fahrenheit float64

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	f := CToF(1.2)

	c := FToC(2.4)

	// 强制转换类型
	a := f - Fahrenheit(c)
	var b float64 = float64(30.00)
	fmt.Println(a, b)
	fmt.Println(hah)
}
