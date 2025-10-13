package main

import (
	"fmt"
	"sort"
)

/*
以数组 intervals 表示若干个区间的集合，其中单个区间为 intervals[i] = [starti, endi] 。
请你合并所有重叠的区间，并返回 一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间 。
*/
func main() {
	intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
	fmt.Println(marge(intervals)) // [[1 6] [8 10] [15 18]]
}

func marge(intervals [][]int) [][]int {

	// 如果输入的集合长度为空,直接返回
	if len(intervals) == 0 {
		return nil
	}

	// 把区间集合按照左值排序,方便判断是否有重合
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	// 临时变量,用来记录左值右值,先初始化,用第一个集合的左值右值给他们赋值
	curLeft := intervals[0][0]
	curRight := intervals[0][1]

	var result [][]int

	// 迭代,判断重合,从第二个开始
	for i := 1; i < len(intervals); i++ {

		// 当前集合的左值和右值
		left := intervals[i][0]
		right := intervals[i][1]

		// 如果当前集合的左值,在标记的左值和右值范围内,说明有重合,应该合并
		if curLeft <= left && left <= curRight {
			curRight = right
		} else {
			// 没有重叠,把上一段加到结果集
			result = append(result, []int{curLeft, curRight})

			// 把当前段重新设置到标记里
			curLeft = left
			curRight = right
		}
	}

	// 最后一段加到结果集
	return append(result, []int{curLeft, curRight})
}
