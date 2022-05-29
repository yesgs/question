package q0016

import (
	"fmt"
	"question/common"
	"reflect"
	"testing"
)

func TestThreeSumClosest(t *testing.T) {

	var inputs []common.Input1IntSliceInput2IntOutputInt
	inputs = append(inputs, common.Input1IntSliceInput2IntOutputInt{Input1: []int{-1, 2, 1, -4}, Input2: 1, Output: 2})

	for _, input := range inputs {
		ret := threeSumClosest(input.Input1, input.Input2)

		msg := fmt.Sprintf("输入 %v 期望 %v 实际 %v", fmt.Sprintf("(%v,%v)", input.Input1, input.Input2), input.Output, ret)

		if !reflect.DeepEqual(ret, input.Output) {
			t.Error("[FAIL]" + msg)
		} else {
			t.Log("[OK]" + msg)
		}

	}
}
