package main

import "fmt"

/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1：

输入：strs = ["flower","flow","flight"]
输出："fl"
示例 2：

输入：strs = ["dog","racecar","car"]
输出：""
解释：输入不存在公共前缀。
*/
func main() {
	strings := []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strings))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	// 纵向比对即可,停止时机:有任意一列不同了或者有一个字符串长度不够了
	for i := 0; i < len(strs[0]); i++ {
		// 当前的基准字符
		char := strs[0][i]
		// 对比剩下的字符串当前位置的字符
		for j := 1; j < len(strs); j++ {
			// 当前位置已经比剩下的字符串总长度长 或者 出现了不一致
			if i >= len(strs[j]) || char != strs[j][i] {
				return strs[0][:i]
			}
		}
	}

	// 如果没有提前返回就是完全匹配
	return strs[0]
}
