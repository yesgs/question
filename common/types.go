package common

type InputIntOutputInt struct {
	Input  int
	Output int
}

type InputIntOutputBool struct {
	Input  int
	Output bool
}

type InputIntOutputString struct {
	Input  int
	Output string
}

type Input1IntSliceInput2IntSliceOutputIntSlice struct {
	Input1 []int
	Input2 []int
	Output []int
}

type InputIntSliceOutputInt struct {
	Input  []int
	Output int
}

type Input1IntSliceInput2IntOutputInt struct {
	Input1 []int
	Input2 int
	Output int
}

type Input1IntSliceInput2IntOutputIntSlice struct {
	Input1 []int
	Input2 int
	Output []int
}

type InputIntSliceOutputIntSliceSlice struct {
	Input  []int
	Output [][]int
}

type Input1IntSliceInput2IntOutputIntSliceSlice struct {
	Input1 []int
	Input2 int
	Output [][]int
}

type InputStringOutputInt struct {
	Input  string
	Output int
}

type InputStringOutputBool struct {
	Input  string
	Output bool
}

type InputStringOutputString struct {
	Input  string
	Output string
}

type Input1StringInput2IntOutputString struct {
	Input1 string
	Input2 int
	Output string
}

type InputStringSliceOutputString struct {
	Input  []string
	Output string
}

type InputStringOutputStringSlice struct {
	Input  string
	Output []string
}
