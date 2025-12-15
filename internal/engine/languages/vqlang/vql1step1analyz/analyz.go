package vql1step1analyz

type TCode struct {
	Len   int
	Pos   int
	Lines []string
}

func Analyz(code string) TCode {

	return TCode{
		Len: 0,
		Pos: 0,
	}
}
