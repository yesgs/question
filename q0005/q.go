package q0005

func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	//babad

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
		//偶数长度 中间两个必相等
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
