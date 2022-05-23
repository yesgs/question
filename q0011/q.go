package q0011

import "question/common"

//11. 盛最多水的容器
//给定一个长度为 n 的整数数组height。有n条垂线，第 i 条线的两个端点是(i, 0)和(i, height[i])。
//
//找出其中的两条线，使得它们与x轴共同构成的容器可以容纳最多的水。
//
//返回容器可以储存的最大水量。
//
//说明：你不能倾斜容器。

//1,8,6,2,5,4,8,3,7

//面积公式
//s(i,j) =min(h[i],h[j]) * (j-i)
//初始状态，两个分别指针指向一头一尾
//移动任意一个挡板，无论移动最高的还是最低的，底部的宽度都会减少1
//移动挡板有两种选择
//1：假设移动高的挡板
//如果下一块挡板比当前高，参与计算的高度不变，面积肯定变小
//如果下一块挡板比当前低，参与计算的高度不变或者变小，面积肯定变小
//2：假设移动低的挡板
//如果下一块挡板比当前高，参与计算的高度变大，面积肯定变大
//如果下一块挡板比当前低，参与计算的高度变小，面积肯定变小
//综上所述，只有移动短板才有可能是面积变大

func maxArea(height []int) int {

	length := len(height)
	if length == 2 {
		return common.Min(height[0], height[1])
	}

	i, j, area := 0, length-1, 0

	for i < j {
		area = common.Max(area, common.Min(height[i], height[j])*(j-i))
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}

	return area
}
