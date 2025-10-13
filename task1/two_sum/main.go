package main

import "fmt"

// 两数之和
// 考察：数组遍历、map使用
// 题目：给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那两个整数
func main() {
	ints := []int{2, 7, 11, 15}
	target := 9
	result := findTarget(ints, target)
	fmt.Println(result)
}

/*
入参nums:给定的整数数组
入参target:目标值

出参两个数的数组下标
*/
func findTarget(nums []int, target int) []int {
	// 创建一个map,k是数组的值,v是数组的index
	result := make(map[int]int)
	for i, v := range nums {
		// 对于当前值,map里是否已经存在目标值-当前值的结果
		need := target - v
		// 如果存在的话,那么目标值就是由这两个数构成
		if j, ok := result[need]; ok {
			return []int{j, i}
		}

		// 不存在就把这个数加到map中
		result[v] = i
	}

	// 找不到返回空
	return nil
}
