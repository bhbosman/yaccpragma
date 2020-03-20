package yaccpragma

import (
	"github.com/bhbosman/lexpragmaids"
	"github.com/bhbosman/yaccpragmaids"
)

type IYaccPragmaIdsLexer interface {
	yaccpragmaids.YaccPragmaIdsLexer
	ErrorOccurred() bool
	ErrorMessage() string
	Version() yaccpragmaids.ITypeVersion
}

type IdsLexerWrapper struct {
	handler       *lexpragmaids.Handler
	version       yaccpragmaids.ITypeVersion
	errorOccurred bool
	errorMessage  string
}

func (self *IdsLexerWrapper) Version() yaccpragmaids.ITypeVersion {
	return self.version
}

func (self *IdsLexerWrapper) ErrorOccurred() bool {
	return self.errorOccurred
}

func (self *IdsLexerWrapper) ErrorMessage() string {
	return self.errorMessage
}

func (self *IdsLexerWrapper) SetTypeValue(version yaccpragmaids.ITypeVersion) {
	self.version = version
}

func (self IdsLexerWrapper) Lex(lval *yaccpragmaids.YaccPragmaIdsSymType) int {
	readLexem, err := self.handler.ReadLexem()
	if err != nil {

	}
	if readLexem.Eof {
		return 0
	}

	switch readLexem.TypeKind {
	case yaccpragmaids.RWLocal:
		return readLexem.TypeKind
	case yaccpragmaids.RWRmi:
		return readLexem.TypeKind
	case yaccpragmaids.RWDce:
		return readLexem.TypeKind
	case yaccpragmaids.Value:
		lval.Value = readLexem.Value
		return readLexem.TypeKind
	default:
		return readLexem.TypeKind
	}
}

func (self *IdsLexerWrapper) Error(s string) {
	self.errorOccurred = true
	self.errorMessage = s
}

func NewIdsLexerWrapper(handler *lexpragmaids.Handler) IYaccPragmaIdsLexer {
	return &IdsLexerWrapper{
		handler: handler,
	}
}
