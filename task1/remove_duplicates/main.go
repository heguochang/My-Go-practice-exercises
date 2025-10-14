package main

import "fmt"

/*
给你一个 非严格递增排列 的数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。元素的 相对顺序 应该保持 一致 。然后返回 nums 中唯一元素的个数。
考虑 nums 的唯一元素的数量为 k ，你需要做以下事情确保你的题解可以被通过：
更改数组 nums ，使 nums 的前 k 个元素包含唯一元素，并按照它们最初在 nums 中出现的顺序排列。nums 的其余元素与 nums 的大小不重要。
返回 k 。
*/
func main() {
	nums := []int{1, 1, 2, 2, 3, 4, 4, 5}
	k := removeDuplicates(nums)
	fmt.Println("唯一元素个数:", k)
	fmt.Println("修改后的数组:", nums[:k])
}

// 删除
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// 设置一个指针作为标记,记录新元素要写入的位置,当前指向第二个位置
	k := 1

	// 设置第二个指针,让他从第二个位置开始,去寻找新出现的元素,找到就写到第一个指针标记的位置
	for i := 1; i < len(nums); i++ {
		// 发现了,把新发现的元素写过来,然后写入标记后移一个位置
		if nums[i] != nums[i-1] {
			nums[k] = nums[i]
			k++
		}
	}
	return k
}
