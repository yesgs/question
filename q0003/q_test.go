package q0003

import (
	"fmt"
	"testing"
)

type stu struct {
	Q string
	A int
}

func TestLengthOfLongestSubstring(t *testing.T) {

	var inputs = []stu{}
	inputs = append(inputs, stu{Q: "asd", A: 3})
	inputs = append(inputs, stu{Q: "abba", A: 2})
	inputs = append(inputs, stu{Q: "abcabcbb", A: 3})
	inputs = append(inputs, stu{Q: "pwwkew", A: 3})
	inputs = append(inputs, stu{Q: "dvdf", A: 3})

	for _, input := range inputs {
		ret := lengthOfLongestSubstring(input.Q)
		msg := fmt.Sprintf("%v want %v get %v", input.Q, input.A, ret)

		if ret != input.A {
			t.Error(msg)
		} else {
			t.Log(msg)
		}
	}
}
