package q0015

import (
	"sort"
)

//三数之和为0

//给你一个包含 n 个整数的数组nums，判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。
//
//注意：答案中不可以包含重复的三元组。
//
//
//示例 1：
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]

//固定 3 个指针中最左（最小）数字的指针 k
//双指针 i，j 分设在数组索引 (k, len(nums)) 两端
//通过双指针交替向中间移动，记录对于每个固定指针 k 的所有满足 nums[k] + nums[i] + nums[j] == 0 的 i,j 组合

func threeSum(nums []int) [][]int {
	var ret [][]int
	var s = 0
	i, j, k := 0, 0, 0

	//先排序
	sort.Ints(nums)

	//-4 -1 -1 0 1 2
	for k = 0; k < len(nums); k++ {
		//当 nums[k] > 0 时直接break跳出
		//因为 nums[j] >= nums[i] >= nums[k] > 0
		//即 3 个数字都大于 0 ，在此固定指针 k 之后不可能再找到结果了
		if nums[k] > 0 {
			break
		}

		//当 k > 0且nums[k] == nums[k - 1]时即跳过此元素nums[k]
		//因为已经将 nums[k - 1] 的所有组合加入到结果中，本次双指针搜索只会得到重复组合。
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		//i，j 分设在数组索引 (k, len(nums)) 两端
		//当i < j时循环计算s = nums[k] + nums[i] + nums[j]
		//按照以下规则执行双指针移动：
		//当s < 0时，i = i + 1并跳过所有重复的nums[i]；
		//当s > 0时，j = j - 1并跳过所有重复的nums[j]；
		//当s == 0时，记录组合[k, i, j]至res，执行i += 1和j -= 1并跳过所有重复的nums[i]和nums[j]，防止记录到重复组合。

		i = k + 1
		j = len(nums) - 1
		//双指针循环开始
		for i < j {
			s = nums[k] + nums[i] + nums[j]
			if s < 0 {
				//总和小了 最小的数字稍微调大 i往后挪一个
				i += 1
				//挪的时候判断数字是否重复，重复的话顺延一个
				for i < j && nums[i] == nums[i-1] {
					i += 1
				}
			} else if s > 0 {
				//总和大了 最大的数字稍微调小 j往前挪一个
				j -= 1
				//挪的时候判断数字是否重复，重复的话顺延一个
				for j > i && nums[j] == nums[j+1] {
					j -= 1
				}
			} else {
				//匹配到一个合格的组合 加入结果集合
				ret = append(ret, []int{nums[k], nums[i], nums[j]})
				//计算下一组可能的情况
				i += 1
				for i < j && nums[i] == nums[i-1] {
					i += 1
				}
				j -= 1
				for j > i && nums[j] == nums[j+1] {
					j -= 1
				}
			}
		}
	}

	return ret
}
