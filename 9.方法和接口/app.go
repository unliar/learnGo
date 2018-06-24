package main

import "fmt"

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

// 接口类型是一组方法签名的定义的集合
// 接口类型的值 可以保存任何实现了这些方法的值
type Abser interface {
	Abs() float64
}
type MyFloat float64

type MyStruct struct {
	X, Y float64
}

func (f MyFloat) Abs() float64 {
	return float64(f) - 0.5
}
func (s MyStruct) Abs() float64 {
	return s.X + s.Y
}
func main() {
	// 普通接收者
	v := V{1, 9}
	a := v.Add()
	// 指针接收者
	s := &v
	ss := s.Sum()

	// abser 实现了Abs方法,但是接收者是不同类型的数据
	var aa Abser
	var bb Abser

	mf := MyFloat(2)
	ms := MyStruct{1.0, 2.6}

	aa = mf
	fmt.Println(aa.Abs())

	bb = ms
	fmt.Println(bb.Abs())
	fmt.Println(aa)
	fmt.Println(a, ss)
}
