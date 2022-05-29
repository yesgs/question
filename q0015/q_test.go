package q0015

import (
	"fmt"
	"log"
	"question/common"
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {

	var inputs []common.InputIntSliceOutputIntSliceSlice
	inputs = append(inputs, common.InputIntSliceOutputIntSliceSlice{Input: []int{-1, 0, 1, 2, -1, -4}, Output: [][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}}})
	inputs = append(inputs, common.InputIntSliceOutputIntSliceSlice{Input: []int{}, Output: [][]int{}})
	inputs = append(inputs, common.InputIntSliceOutputIntSliceSlice{Input: []int{0}, Output: [][]int{}})

	for _, input := range inputs {
		ret := threeSum(input.Input)
		log.Println(ret)

		msg := fmt.Sprintf("输入 %v 期望 %v 实际 %v", input.Input, input.Output, ret)

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
