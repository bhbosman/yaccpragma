package yaccpragma

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPragmaVersion(t *testing.T) {
	t.Run("", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RwPragma, data: nil},
				{token: RwVersion, data: ""},
				{token: Identifier, data: "A"},
				{token: IntLiteral, data: int64(3)},
				{token: '.', data: nil},
				{token: IntLiteral, data: int64(9)},
			})

		parser := YaccPragmaNewParser()
		parser.Parse(lexer)
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetPragma())
		assert.Equal(t, "Version", lexer.GetPragma().PragmaType())
		if !assert.Implements(t, (*IPragmaVersionNode)(nil), lexer.GetPragma()) {
			return
		}
		idPragmaNode, ok1 := lexer.GetPragma().(IPragmaVersionNode)
		assert.True(t, ok1)
		assert.Equal(t, "A", idPragmaNode.Identifier())
		assert.Equal(t, int64(3), idPragmaNode.Major())
		assert.Equal(t, int64(9), idPragmaNode.Minor())
	})
}
