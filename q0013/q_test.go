package q0013

import (
	"question/common"
	"testing"
)

func TestRomanToInt(t *testing.T) {

	var inputs []common.InputStringOutputInt
	inputs = append(inputs, common.InputStringOutputInt{Input: "III", Output: 3})
	inputs = append(inputs, common.InputStringOutputInt{Input: "IV", Output: 4})
	inputs = append(inputs, common.InputStringOutputInt{Input: "IX", Output: 9})
	inputs = append(inputs, common.InputStringOutputInt{Input: "LVIII", Output: 58})
	inputs = append(inputs, common.InputStringOutputInt{Input: "MCMXCIV", Output: 1994})

	for _, input := range inputs {
		ret := romanToInt(input.Input)
		common.AssertEqual(t, input.Input, input.Output, ret)
	}
}
