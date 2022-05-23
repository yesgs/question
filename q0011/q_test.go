package q0011

import (
	"question/common"
	"testing"
)

func TestMaxArea(t *testing.T) {

	var inputs []common.InputIntSliceOutputInt
	inputs = append(inputs, common.InputIntSliceOutputInt{Input: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}, Output: 49})
	inputs = append(inputs, common.InputIntSliceOutputInt{Input: []int{1, 1}, Output: 1})
	inputs = append(inputs, common.InputIntSliceOutputInt{Input: []int{1, 2}, Output: 1})

	for _, input := range inputs {
		ret := maxArea(input.Input)
		common.AssertEqual(t, input.Input, input.Output, ret)
	}
}
