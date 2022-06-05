package q0023

import (
	"fmt"
	"question/common"
	"reflect"
	"testing"
)

func TestMergeKLists(t *testing.T) {

	var inputs []common.InputIntSliceSliceOutputIntSlice
	inputs = append(inputs, common.InputIntSliceSliceOutputIntSlice{Input: [][]int{{1, 4, 5}, {1, 3, 4}, {2, 6}}, Output: []int{1, 1, 2, 3, 4, 4, 5, 6}})
	inputs = append(inputs, common.InputIntSliceSliceOutputIntSlice{Input: [][]int{}, Output: []int{}})
	inputs = append(inputs, common.InputIntSliceSliceOutputIntSlice{Input: [][]int{{}}, Output: []int{}})

	for _, input := range inputs {
		i1 := SliceToListNodeSlice(input.Input)
		//ret := mergeKLists(i1)
		ret := mergeKLists2(i1)
		ret2 := ListNodeToSlice(ret)
		msg := fmt.Sprintf("输入 %v 期望 %v 实际 %v", input.Input, input.Output, ret2)
		if len(input.Output) == 0 && len(ret2) == 0 {
			t.Log("[OK]" + msg)
		} else {
			if !reflect.DeepEqual(ret2, input.Output) {
				t.Error("[FAIL]" + msg)
			} else {
				t.Log("[OK]" + msg)
			}
		}
	}

}
