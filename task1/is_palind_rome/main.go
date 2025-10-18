package main

import "fmt"

/*
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。

回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。

例如，121 是回文，而 123 不是。
*/
func main() {
	tests := []int{121, -121, 10, 12321, 1234321, 0}
	for _, v := range tests {
		fmt.Printf("%d -> str=%v \n", v, isPalindrome(v))
	}
}

func isPalindrome(num int) bool {

	// 负数,不等于零且末位为0,一定不是
	if num < 0 || (num%10 == 0 && num != 0) {
		return false
	}

	// 转字符串,然后双指针法
	s := fmt.Sprintf("%d", num)
	n := len(s)
	for i := 0; i < n/2; i++ {
		// 左指针不等于右指针就不是回文数
		if s[i] != s[n-i-1] {
			return false
		}
	}
	return true
}
