package q0012

import (
	"question/common"
	"testing"
)

func TestIntToRoman(t *testing.T) {

	var inputs []common.InputIntOutputString
	inputs = append(inputs, common.InputIntOutputString{Input: 3, Output: "III"})
	inputs = append(inputs, common.InputIntOutputString{Input: 4, Output: "IV"})
	inputs = append(inputs, common.InputIntOutputString{Input: 9, Output: "IX"})
	inputs = append(inputs, common.InputIntOutputString{Input: 58, Output: "LVIII"})
	inputs = append(inputs, common.InputIntOutputString{Input: 1994, Output: "MCMXCIV"})

	for _, input := range inputs {
		ret := intToRoman(input.Input)
		common.AssertEqual(t, input.Input, input.Output, ret)
	}
}
