package yaccpragma

import (
	"fmt"
	"github.com/bhbosman/lexpragmaids"
	"github.com/bhbosman/yaccpragmaids"
)

type lexem struct {
	token int
	data  interface{}
}

type YaccPragmaLexerImplForTesting struct {
	pos          int
	data         []lexem
	error        bool
	errorMessage string
	pragmaNode   IPragmaNode
}

func (y *YaccPragmaLexerImplForTesting) Read(s string) (yaccpragmaids.ITypeVersion, error) {

	lex, err := lexpragmaids.NewPragmaLexIdsFromData("", s)
	if err != nil {
		return nil, err
	}
	lexWrapper := NewIdsLexerWrapper(lex)

	parser := yaccpragmaids.YaccPragmaIdsNewParser()
	v := parser.Parse(lexWrapper)
	if v != 0 || lexWrapper.ErrorOccurred() {
		return nil, fmt.Errorf("ErrorOccured in parsing pagmaids. Error %v: ", lexWrapper.ErrorMessage())
	}
	return lexWrapper.Version(), nil
}

func (y *YaccPragmaLexerImplForTesting) SetPragma(node IPragmaNode) {
	y.pragmaNode = node
}

func NewYaccIdlLexerImpl(pos int, data []lexem) *YaccPragmaLexerImplForTesting {
	return &YaccPragmaLexerImplForTesting{
		pos:  pos,
		data: data,
	}
}

func (y *YaccPragmaLexerImplForTesting) Lex(lval *YaccPragmaSymType) int {
	y.pos++
	if y.pos == len(y.data)+1 {
		return 0
	}
	switch y.data[y.pos-1].token {
	case StringLiteral, Identifier:
		if s, ok := y.data[y.pos-1].data.(string); ok {
			lval.StringLiteral = s
		}
		return y.data[y.pos-1].token
	case IntLiteral:
		if s, ok := y.data[y.pos-1].data.(int64); ok {
			lval.IntLiteral = s
		}
		return y.data[y.pos-1].token
	}

	return y.data[y.pos-1].token

}

func (y *YaccPragmaLexerImplForTesting) Error(s string) {
	y.error = true
	y.errorMessage = s
}

func (y *YaccPragmaLexerImplForTesting) GetPragma() IPragmaNode {
	return y.pragmaNode
}
