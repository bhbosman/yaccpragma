package yaccpragma

type lexem struct {
	token int
	data  interface{}
}

type YaccIdlLexerImpl struct {
	pos  int
	data []lexem
	//t            *testing.T
	error        bool
	errorMessage string
}

func NewYaccIdlLexerImpl(pos int, data []lexem) *YaccIdlLexerImpl {
	return &YaccIdlLexerImpl{
		pos:  pos,
		data: data,
	}
}

func (y *YaccIdlLexerImpl) Lex(lval *YaccPragmaSymType) int {
	y.pos++
	if y.pos == len(y.data)+1 {
		return 0
	}
	switch y.data[y.pos-1].token {
	case StringLiteral:
		if s, ok := y.data[y.pos-1].data.(string); ok {
			lval.StringLiteral = s
		}
		return y.data[y.pos-1].token

	}

	return y.data[y.pos-1].token

}

func (y *YaccIdlLexerImpl) Error(s string) {
	y.error = true
	y.errorMessage = s
}
