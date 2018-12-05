package main

import (
	"fmt"
)

func Sum(arr []int64) int64 {
	if len(arr) == 1 {
		return arr[0]
	}
	return arr[0] + Sum(arr[1:])
}

func Max(arr []int64) int64 {
	if len(arr) == 0 {
		panic("数组为空")
	}
	if len(arr) == 1 {
		return arr[0]
	} else if arr[0] > Max(arr[1:]) {
		return arr[0]
	} else {
		return Max(arr[1:])
	}

}

func Total(arr []int64) int64 {
	if len(arr) == 0 {
		return 0
	}
	return 1 + Total(arr[1:])
}

func Sort(arr []int64) []int64 {
	if len(arr) < 2 {
		return arr
	}
	i := arr[0]
	var l []int64
	var r []int64
	for _, item := range arr[1:] {
		if item <= i {
			l = append(l, item)
		} else {
			r = append(r, item)
		}
	}

	ll := append(Sort(l), i)
	tt := append(ll, Sort(r)...)
	return tt

}
func main() {
	a := []int64{1, 2, 3, 4, 5, 7, 8, 8, 9, 90, 434, 34, 3, 43, 43, 4, 34, 3, 43, 4, 34, 3, 43, 4, 3, 4, 34, 3, 4}
	s := Sum(a)
	m := Max(a)
	t := Total(a)
	fmt.Println("sum:", s)
	fmt.Println("max:", m)
	fmt.Println("total:", t)
	fmt.Println(Sort(a))
}
