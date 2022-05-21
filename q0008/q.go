package q0008

//请你来实现一个myAtoi(string s)函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。
//
//函数myAtoi(string s) 的算法如下：
//
//读入字符串并丢弃无用的前导空格
//检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
//读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
//将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
//如果整数数超过 32 位有符号整数范围 [−2^31, 2^31− 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −2^31 的整数应该被固定为 −2^31 ，大于 2^31− 1 的整数应该被固定为 2^31− 1 。
//返回整数作为最终结果。
//注意：
//
//本题中的空白字符只包括空格字符 ' ' 。
//除前导空格或数字后的其余字符串外，请勿忽略 任何其他字符。

//示例1：
//
//输入：s = "42"
//输出：42
//解释：加粗的字符串为已经读入的字符，插入符号是当前读取的字符。
//第 1 步："42"（当前没有读入字符，因为没有前导空格）
//^
//第 2 步："42"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
//^
//第 3 步："42"（读入 "42"）
//^
//解析得到整数 42 。
//由于 "42" 在范围 [-231, 231 - 1] 内，最终结果为 42 。
//示例2：
//
//输入：s = "   -42"
//输出：-42
//解释：
//第 1 步："   -42"（读入前导空格，但忽视掉）
//^
//第 2 步："   -42"（读入 '-' 字符，所以结果应该是负数）
//^
//第 3 步："   -42"（读入 "42"）
//^
//解析得到整数 -42 。
//由于 "-42" 在范围 [-231, 231 - 1] 内，最终结果为 -42 。
//示例3：
//
//输入：s = "4193 with words"
//输出：4193
//解释：
//第 1 步："4193 with words"（当前没有读入字符，因为没有前导空格）
//^
//第 2 步："4193 with words"（当前没有读入字符，因为这里不存在 '-' 或者 '+'）
//^
//第 3 步："4193 with words"（读入 "4193"；由于下一个字符不是一个数字，所以读入停止）
//^
//解析得到整数 4193 。
//由于 "4193" 在范围 [-231, 231 - 1] 内，最终结果为 4193 。
//

func myAtoi(s string) int {
	maxInt, signAllowed, whitespaceAllowed, sign, digits := int64(2<<30), true, true, 1, []int{}

	for i := 0; i < len(s); i++ {
		c := s[i]
		//过滤前导空格字符
		if c == ' ' && whitespaceAllowed {
			continue
		}
		//判断正负号
		if signAllowed {
			//如果有了正负号，则不允许有空字符
			if c == '+' {
				signAllowed = false
				whitespaceAllowed = false
				continue
			}
			if c == '-' {
				sign = -1
				signAllowed = false
				whitespaceAllowed = false
				continue
			}
		}
		//非数字字符退出
		if c < '0' || c > '9' {
			break
		}
		//如果有了一个数字，则不允许出现空格或空字符
		whitespaceAllowed = false
		signAllowed = false
		digits = append(digits, int(c-48))
	}

	//[0,0,1,2]
	//找出最后一个0的位置
	lastLeading0Index := -1
	for i, d := range digits {
		if d == 0 {
			lastLeading0Index = i
		} else {
			break
		}
	}
	//删除前导0
	if lastLeading0Index > -1 {
		digits = digits[lastLeading0Index+1:]
	}

	var rtnMax int64
	if sign > 0 {
		//正数
		rtnMax = maxInt - 1
	} else {
		//负数
		rtnMax = maxInt
	}

	var num, place int64

	place, num = 1, 0

	//[4,2]
	digitsCount := len(digits)
	if digitsCount > 10 {
		return int(int64(sign) * rtnMax)
	}
	for i := digitsCount - 1; i >= 0; i-- {
		num += int64(digits[i]) * place
		place *= 10
		if num > rtnMax {
			return int(int64(sign) * rtnMax)
		}
	}

	num *= int64(sign)
	return int(num)
}
