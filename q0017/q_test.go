package q0017

import (
	"fmt"
	"question/common"
	"reflect"
	"testing"
)

func TestLetterCombinations(t *testing.T) {

	var inputs []common.InputStringOutputStringSlice
	inputs = append(inputs, common.InputStringOutputStringSlice{Input: "23", Output: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}})
	inputs = append(inputs, common.InputStringOutputStringSlice{Input: "", Output: []string{}})
	inputs = append(inputs, common.InputStringOutputStringSlice{Input: "2", Output: []string{"a", "b", "c"}})

	for _, input := range inputs {
		ret := letterCombinations(input.Input)

		msg := fmt.Sprintf("输入 %v 期望 %v 实际 %v", input.Input, input.Output, ret)

		if len(input.Output) == 0 {
			t.Log("[OK]" + msg)
		} else {

			if !reflect.DeepEqual(ret, input.Output) {
				t.Error("[FAIL]" + msg)
			} else {
				t.Log("[OK]" + msg)
			}
		}

	}
}
