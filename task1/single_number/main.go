package main

import "fmt"

/*
给你一个 非空 整数数组 nums ，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。
*/
func main() {
	nums := []int{4, 1, 2, 1, 2}
	fmt.Println(singleNumber(nums)) // 输出 4
}

func singleNumber(nums []int) int {

	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return -1
}
