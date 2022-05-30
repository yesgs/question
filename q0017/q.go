package q0017

//电话号码的字母组合
//给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合。
//答案可以按 任意顺序 返回。
//给出数字到字母的映射如下（与电话按键相同）。
//注意 1 不对应任何字母。

var letterMap = []string{
	" ",    //0
	"",     //1
	"abc",  //2
	"def",  //3
	"ghi",  //4
	"jkl",  //5
	"mno",  //6
	"pqrs", //7
	"tuv",  //8
	"wxyz", //9
}
var res []string

//输入：digits = "23"
//输出：["ad","ae","af","bd","be","bf","cd","ce","cf"]

func letterCombinations(digits string) []string {
	res = []string{} //测试的时候保证每个case初始化的res为空

	if digits == "" {
		return res
	}

	//初始状态 求深度为0 空字符串为前缀的解
	dfs(digits, 0, "")
	return res
}

func dfs(digits string, depth int, resStr string) {
	//输入的字符串 2 3 4
	//如果当前深度已经是最深了
	//把res 添加的结果集
	if depth == len(digits) {
		res = append(res, resStr)
		return
	}

	//输入的数字
	num := digits[depth] // => 2
	//数字对应的字符串
	letter := letterMap[num-'0'] // => a b c

	//遍历深度为depth的字符串
	//以当前某个字符为前缀，求下一级的解
	for i := 0; i < len(letter); i++ {
		dfs(digits, depth+1, resStr+string(letter[i]))
	}
	return
}

//输入 2 3 4
//遍历 a b c
//求 a 为前缀的解
//	遍历 d e f
//	求 ad 为前缀的解
//		遍历 g h i
//		求 adg 为前缀的解
//		求 adh 为前缀的解
//		求 adi 为前缀的解

//	求 ae 为前缀的解
//		遍历 g h i
//		求 aeg 为前缀的解
//		求 aeh 为前缀的解
//		求 aei 为前缀的解

//	求 af 为前缀的解

//求 b 为前缀的解
//	遍历 d e f
//	求 bd 为前缀的解
//	求 be 为前缀的解
//	求 bf 为前缀的解

//求 c 为前缀的解
//	遍历 d e f
//	求 cd 为前缀的解
//	求 ce 为前缀的解
//	求 cf 为前缀的解
