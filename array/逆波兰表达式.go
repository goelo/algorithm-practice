package main

import "strconv"

func evalRPN(tokens []string) int {
	oper := map[string]bool{
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}

	stack := make([]string, 0, len(tokens))
	for _, token := range tokens {
		// 如果是符号，出栈
		if _, ok := oper[token]; ok {
			// 弹出栈顶的两个元素,计算结果并压栈
			num1, num2 := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[0 : len(stack)-2]
			num := operator(num1, num2, token)
			stack = append(stack, num)
			continue
		}
		// 数字，压栈
		stack = append(stack, token)
	}
	num, _ := strconv.Atoi(stack[len(stack)-1])
	return num
}

func operator(num1, num2 string, oper string) string {
	a, _ := strconv.Atoi(num1)
	b, _ := strconv.Atoi(num2)
	var num int
	switch oper {
	case "+":
		num = a + b
	case "*":
		num = a * b
	case "-":
		num = b - a
	case "/":
		num = b / a
	}
	return strconv.Itoa(num)
}
