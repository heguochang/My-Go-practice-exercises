package main

import "fmt"

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
每个右括号都有一个对应的相同类型的左括号。

示例 1：

输入：s = "()"

输出：true

示例 2：

输入：s = "()[]{}"

输出：true

示例 3：

输入：s = "(]"

输出：false

示例 4：

输入：s = "([])"

输出：true

示例 5：

输入：s = "([)]"

输出：false

提示：

1 <= s.length <= 104
s 仅由括号 '()[]{}' 组成
*/
func main() {
	fmt.Println(isValid("()"))     // true
	fmt.Println(isValid("()[]{}")) // true
	fmt.Println(isValid("(]"))     // false
	fmt.Println(isValid("([])"))   // true
	fmt.Println(isValid("([)]"))   // false
}

func isValid(s string) bool {
	stacke := []rune{}
	pairs := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, ch := range s {
		switch ch {
		case '(', '{', '[':
			stacke = append(stacke, ch)
		case ')', '}', ']':
			if len(stacke) == 0 {
				return false
			}
			// 判断栈顶是否匹配
			top := stacke[len(stacke)-1]
			// 弹出最后一个元素
			stacke = stacke[:len(stacke)-1]
			if top != pairs[ch] {
				// 不相同的时候,就是false
				return false
			}
		}
	}

	return len(stacke) == 0
}
