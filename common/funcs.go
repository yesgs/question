package common

import (
	"fmt"
	"testing"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func AssertEqual(t *testing.T, input, output, ret interface{}) {
	msg := fmt.Sprintf("输入 %v 期望 %v 实际 %v", input, output, ret)
	if ret != input {
		t.Error(msg)
	} else {
		t.Log(msg)
	}
}
