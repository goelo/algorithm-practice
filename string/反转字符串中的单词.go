package main

func reverseWords(s string) string {
	// 不在原地处理, 转成byte因为string不允许更改内部的值
	b := []byte(s)
	// 1. 移除多余空格
	// 双指针移除多余空格
	slowIdx, fastIdx := 0, 0
	// 移除最左边空格
	for len(b) > 0 && fastIdx < len(b) && b[fastIdx] == ' ' {
		fastIdx++
	}
	// 移除单词之间的多余空格
	// 这段代码思想是判断是否有多余重复元素,是用fastIdx-1和fastIdx判断
	// 注意不是fastIdx+1和fastIdx
	for ; fastIdx < len(b); fastIdx++ {
		if fastIdx-1 > 0 && b[fastIdx-1] == b[fastIdx] && b[fastIdx] == ' ' {
			// 跳过空格
			continue
		}
		// slowIdx位置填入值
		b[slowIdx] = b[fastIdx]
		slowIdx++
	}
	// 移除尾部冗余空格
	// 根据上面的算法，slowIdx会移动至最后一位
	// 因为当末尾有两个空格的时候，slowIdx还是会移动到最后一位。
	// 如果超过两个空格，slowIdx会停止在末尾的第二个空格位置。
	// 因此应该判断slowIdx-1的位置是否是空格，才能移除所有尾部多余的空格
	if slowIdx-1 > 0 && b[slowIdx-1] == ' ' {
		b = b[:slowIdx-1]
	} else {
		// 末尾只有一个空格
		b = b[:slowIdx]
	}
	// 2. 反转整个字符串
	reverseWord([]byte(b), 0, len(b)-1)
	// 3. 然后反转内部的单个单词
	//
	start := 0
	for start < len(b) {
		end := start
		for end < len(b) && b[end] != ' ' {
			end++
		}
		// 反转单词
		reverseWord([]byte(b), start, end-1)
		start = end
		start++
	}
	return string(b)
}

func reverseWord(s []byte, start, end int) {
	for start < end {
		s[start], s[end] = s[end], s[start]
		start++
		end--
	}
}
