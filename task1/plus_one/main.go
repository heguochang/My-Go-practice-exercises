package main

import "fmt"

/*
给定一个表示 大整数 的整数数组 digits，其中 digits[i] 是整数的第 i 位数字。这些数字按从左到右，从最高位到最低位排列。这个大整数不包含任何前导 0。
将大整数加 1，并返回结果的数字数组。
示例 1：

输入：digits = [1,2,3]
输出：[1,2,4]
解释：输入数组表示数字 123。
加 1 后得到 123 + 1 = 124。
因此，结果应该是 [1,2,4]。
示例 2：

输入：digits = [4,3,2,1]
输出：[4,3,2,2]
解释：输入数组表示数字 4321。
加 1 后得到 4321 + 1 = 4322。
因此，结果应该是 [4,3,2,2]。
示例 3：

输入：digits = [9]
输出：[1,0]
解释：输入数组表示数字 9。
加 1 得到了 9 + 1 = 10。
因此，结果应该是 [1,0]。
*/
func main() {
	fmt.Println(plusOne([]int{1, 2, 3})) // [1 2 4]
	fmt.Println(plusOne([]int{9}))       // [1 0]
	fmt.Println(plusOne([]int{9, 9, 9})) // [1 0 0 0]
}
func plusOne(digits []int) []int {
	// 从后往前遍历
	add := 0
	for i := len(digits) - 1; i >= 0; i-- {
		current := digits[i]

		// 如果是最后一个数,先加1
		if i == len(digits)-1 {
			current = current + 1
		}

		// 上次有进位
		if add != 0 {
			current = current + 1
			add = 0
		}

		// 判断是否要进位,要进位的话要记录进位标记,同时把当前位设置为0
		if current == 10 {
			add = 1
			digits[i] = 0
		} else {
			digits[i] = current
		}
	}

	// 查看最后的进位标记是否为
	if add == 0 {
		return digits
	} else {
		return append([]int{1}, digits...)
	}
}
