package main

// 这种匹配的问题，可以考虑栈来组合。
// golang中没有栈，需要自己动手实现
func IsValid(s string) bool {
	stack := make([]string, 0)
	combine := map[string]string{
		"(": ")",
		"{": "}",
		"[": "]",
	}
	// "(])"
	for _, v := range s {
		// 先判断栈是否为空,如果为空, 则压栈
		if len(stack) == 0 {
			stack = append(stack, string(v))
			continue
		}

		// 如果不为空， 则出栈，看是否匹配。一旦不匹配，则可以返回
		if string(v) == combine[stack[len(stack)-1]] {
			// 模拟出栈操作
			stack = stack[:len(stack)-1]
		} else {
			// 不匹配，则压栈
			stack = append(stack, string(v))
		}
	}

	return len(stack) == 0
}
