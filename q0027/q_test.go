package q0027

import (
	"fmt"
	"question/common"
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {

	var inputs []common.Input1IntSliceInput2IntOutputInt
	inputs = append(inputs, common.Input1IntSliceInput2IntOutputInt{Input1: []int{3, 2, 2, 3}, Input2: 3, Output: 2})
	inputs = append(inputs, common.Input1IntSliceInput2IntOutputInt{Input1: []int{0, 1, 2, 2, 3, 0, 4, 2}, Input2: 2, Output: 5})
	inputs = append(inputs, common.Input1IntSliceInput2IntOutputInt{Input1: []int{2, 3, 3}, Input2: 2, Output: 2})
	inputs = append(inputs, common.Input1IntSliceInput2IntOutputInt{Input1: []int{2, 3, 3}, Input2: 3, Output: 1})
	inputs = append(inputs, common.Input1IntSliceInput2IntOutputInt{Input1: []int{1}, Input2: 1, Output: 0})
	inputs = append(inputs, common.Input1IntSliceInput2IntOutputInt{Input1: []int{3, 3}, Input2: 3, Output: 0})
	inputs = append(inputs, common.Input1IntSliceInput2IntOutputInt{Input1: []int{4, 5}, Input2: 4, Output: 1})

	for _, input := range inputs {
		ret := removeElement(input.Input1, input.Input2)
		msg := fmt.Sprintf("输入 (%v,%v) 期望 %v 实际 %v", input.Input1, input.Input2, input.Output, ret)
		if !reflect.DeepEqual(ret, input.Output) {
			t.Error("[FAIL]" + msg)
		} else {
			t.Log("[OK]" + msg)
		}
	}

}
