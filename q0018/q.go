package q0018

import (
	"sort"
)

//四数之和

//给你一个由 n 个整数组成的数组nums ，和一个目标值 target 。
//请你找出并返回满足下述全部条件且不重复的四元组[nums[a], nums[b], nums[c], nums[d]]
//（若两个四元组元素一一对应，则认为两个四元组重复）：

//0 <= a, b, c, d< n
//a、b、c 和 d 互不相同
//nums[a] + nums[b] + nums[c] + nums[d] == target

//示例 1：

//输入：nums = [1,0,-1,0,-2,2], target = 0
//输出：[[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]

func fourSum(nums []int, target int) [][]int {
	var ret [][]int

	if len(nums) < 4 {
		return [][]int{}
	}

	sort.Ints(nums)

	i, j, l, r, t, length := 0, 0, 0, 0, 0, len(nums)

	for i = 0; i < length-3; i++ {
		//当 nums[i] 的值与前面的值相等时忽略
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		//获取当前最小值,如果最小值比目标值大,说明后面越来越大的值根本没戏
		if nums[i]+nums[i+1]+nums[i+2]+nums[i+3] > target {
			//这里使用的break,直接退出此次循环,因为后面的数只会更大
			break
		}
		//获取当前最大值,如果最大值比目标值小,说明后面越来越小的值根本没戏,忽略
		if nums[i]+nums[length-3]+nums[length-2]+nums[length-1] < target {
			//这里使用continue,继续下一次循环,因为下一次循环有更大的数
			continue
		}
		for j = i + 1; j < length-2; j++ {

			if j > i+1 && nums[j] == nums[j-1] {
				//当j的值与前面的值相等时忽略
				continue
			}
			//最小的数加起来比目标大，退出整个循环，因为后面只会更大
			if nums[i]+nums[j]+nums[j+1]+nums[j+2] > target {
				break
			}
			//最大的数加起来比目标小，退出本次循环，因为后面有更大的
			if nums[i]+nums[j]+nums[length-2]+nums[length-1] < target {
				continue
			}
			l, r = j+1, length-1
			//双指针遍历
			//如果等于目标值,left++并去重,right--并去重,
			//当前和大于目标值时right--
			//当当前和小于目标值时left++
			for l < r {
				t = nums[i] + nums[j] + nums[l] + nums[r]

				if t > target {
					//和超过目标了，把最大的数往小了调
					r--
					//调的时候跳过重复的数
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				} else if t < target {
					//和不到目标，把最小的数往大了调
					l++
					//调的时候跳过重复的数
					for l < r && nums[l] == nums[l-1] {
						l++
					}

				} else {
					ret = append(ret, []int{nums[i], nums[j], nums[l], nums[r]})
					l++
					for l < r && nums[l] == nums[l-1] {
						l++
					}
					r--
					for l < r && nums[r] == nums[r+1] {
						r--
					}
				}

			}

		}
	}

	return ret
}
