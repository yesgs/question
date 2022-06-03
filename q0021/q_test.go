package q0021

import (
	"fmt"
	"question/common"
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {

	var inputs []common.Input1IntSliceInput2IntSliceOutputIntSlice
	inputs = append(inputs, common.Input1IntSliceInput2IntSliceOutputIntSlice{Input1: []int{1, 2, 3}, Input2: []int{1, 3, 4}, Output: []int{1, 1, 2, 3, 3, 4}})

	for _, input := range inputs {
		i1 := SliceToListNode(input.Input1)
		i2 := SliceToListNode(input.Input2)
		ret := q0021(i1, i2)
		ret2 := ListNodeToSlice(ret)
		msg := fmt.Sprintf("输入 (%v,%v) 期望 %v 实际 %v", input.Input1, input.Input2, input.Output, ret2)
		if !reflect.DeepEqual(ret2, input.Output) {
			t.Error("[FAIL]" + msg)
		} else {
			t.Log("[OK]" + msg)
		}
	}

}
