package main

var combinations []string
var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}
func letterCombinations(digits string) []string {
	// 全部组合一般是用搜索，搜索常用的是DFS和BFS
	// 定义map对应字符串
	// 本题用到了DFS+回溯算法
	if len(digits) == 0  {
		return []string{}
	}


	combinations = []string{}
	backTracking(digits, 0, "")
	return combinations

}

func backTracking(digits string, index int, combination string) {
	// 这里的递归退出条件，很巧妙。第一次思考不到。要记住。
	// 当这里index==len(digits)， 也意味着dfs到底了。可以退出了
	if index == len(digits) {
		combinations = append(combinations, combination)
		return
	}
	// 没有走完递归，处理
	// 找到每个字母对应的选择
	letters := phoneMap[string(digits[index])]
	// 遍历每个字母, 来逐级组合
	for i := 0; i<len(letters); i++{
		// 往下一个字母看组合，一直递归底部，然后从底部开始合并
		backTracking(digits, index+1, combination+string(letters[i]))
	}
}
