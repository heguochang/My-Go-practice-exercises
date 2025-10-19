package main

import "fmt"

/*
题目 ：实现一个函数，接收一个整数切片的指针，将切片中的每个元素乘以2。
考察点 ：指针运算、切片操作。
*/
func main() {
	s := []int{1, 2, 3}
	process(&s)
	fmt.Println(s)
}

func process(nums *[]int) {
	s := *nums
	for i, _ := range s {
		s[i] *= 2
	}
}
