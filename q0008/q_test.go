package q0008

import (
	"question/common"
	"testing"
)

func TestMyAtoi(t *testing.T) {

	var inputs []common.InputStringOutputInt
	inputs = append(inputs, common.InputStringOutputInt{Input: "42", Output: 42})
	inputs = append(inputs, common.InputStringOutputInt{Input: "   -42", Output: -42})
	inputs = append(inputs, common.InputStringOutputInt{Input: "4193 with words", Output: 4193})

	for _, input := range inputs {
		ret := myAtoi(input.Input)
		common.AssertEqual(t, input.Input, input.Output, ret)
	}

}
