package vql1step3simant

import "github.com/Kwynto/vql-prototype/internal/engine/languages/vqlang/vql1step2sintax"

type TCodeSimant struct {
	Len   int
	Pos   int
	Lines []string
}

func Simant(code vql1step2sintax.TCodeSintax) TCodeSimant {

	return TCodeSimant{
		Len: 0,
		Pos: 0,
	}
}
