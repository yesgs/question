package q0028

import (
	"fmt"
	"question/common"
	"testing"
)

func TestStrStr(t *testing.T) {

	var inputs []common.Input1StringInput2StringOutputInt
	//inputs = append(inputs, common.Input1StringInput2StringOutputInt{Input1: "hello", Input2: "ll", Output: 2})
	//inputs = append(inputs, common.Input1StringInput2StringOutputInt{Input1: "aaaaa", Input2: "bba", Output: -1})
	//inputs = append(inputs, common.Input1StringInput2StringOutputInt{Input1: "aaa", Input2: "aaaa", Output: -1})
	//inputs = append(inputs, common.Input1StringInput2StringOutputInt{Input1: "mississippi", Input2: "mississippi", Output: 0})
	//inputs = append(inputs, common.Input1StringInput2StringOutputInt{Input1: "mississippi", Input2: "issi", Output: 1})
	inputs = append(inputs, common.Input1StringInput2StringOutputInt{Input1: "mississippi", Input2: "issipi", Output: -1})

	for _, input := range inputs {
		ret := strStr(input.Input1, input.Input2)
		common.AssertEqual(t, fmt.Sprintf("(%v,%v)", input.Input1, input.Input2), input.Output, ret)
	}
}
