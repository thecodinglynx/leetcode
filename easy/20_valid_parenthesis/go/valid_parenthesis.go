package leetcode

import "github.com/thecodinglynx/stackutil"

func isValid(s string) bool {
	stack := stackutil.Stack[rune]{}
	for _, v := range s {
		if v == '[' || v == '{' || v == '(' {
			stack.Push(v)
		} else {
			if stack.IsEmpty() {
				return false
			}
			popped, _ := stack.Pop()
			if (popped == '[' && v == ']') || (popped == '(' && v == ')') || (popped == '{' && v == '}') {
				continue
			}
			return false
		}
	}
	return stack.IsEmpty()
}
