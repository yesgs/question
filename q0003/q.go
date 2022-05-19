package q0003

import "question/common"

//给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。

//示例1:
//
//输入: s = "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: s = "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: s = "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是"wke"，所以其长度为 3。
//    请注意，你的答案必须是 子串 的长度，"pwke"是一个子序列，不是子串。
//

func lengthOfLongestSubstring(s string) int {
	left, right, length, ret := 0, 0, len(s), 0
	if length == 0 {
		return ret
	}

	var indexMap = make(map[byte]int)

	for right < length {
		ch := s[right]

		//如果右边发现重复的，将左边的窗口挪到当前元素上一次出现的位置后面
		if i, ok := indexMap[ch]; ok {
			left = common.Max(i+1, left)
			//遇到 abba 这种，确保左边界不会回拨
			//当左边界走到第二个b时，下一步判断a，a在b前面，就不要回去了
		}

		ret = common.Max(ret, right-left+1) //取最大值
		indexMap[ch] = right
		right++
	}
	return ret
}
