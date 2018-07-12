package main

import "fmt"

/*
 并发编程

1. 多进程

 开销最大的模式,系统层面的基本模式

2. 多线程

 使用最多的一种方式

3. 基于回调的异步io

 事件驱动使用异步IO 高并发的模式下会消耗cpu资源和内存

4. 协程

 不需要操作系统来进行调度,系统开销小,寄存于线程中

*/

func Count(ch chan int, i int) {
	ch <- i
	fmt.Println("counting--->")
}

func main() {

	chs := make([]chan int, 10)

	for i := 0; i < 10; i++ {
		chs[i] = make(chan int)
		go Count(chs[i], i)
	}

	for _, ch := range chs {
		i := <-ch
		fmt.Println(i)
	}
}
