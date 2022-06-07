package q0025

import (
	"fmt"
	"question/common"
	"reflect"
	"testing"
)

func TestReverseKGroup(t *testing.T) {

	var inputs []common.Input1IntSliceInput2IntOutputIntSlice
	//inputs = append(inputs, common.Input1IntSliceInput2IntOutputIntSlice{Input1: []int{1, 2, 3, 4, 5}, Input2: 2, Output: []int{2, 1, 4, 3, 5}})
	//inputs = append(inputs, common.Input1IntSliceInput2IntOutputIntSlice{Input1: []int{1, 2, 3, 4, 5}, Input2: 3, Output: []int{3, 2, 1, 4, 5}})
	inputs = append(inputs, common.Input1IntSliceInput2IntOutputIntSlice{Input1: []int{1, 2, 3, 4, 5, 6, 7, 8}, Input2: 3, Output: []int{3, 2, 1, 6, 5, 4, 7, 8}})

	for _, input := range inputs {
		i1 := common.SliceToListNode(input.Input1)
		//ret := reverseKGroup(i1, input.Input2)
		ret := reverseKGroup2(i1, input.Input2)
		ret2 := common.ListNodeToSlice(ret)
		msg := fmt.Sprintf("输入 (%v,%v) 期望 %v 实际 %v", input.Input1, input.Input2, input.Output, ret2)
		if !reflect.DeepEqual(ret2, input.Output) {
			t.Error("[FAIL]" + msg)
		} else {
			t.Log("[OK]" + msg)
		}
	}
}
