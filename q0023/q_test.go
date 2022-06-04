package q0023

import (
	"fmt"
	"question/common"
	"reflect"
	"testing"
)

func TestSwapPairs(t *testing.T) {

	var inputs []common.InputIntSliceOutputIntSlice
	inputs = append(inputs, common.InputIntSliceOutputIntSlice{Input: []int{1, 2, 3, 4}, Output: []int{2, 1, 4, 3}})
	inputs = append(inputs, common.InputIntSliceOutputIntSlice{Input: []int{}, Output: []int{}})
	inputs = append(inputs, common.InputIntSliceOutputIntSlice{Input: []int{1}, Output: []int{1}})

	for _, input := range inputs {
		i1 := SliceToListNode(input.Input)
		ret := swapPairs(i1)
		ret2 := ListNodeToSlice(ret)
		msg := fmt.Sprintf("输入 %v 期望 %v 实际 %v", input.Input, input.Output, ret2)
		if !reflect.DeepEqual(ret2, input.Output) {
			t.Error("[FAIL]" + msg)
		} else {
			t.Log("[OK]" + msg)
		}
	}

}
