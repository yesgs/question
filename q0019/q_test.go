package q0019

import (
	"fmt"
	"question/common"
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {

	var inputs []common.Input1IntSliceInput2IntOutputIntSlice
	inputs = append(inputs, common.Input1IntSliceInput2IntOutputIntSlice{Input1: []int{1, 2, 3, 4, 5}, Input2: 2, Output: []int{1, 2, 3, 5}})
	inputs = append(inputs, common.Input1IntSliceInput2IntOutputIntSlice{Input1: []int{1}, Input2: 1, Output: []int{}})
	inputs = append(inputs, common.Input1IntSliceInput2IntOutputIntSlice{Input1: []int{1, 2}, Input2: 1, Output: []int{1}})
	inputs = append(inputs, common.Input1IntSliceInput2IntOutputIntSlice{Input1: []int{1, 2}, Input2: 2, Output: []int{2}})

	for _, input := range inputs {
		i1 := SliceToListNode(input.Input1)
		ret := removeNthFromEnd(i1, input.Input2)
		ret2 := ListNodeToSlice(ret)
		msg := fmt.Sprintf("输入 (%v,%v) 期望 %v 实际 %v", input.Input1, input.Input2, input.Output, ret2)
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
