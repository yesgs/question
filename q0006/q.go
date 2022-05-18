package q0006

import (
	"question/common"
	"strings"
)

//Z 字形变换
//将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行Z 字形排列。
//
//比如输入字符串为 "PAYPALISHIRING"行数为 3 时，排列如下：
//
//P   A   H   N
//A P L S I I G
//Y   I   R
//之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
//
//请你实现这个将字符串进行指定行数变换的函数：
//
//string convert(string s, int numRows);
//

//示例 1：
//
//输入：s = "PAYPALISHIRING", numRows = 3
//输出："PAHNAPLSIIGYIR"
//示例 2：
//输入：s = "PAYPALISHIRING", numRows = 4
//输出："PINALSIGYAHRPI"
//解释：
//P     I    N
//A   L S  I G
//Y A   H R
//P     I
//示例 3：
//
//输入：s = "A", numRows = 1
//输出："A"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	//字符串长为1,行数为2，显然不可能，即行数输入的有问题
	numRows = common.Min(len(s), numRows)
	var rows []string = make([]string, numRows) //按照行数初始化

	loc := 0      //行号
	down := false //字符串遍历方向
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		rows[loc] += c

		//处于第一行 往下走
		//处于最后一行 往上走
		if loc == 0 || loc == numRows-1 {
			down = !down
		}
		//如果要向下走 行号进1
		if down {
			loc = loc + 1
		} else {
			//向上走 行号退1
			loc = loc - 1
		}
	}

	return strings.Join(rows, "")
}

func q0006Convert(s string, numRows int) string {
	return convert(s, numRows)
}
