package q0003

import (
	"question/common"
	"testing"
)

func TestLengthOfLongestSubstring(t *testing.T) {

	var inputs []common.InputStringOutputInt
	inputs = append(inputs, common.InputStringOutputInt{Input: "asd", Output: 3})
	inputs = append(inputs, common.InputStringOutputInt{Input: "abba", Output: 2})
	inputs = append(inputs, common.InputStringOutputInt{Input: "abcabcbb", Output: 3})
	inputs = append(inputs, common.InputStringOutputInt{Input: "pwwkew", Output: 3})
	inputs = append(inputs, common.InputStringOutputInt{Input: "dvdf", Output: 3})

	for _, input := range inputs {
		ret := lengthOfLongestSubstring(input.Input)
		common.AssertEqual(t, input.Input, input.Output, ret)
	}
}
