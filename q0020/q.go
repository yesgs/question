package q0020

//
//有效的括号
//
//给定一个只包括 '('，')'，'{'，'}'，'['，']'的字符串 s ，判断字符串是否有效。
//
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。

func isValid(s string) bool {
	length := len(s)
	if length%2 != 0 {
		return false
	}

	var ret = make([]uint8, 0, length)
	for i := 0; i < length; i++ {
		switch s[i] {
		case '(':
			fallthrough
		case '{':
			fallthrough
		case '[':
			ret = append(ret, s[i])
		case ')':
			if len(ret) == 0 {
				return false
			}
			if ret[len(ret)-1] != '(' {
				return false
			} else {
				ret = ret[:len(ret)-1]
			}
		case '}':
			if len(ret) == 0 {
				return false
			}
			if ret[len(ret)-1] != '{' {
				return false
			} else {
				ret = ret[:len(ret)-1]
			}
		case ']':
			if len(ret) == 0 {
				return false
			}
			if ret[len(ret)-1] != '[' {
				return false
			} else {
				ret = ret[:len(ret)-1]
			}
		}
	}
	return len(ret) == 0
}

func q0020isValid(s string) bool {
	return isValid(s)
}
