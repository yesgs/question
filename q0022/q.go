package q0022

//重点关注

//括号生成
//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。

//示例 1：
//
//输入：n = 3
//输出：["((()))","(()())","(())()","()(())","()()()"]
var res []string

func generateParenthesis(n int) []string {
	res = []string{}

	if n == 0 {
		return res
	}

	dfs("", 0, 0, n)

	var res2 []string
	dfs2(n, n, "", &res2)

	return res
}

//n 左括号、右括号一共得用几对
//leftCount 是左括号数量
//rightCount 是右括号数
//curStr 当前递归得到的结果

func dfs(curStr string, leftCount int, rightCount int, n int) {

	//退出条件
	//左右括号数成对
	if leftCount == n && rightCount == n {
		//递归边界
		res = append(res, curStr)
		return
	}

	// 剪枝
	//左括号 比 右括号 少
	if leftCount < rightCount {
		return
	}

	//如果左括号个数没达到要求
	//就在当前的字符串上拼左括号
	if leftCount < n {
		dfs(curStr+"(", leftCount+1, rightCount, n)
	}

	//如果右括号个数没有达到要求 且 左括号数一定要大于右括号数
	//就在当前的字符串上拼右括号
	if rightCount < n && leftCount > rightCount {
		dfs(curStr+")", leftCount, rightCount+1, n)
	}
}

func dfs2(leftIndex, rightIndex int, str string, res *[]string) {
	if leftIndex == 0 && rightIndex == 0 {
		*res = append(*res, str)
		return
	}
	if leftIndex > 0 {
		dfs2(leftIndex-1, rightIndex, str+"(", res)
	}
	if rightIndex > 0 && leftIndex < rightIndex {
		dfs2(leftIndex, rightIndex-1, str+")", res)
	}
}
