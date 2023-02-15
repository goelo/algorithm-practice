package main

var numMap = map[string]int{
	"0": 0,
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
}

var symbolMap = map[string]bool{
	".": true,
	" ": true,
}

func myAtoi(s string) int {
	b := []byte(s)
	// 确定开始位置
	var i int
	for i = 0; i < len(s); i++ {
		if s[i] != ' ' {
			break
		}
	}
	b = b[i : len(b)-1]
}
