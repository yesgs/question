package q0027

//27. 移除元素

//给你一个数组 nums和一个值 val，
//你需要 原地 移除所有数值等于val的元素，并返回移除后数组的新长度。
//不要使用额外的数组空间，
//你必须仅使用 O(1) 额外空间并 原地 修改输入数组。
//元素的顺序可以改变。你不需要考虑数组中超出新长度后面的元素。

//输入：nums = [3,2,2,3], val = 3
//输出：2, nums = [2,2]

func removeElement(nums []int, val int) int {
	//log.Println("start", nums)

	length := len(nums)
	i := 0
	j := 0
	for {
		if j >= length {
			break
		}
		//0, 1, 2, 2, 3, 0, 4, 2  => val=2

		//遍历数组每个元素
		//如果元素不等于指定的val，i,j 齐头并进
		//当发现元素等于指定val时，i不动，j往前走
		//直到再次找到不等于指定val时，重新搬运值

		if nums[j] != val {
			nums[i] = nums[j]
			i++
		}

		j++
	}
	return i
}
