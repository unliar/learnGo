package main

import (
	"fmt"
	"time"
)

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

	/*
		数组是一个固定长度的元素序列,一个数组可以由0个或者多个组成

		因为数组是固定长度的,因此Go中更加常用的是slice
	*/
	var a [3]int = [3]int{1, 2, 3}

	// 省略长度
	b := [...]int{4, 5, 6}

	// 索引赋值
	c := [...]int{0: 1, 1: 2, 2: 4}

	d := [...]int{99: -10}
	// 获取末尾的元素
	fmt.Println(a[len(a)-1])

	fmt.Println(b, c, d)

	// 数组可以用于对比相等 两个数组的所有元素相等时候才相等
	for i, v := range a {
		fmt.Printf("%d %d \n", i, v)
	}

	/*
	   slice

	   切片Slice代表变长的序列,没个序列中元素都有相同的类型.
	   slice的语法和数组类似,只是没有固定长度而已
	   一个切片由指针 长度 和容量组成
	   长度len不能超过容量cap
	   slice的切片 s[i:j]
	   复制一个slice只是创建了一个slice的别名
	   slice之间不能比较
	   slice唯一可以比较的是和nil进行比较
	*/

	months := []string{"a", "b", "c"}

	// 使用make函数创建一个长度为10的切片,容量为20
	years := make([]int, 10, 20)
	fmt.Println(months, years)

	/*
	 append函数用于向slice增加元素
	 在每次append之前先检测原有slice是否可容纳元素
	 如果不是
	 则自动扩容
	 自动扩展一个原来两倍的切片
	 为什么是两倍...因为为了提高内存使用效率,避免频繁复制
	*/
	for i := 0; i < 200; i++ {
		years = append(years, i)

	}
	fmt.Println(len(years), cap(years))

	/*
	  Map

	  Go语言中,一个Map就是对哈希表的一个引用map类型写作map[K][V]
	  map中的所有key,所有value都有相同的数据类型
	  k必须是支持比较运算符的数据类型,所以可以通过这个来检测某个key是否存在
	*/
	maps := map[string]string{
		"x": "我是猪头",
		"y": "你才是猪头",
	}
	// 取出一个并不存在的key对应的值,返回值有两个z代表值,ok代表是否被赋值
	z, ok := maps["z"]
	// map和slice一样也不能进行相等比较

	// Go语言中没有set的概念 但是可以用key去实现,map[string]bool{}
	// 如果想用slice当key 可以把slice转换成string
	fmt.Println(maps["x"], z, ok)

	/*
	 结构体 struct ---->js的对象类似

	 结构体是一种聚合的数据类型
	 可以由0或者多个任意类型聚合成的实体
	*/
	type Emp struct {
		name  string
		id    int
		items map[string]string
		time  time.Time
	}

	emp := Emp{name: "天才小学生", id: 1, items: maps, time: time.Now()}
	// 访问time字段
	fmt.Println(emp.time)
	// 取出time字段的地址然后访问
	pTime := &emp.time
	fmt.Println(*pTime)
	// 完整结构体指针访问
	pEmp := &emp
	fmt.Println(pEmp.name)

	//通常一个结构体的成员名字在前 类型在后 但是如果相邻的成员类型相同可以合并到一行
	// 结构体的字段顺序不一样会产生不同的结构体类型
	type Cp struct {
		x, y, z int
	}
	cp := Cp{1, 2, 3}

	fmt.Println(*&cp.x)
}
