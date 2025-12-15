package vql1step2sintax

import "github.com/Kwynto/vql-prototype/internal/engine/languages/vqlang/vql1step1analyz"

type TCodeSintax struct {
	Len   int
	Pos   int
	Lines []string
}

func Sintax(code vql1step1analyz.TCode) TCodeSintax {

	return TCodeSintax{
		Len: 0,
		Pos: 0,
	}
}
