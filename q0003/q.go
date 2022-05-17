package q0003

import "question/common"

func lengthOfLongestSubstring(s string) int {
	left, right, length, ret := 0, 0, len(s), 0
	if length == 0 {
		return ret
	}

	var indexMap = make(map[byte]int)

	for right < length {
		//如果右边发现重复的，将左边的窗口挪到当前元素上一次出现的位置后面
		ch := s[right]

		if i, ok := indexMap[ch]; ok {
			left = common.Max(i+1, left) //遇到 abba 这种，确保左边界不会回拨
		}

		ret = common.Max(ret, right-left+1)
		indexMap[ch] = right
		right++
	}
	return ret
}
