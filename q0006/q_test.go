package q0006

import (
	"fmt"
	"question/common"
	"testing"
)

func TestQ0006Convert(t *testing.T) {

	var inputs []common.Input1StringInput2IntOutputString
	inputs = append(inputs, common.Input1StringInput2IntOutputString{Input1: "PAYPALISHIRING", Input2: 3, Output: "PAHNAPLSIIGYIR"})
	inputs = append(inputs, common.Input1StringInput2IntOutputString{Input1: "PAYPALISHIRING", Input2: 4, Output: "PINALSIGYAHRPI"})
	inputs = append(inputs, common.Input1StringInput2IntOutputString{Input1: "A", Input2: 1, Output: "A"})
	inputs = append(inputs, common.Input1StringInput2IntOutputString{Input1: "LEETCODE", Input2: 3, Output: "LCETOEED"})
	inputs = append(inputs, common.Input1StringInput2IntOutputString{Input1: "LALALADEMAXIYA", Input2: 4, Output: "LDYAAEIALLMXAA"})

	for _, input := range inputs {
		ret := q0006Convert(input.Input1, input.Input2)
		msg := fmt.Sprintf("输入 %v 期望 %v 实际 %v", fmt.Sprintf("(%v,%v)", input.Input1, input.Input2), input.Output, ret)
		if ret != input.Output {
			t.Error(msg)
		} else {
			t.Log(msg)
		}
	}
}
