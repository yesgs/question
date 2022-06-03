package q0020

import (
	"question/common"
	"testing"
)

func TestQ0020isValid(t *testing.T) {

	var inputs []common.InputStringOutputBool
	inputs = append(inputs, common.InputStringOutputBool{Input: "()", Output: true})
	inputs = append(inputs, common.InputStringOutputBool{Input: "()[]{}", Output: true})
	inputs = append(inputs, common.InputStringOutputBool{Input: "(]", Output: false})
	inputs = append(inputs, common.InputStringOutputBool{Input: "([)]", Output: false})
	inputs = append(inputs, common.InputStringOutputBool{Input: "{[]}", Output: true})
	inputs = append(inputs, common.InputStringOutputBool{Input: "){", Output: false})

	for _, input := range inputs {
		ret := q0020isValid(input.Input)
		common.AssertEqual(t, input.Input, input.Output, ret)
	}
}
