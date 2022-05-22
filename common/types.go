package common

type InputIntOutputInt struct {
	Input  int
	Output int
}

type InputIntOutputBool struct {
	Input  int
	Output bool
}

type Input1IntSliceInput2IntSliceOutputIntSlice struct {
	Input1 []int
	Input2 []int
	Output []int
}

type InputStringOutputInt struct {
	Input  string
	Output int
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
