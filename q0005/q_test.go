package q0005

import (
	"question/common"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {

	var inputs []common.InputStringOutputString
	inputs = append(inputs, common.InputStringOutputString{Input: "babad", Output: "bab"})
	inputs = append(inputs, common.InputStringOutputString{Input: "cbbd", Output: "bb"})

	for _, input := range inputs {
		ret := longestPalindrome(input.Input)
		common.AssertEqual(t, input.Input, input.Output, ret)
	}
}
