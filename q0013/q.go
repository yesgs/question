package q0013

//给定一个罗马数字，将其转换成整数。
//示例1:
//
//输入:s = "III"
//输出: 3
//示例2:
//
//输入:s = "IV"
//输出: 4
//示例3:
//
//输入:s = "IX"
//输出: 9
//示例4:
//
//输入:s = "LVIII"
//输出: 58
//解释: L = 50, V= 5, III = 3.
//示例5:
//
//输入:s = "MCMXCIV"
//输出: 1994
//解释: M = 1000, CM = 900, XC = 90, IV = 4.

func romanToInt(s string) int {
	var valueMap = make(map[uint8]int)
	valueMap['I'] = 1
	valueMap['V'] = 5
	valueMap['X'] = 10
	valueMap['L'] = 50
	valueMap['C'] = 100
	valueMap['D'] = 500
	valueMap['M'] = 1000

	//罗马数字由 I,V,X,L,C,D,M 构成；
	//当小值在大值的左边，则减小值，如 IV=5-1=4；
	//当小值在大值的右边，则加小值，如 VI=5+1=6；
	//由上可知，右值永远为正，因此最后一位必然为正。

	//XXVII
	sum := 0
	preNum := valueMap[s[0]]
	for i := 1; i < len(s); i++ {
		num := valueMap[s[i]]
		if preNum < num {
			sum = sum - preNum
		} else {
			sum = sum + preNum
		}
		preNum = num
	}
	//退出循环时，还剩最后一位没有加

	sum += preNum
	return sum
}
