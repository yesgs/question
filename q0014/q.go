package q0014

//最长公共前缀
//
//编写一个函数来查找字符串数组中的最长公共前缀。
//
//如果不存在公共前缀，返回空字符串""。
//
//
//示例 1：
//
//输入：strs = ["flower","flow","flight"]
//输出："fl"
//示例 2：
//
//输入：strs = ["dog","racecar","car"]
//输出：""
//解释：输入不存在公共前缀。
//

func longestCommonPrefix(strs []string) string {
	ret := ""
	if len(strs) == 0 {
		return ret
	}

	//遍历第一个字符串的长度
	for i := 0; i < len(strs[0]); i++ {
		//遍历第一个字符串之后的所有字符串
		for j := 1; j < len(strs); j++ {

			//如果第一个字符串长度超过第j个字符串的长度
			if i == len(strs[j]) {
				//直接退出
				return strs[0][:i]
			}

			//如果第一个字符串的第i个字符不等于第j个字符串的第i个字符
			if strs[0][i] != strs[j][i] {
				return strs[0][:i]
			}
		}
	}
	//执行到这里说明第一个字符串是最短的
	return strs[0]
}
