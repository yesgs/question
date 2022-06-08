package q0026

func removeDuplicates2(nums []int) int {

	length := len(nums)
	if length == 0 {
		return 0
	}
	i := 0
	for j := 1; j < length; j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}

	return i + 1
}

func removeDuplicates(nums []int) int {
	// 1 1 1 2 3 4
	length := len(nums)
	i := 0
	//i 慢慢走
	//j 从i后面开始遍历数组，j单调递增且不会回拨
	for j := 1; j < length; j++ {
		//找到第一个和i不相等的值
		if nums[j] != nums[i] {
			//将第一个和i 不相等的值赋值给 i 一个后面的位置
			nums[i+1] = nums[j]

			//i往后走一个
			i++
		}
	}
	return i + 1
}
