package q0018

import (
	"fmt"
	"question/common"
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {

	var inputs []common.Input1IntSliceInput2IntOutputIntSliceSlice
	inputs = append(inputs, common.Input1IntSliceInput2IntOutputIntSliceSlice{Input1: []int{1, 0, -1, 0, -2, 2}, Input2: 0, Output: [][]int{[]int{-2, -1, 1, 2}, []int{-2, 0, 0, 2}, []int{-1, 0, 0, 1}}})
	inputs = append(inputs, common.Input1IntSliceInput2IntOutputIntSliceSlice{Input1: []int{2, 2, 2, 2, 2}, Input2: 8, Output: [][]int{[]int{2, 2, 2, 2}}})
	inputs = append(inputs, common.Input1IntSliceInput2IntOutputIntSliceSlice{Input1: []int{-2, -1, -1, 1, 1, 2, 2}, Input2: 0, Output: [][]int{[]int{-2, -1, 1, 2}, []int{-1, -1, 1, 1}}})

	for _, input := range inputs {
		ret := fourSum(input.Input1, input.Input2)

		msg := fmt.Sprintf("输入 (%v, %v) 期望 %v 实际 %v", input.Input1, input.Input2, input.Output, ret)

		if len(input.Output) == 0 {
			t.Log("[OK]" + msg)
		} else {

			if !reflect.DeepEqual(ret, input.Output) {
				t.Error("[FAIL]" + msg)
			} else {
				t.Log("[OK]" + msg)
			}
		}

	}
}
