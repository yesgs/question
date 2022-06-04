package q0022

import (
	"fmt"
	"question/common"
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {

	var inputs []common.InputIntOutputStringSlice
	inputs = append(inputs, common.InputIntOutputStringSlice{Input: 3, Output: []string{"((()))", "(()())", "(())()", "()(())", "()()()"}})

	for _, input := range inputs {
		ret := generateParenthesis(input.Input)
		msg := fmt.Sprintf("输入 %v 期望 %v 实际 %v", input.Input, input.Output, ret)
		if len(input.Output) == 0 && len(ret) == 0 {
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
