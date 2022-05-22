package q0009

import (
	"question/common"
	"testing"
)

func TestIsPalindrome(t *testing.T) {

	var inputs []common.InputIntOutputBool
	inputs = append(inputs, common.InputIntOutputBool{Input: 121, Output: true})
	inputs = append(inputs, common.InputIntOutputBool{Input: -121, Output: false})
	inputs = append(inputs, common.InputIntOutputBool{Input: 10, Output: false})

	for _, input := range inputs {
		ret := isPalindrome(input.Input)
		common.AssertEqual(t, input.Input, input.Output, ret)
	}
}
