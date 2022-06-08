package q0026

import (
	"fmt"
	"question/common"
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {

	var inputs []common.InputIntSliceOutputInt
	inputs = append(inputs, common.InputIntSliceOutputInt{Input: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, Output: 5})
	inputs = append(inputs, common.InputIntSliceOutputInt{Input: []int{1, 1, 1, 1, 1, 2, 2, 3, 3}, Output: 3})
	inputs = append(inputs, common.InputIntSliceOutputInt{Input: []int{1, 1, 1, 2, 3, 4}, Output: 4})

	for _, input := range inputs {
		ret := removeDuplicates(input.Input)
		msg := fmt.Sprintf("输入 %v 期望 %v 实际 %v", input.Input, input.Output, ret)
		if !reflect.DeepEqual(ret, input.Output) {
			t.Error("[FAIL]" + msg)
		} else {
			t.Log("[OK]" + msg)
		}
	}

}
