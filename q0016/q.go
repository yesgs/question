package q0016

import (
	"math"
	"sort"
)

//最接近的三数之和
//给你一个长度为 n 的整数数组 nums 和 一个目标值 target。请你从 nums 中选出三个整数，使它们的和与 target 最接近。
//
//返回这三个数的和。
//
//假定每组输入只存在恰好一个解。

//示例 1：
//
//输入：nums = [-1,2,1,-4], target = 1
//输出：2
//解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。
//示例 2：
//
//输入：nums = [0,0,0], target = 1
//输出：0
//

func threeSumClosest(nums []int, target int) int {
	var ret = math.MaxInt
	s, k, i, j := 0, 0, 0, 0

	//先排序
	sort.Ints(nums)

	//-4 -1 1 2
	//固定一个指针
	for k = 0; k < len(nums); k++ {

		//i,j 分别指向一头一尾
		i = k + 1
		j = len(nums) - 1
		//双指针循环开始
		for i < j {
			//求和
			s = nums[k] + nums[i] + nums[j]
			//ret取差值小的那个
			if math.Abs(float64(target-s)) < math.Abs(float64(target-ret)) {
				ret = s
			}
			if s > target {
				//和大于目标值 挪高位指针
				j--
			} else if s < target {
				//和小于目标值 挪低位指针
				i++
			} else {
				//如果正好和等于目标值直接返回
				return ret
			}
		}
	}

	return ret
}
