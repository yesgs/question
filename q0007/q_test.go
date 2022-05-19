package q0007

import (
	"question/common"
	"testing"
)

func TestQ0007Reverse(t *testing.T) {

	var inputs []common.InputIntOutputInt
	inputs = append(inputs, common.InputIntOutputInt{Input: 123, Output: 321})
	inputs = append(inputs, common.InputIntOutputInt{Input: -123, Output: -321})
	inputs = append(inputs, common.InputIntOutputInt{Input: 120, Output: 21})
	inputs = append(inputs, common.InputIntOutputInt{Input: 0, Output: 0})
	inputs = append(inputs, common.InputIntOutputInt{Input: 1534236469, Output: 0})
	inputs = append(inputs, common.InputIntOutputInt{Input: 1563847412, Output: 0})

	for _, input := range inputs {
		ret := q0007Reverse(input.Input)
		common.AssertEqual(t, input.Input, input.Output, ret)
	}
}
