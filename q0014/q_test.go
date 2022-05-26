package q0014

import (
	"question/common"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {

	var inputs []common.InputStringSliceOutputString
	inputs = append(inputs, common.InputStringSliceOutputString{Input: []string{"flower", "flow", "flight"}, Output: "fl"})
	inputs = append(inputs, common.InputStringSliceOutputString{Input: []string{"flow", "flower", "flight"}, Output: "fl"})
	inputs = append(inputs, common.InputStringSliceOutputString{Input: []string{"dog", "racecar", "car"}, Output: ""})
	inputs = append(inputs, common.InputStringSliceOutputString{Input: []string{}, Output: ""})
	inputs = append(inputs, common.InputStringSliceOutputString{Input: []string{"leetcode"}, Output: "leetcode"})

	for _, input := range inputs {
		ret := longestCommonPrefix(input.Input)
		common.AssertEqual(t, input.Input, input.Output, ret)
	}
}
