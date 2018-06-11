package main

import "fmt"

/*
数组和结构体都有固定内存大小的数据结构
slice和map则是动态的数据结构

数组: 数组是由同构的元素组成的,每个数组中的元素都是由完全相同的类型

结构体:则是由异构的元素组成
*/

type Messages struct {
	id      int
	name    string
	content string
	from    string
}

func main() {
	// 数组是一个固定长度的元素序列,一个数组可以由0个或者多个组成
	// 因为数组是固定长度的,因此Go中更加常用的是slice
	var a [3]int = [3]int{1, 2, 3}

	// 省略长度
	b := [...]int{4, 5, 6}

	// 索引赋值
	c := [...]int{0: 1, 1: 2, 2: 4}

	d := [...]int{99: -10}
	// 获取末尾的元素
	fmt.Println(a[len(a)-1])

	fmt.Println(b, c, d)
	for i, v := range a {
		fmt.Printf("%d %d \n", i, v)
	}
}
