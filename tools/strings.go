package tools

type Strings struct {
}

func (t Strings) EmptyNewLine() string {
	return "\n‎"
}

func (t Strings) NoBreakSpace() string {
	return "‎"
}

func (t Strings) RightToLeftMark() string {
	return "‎"
}

func (t Strings) LeftToRightMark() string {
	return "‎"
}
