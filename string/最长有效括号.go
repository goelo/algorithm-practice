package main

func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	stack := make([]string, 0, len(s))
	combine := map[string]string{
		"(": ")",
	}
	var max int
	for _, v := range s {
		var count int
		if len(stack) == 0 || string(v) == "(" {
			stack = append(stack, string(v))
		}
		// 如果匹配，++
		if string(v) == combine[stack[len(stack)-1]] {
			count += 2
			stack = stack[0 : len(stack)-1]
			if count > max {
				max = count
			}
		} else {
			count = 0
			stack = append(stack, string(v))
		}
	}
	return max
}
