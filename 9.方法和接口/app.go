package main

/*
方法

就是一类带特殊接受者参数的函数

*/

type V struct {
	X, Y int
}

// 函数使用 V结构体类型的数据作为接受者
func (v V) Add() int {
	return v.Y + v.X
}

// 函数使用指针作为接收者.
// 隐式转换,可以使用指针 也可以使用普通值
func (v *V) Sum() int {
	return v.X + v.Y
}
func main() {
	// 普通接收者
	v := V{1, 9}
	a := v.Add()
	// 指针接收者
	s := &v
	ss := s.Sum()

	println(a, ss)
}
