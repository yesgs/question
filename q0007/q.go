package q0007

//整数反转
//给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
//如果反转后整数超过 32 位的有符号整数的范围[−2^31, 2^31− 1] ，就返回 0。
//示例 1：
//
//输入：x = 123
//输出：321
//示例 2：
//
//输入：x = -123
//输出：-321
//示例 3：
//
//输入：x = 120
//输出：21
//示例 4：
//
//输入：x = 0
//输出：0
//

func reverse(x int) int {
	var flag = 1
	var ret = 0

	if x == 0 {
		return 0
	}
	if x < 0 {
		flag = -1
		x = x * -1
	}

	for x > 0 {
		ret = ret*10 + x%10
		x = x / 10
	}
	if ret > 1<<31-1 {
		return 0
	}
	return ret * flag
}

func q0007Reverse(x int) int {
	return reverse(x)
}
