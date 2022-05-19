package q0005

//给你一个字符串 s，找到 s 中最长的回文子串。

//示例 1：
//
//输入：s = "babad"
//输出："bab"
//解释："aba" 同样是符合题意的答案。
//示例 2：
//
//输入：s = "cbbd"
//输出："bb"

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}

	//遍历原始字符串的每一个字符
	res := ""
	for i := 0; i < len(s); i++ {
		odd := maxPalindrome(s, i, i)    //找出奇数长度的回文串，即找到以 s[i] 为中心的回文串
		even := maxPalindrome(s, i, i+1) //找出偶数长度的回文串，即找到以 s[i] 和 s[i+1] 为中心的回文串

		//比较奇数回文串和偶数回文串的长度 取最大的
		max := even
		if odd[1] > even[1] {
			max = odd
		}
		if len(res) < max[1] {
			res = s[max[0] : max[0]+max[1]]
		}
	}
	return res
}

func maxPalindrome(s string, left, right int) []int {
	length := len(s)
	for left >= 0 && right < length {
		//偶数长度 中间两个必须相等
		//奇数长度 中间点只有一个
		if s[left] == s[right] {
			left--
			right++
		} else {
			break
		}
	}
	//第一个元素 左边界起始下标
	//第二个元素 子串长度
	return []int{left + 1, right - (left + 1)}
}
