package yaccpragma

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPragmaPrefix(t *testing.T) {
	t.Run("", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RwPragma, data: nil},
				{token: RwPrefix, data: nil},
				{token: StringLiteral, data: "ABCDEF"},
			})

		parser := YaccPragmaNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetPragma())
		assert.Equal(t, "Prefix", lexer.GetPragma().PragmaType())
		if !assert.Implements(t, (*IPragmaPrefixNode)(nil), lexer.GetPragma()) {
			return
		}
		prefixPragmaNode, ok1 := lexer.GetPragma().(IPragmaPrefixNode)
		assert.True(t, ok1)
		assert.Equal(t, "ABCDEF", prefixPragmaNode.Prefix())
	})
}
