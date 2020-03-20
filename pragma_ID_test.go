package yaccpragma

import (
	"github.com/bhbosman/yaccpragmaids"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPragmaId(t *testing.T) {
	t.Run("fail test", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RwPragma, data: nil},
				{token: RwId, data: nil},
				{token: Identifier, data: "A"},
				{token: StringLiteral, data: "ABCDEF"},
			})

		parser := YaccPragmaNewParser()
		v := parser.Parse(lexer)
		assert.Equal(t, int(1), v)
	})
	t.Run("IDL(1)", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RwPragma, data: nil},
				{token: RwId, data: nil},
				{token: Identifier, data: "A"},
				{token: StringLiteral, data: "IDL:A:1.1"},
			})

		parser := YaccPragmaNewParser()
		v := parser.Parse(lexer)
		if !assert.Equal(t, 0, v) {
			return
		}
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetPragma())
		assert.Equal(t, "Id", lexer.GetPragma().PragmaType())
		if !assert.Implements(t, (*IPragmaIdentifierNode)(nil), lexer.GetPragma()) {
			return
		}
		idPragmaNode, ok1 := lexer.GetPragma().(IPragmaIdentifierNode)
		assert.True(t, ok1)
		assert.Equal(t, "A", idPragmaNode.Identifier())
		nodeValue := idPragmaNode.Value()
		assert.NotNil(t, nodeValue)
		assert.Equal(t, "IDL", nodeValue.VersionType())
		idlVersion, ok := nodeValue.(yaccpragmaids.IIdlVersion)
		assert.True(t, ok)
		assert.Equal(t, uint64(1), idlVersion.Value().Major())
		assert.Equal(t, uint64(1), idlVersion.Value().Minor())
	})
	t.Run("IDL(2)", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RwPragma, data: nil},
				{token: RwId, data: nil},
				{token: Identifier, data: "A"},
				{token: StringLiteral, data: "IDL:A:1."},
			})

		parser := YaccPragmaNewParser()
		v := parser.Parse(lexer)
		if !assert.Equal(t, 0, v) {
			return
		}
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetPragma())
		assert.Equal(t, "Id", lexer.GetPragma().PragmaType())
		if !assert.Implements(t, (*IPragmaIdentifierNode)(nil), lexer.GetPragma()) {
			return
		}
		idPragmaNode, ok1 := lexer.GetPragma().(IPragmaIdentifierNode)
		assert.True(t, ok1)
		assert.Equal(t, "A", idPragmaNode.Identifier())
		nodeValue := idPragmaNode.Value()
		assert.NotNil(t, nodeValue)
		assert.Equal(t, "IDL", nodeValue.VersionType())
		idlVersion, ok := nodeValue.(yaccpragmaids.IIdlVersion)
		assert.True(t, ok)
		assert.Equal(t, uint64(1), idlVersion.Value().Major())
		assert.Equal(t, uint64(0), idlVersion.Value().Minor())
	})
	t.Run("IDL(3)", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RwPragma, data: nil},
				{token: RwId, data: nil},
				{token: Identifier, data: "A"},
				{token: StringLiteral, data: "IDL:A:1"},
			})

		parser := YaccPragmaNewParser()
		v := parser.Parse(lexer)
		if !assert.Equal(t, 0, v) {
			return
		}
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetPragma())
		assert.Equal(t, "Id", lexer.GetPragma().PragmaType())
		if !assert.Implements(t, (*IPragmaIdentifierNode)(nil), lexer.GetPragma()) {
			return
		}
		idPragmaNode, ok1 := lexer.GetPragma().(IPragmaIdentifierNode)
		assert.True(t, ok1)
		assert.Equal(t, "A", idPragmaNode.Identifier())
		nodeValue := idPragmaNode.Value()
		assert.NotNil(t, nodeValue)
		assert.Equal(t, "IDL", nodeValue.VersionType())
		idlVersion, ok := nodeValue.(yaccpragmaids.IIdlVersion)
		assert.True(t, ok)
		assert.Equal(t, uint64(1), idlVersion.Value().Major())
		assert.Equal(t, uint64(0), idlVersion.Value().Minor())
	})
	t.Run("RMI(1)", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RwPragma, data: nil},
				{token: RwId, data: nil},
				{token: Identifier, data: "A"},
				{token: StringLiteral, data: "RMI:Classname:AABBCCDDEEFF0011"},
			})

		parser := YaccPragmaNewParser()
		v := parser.Parse(lexer)
		if !assert.Equal(t, 0, v, lexer.errorMessage) {
			return
		}
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetPragma())
		assert.Equal(t, "Id", lexer.GetPragma().PragmaType())
		if !assert.Implements(t, (*IPragmaIdentifierNode)(nil), lexer.GetPragma()) {
			return
		}
		idPragmaNode, ok1 := lexer.GetPragma().(IPragmaIdentifierNode)
		assert.True(t, ok1)

		assert.Equal(t, "A", idPragmaNode.Identifier())
		nodeValue := idPragmaNode.Value()
		assert.NotNil(t, nodeValue)
		assert.Equal(t, "RMI", nodeValue.VersionType())

		rmiVersion, ok := nodeValue.(yaccpragmaids.IRmiVersion)
		assert.True(t, ok)
		assert.Equal(t, "Classname", rmiVersion.ClassName())
		assert.Equal(t, uint64(0xAABBCCDDEEFF0011), rmiVersion.Hash())
	})
	t.Run("RMI(2)", func(t *testing.T) {
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RwPragma, data: nil},
				{token: RwId, data: nil},
				{token: Identifier, data: "A"},
				{token: StringLiteral, data: "RMI:AA:FFFFFFFFFFFFFFFF:AAAAAAAAAAAAAAAA"},
			})

		parser := YaccPragmaNewParser()
		v := parser.Parse(lexer)
		if !assert.Equal(t, 0, v, lexer.errorMessage) {
			return
		}
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetPragma())
		assert.Equal(t, "Id", lexer.GetPragma().PragmaType())
		if !assert.Implements(t, (*IPragmaIdentifierNode)(nil), lexer.GetPragma()) {
			return
		}
		idPragmaNode, ok1 := lexer.GetPragma().(IPragmaIdentifierNode)
		assert.True(t, ok1)

		assert.Equal(t, "A", idPragmaNode.Identifier())
		nodeValue := idPragmaNode.Value()
		assert.NotNil(t, nodeValue)
		assert.Equal(t, "RMI", nodeValue.VersionType())

		rmiVersion, ok := nodeValue.(yaccpragmaids.IRmiVersion)
		assert.True(t, ok)
		assert.Equal(t, "AA", rmiVersion.ClassName())
		assert.Equal(t, uint64(0xFFFFFFFFFFFFFFFF), rmiVersion.Hash())
		assert.Equal(t, uint64(0xAAAAAAAAAAAAAAAA), rmiVersion.SerializationVersionUID())
	})
	t.Run("RMI(3)", func(t *testing.T) {
		// CORBA - Part 1: Interfaces, v3.3
		// 14.7.2 RMI Hashed Format
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RwPragma, data: nil},
				{token: RwId, data: nil},
				{token: Identifier, data: "A"},
				{token: StringLiteral, data: "RMI:foo/bar;:1234567812345678"},
			})

		parser := YaccPragmaNewParser()
		v := parser.Parse(lexer)
		if !assert.Equal(t, 0, v, lexer.errorMessage) {
			return
		}
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetPragma())
		assert.Equal(t, "Id", lexer.GetPragma().PragmaType())
		if !assert.Implements(t, (*IPragmaIdentifierNode)(nil), lexer.GetPragma()) {
			return
		}
		idPragmaNode, ok1 := lexer.GetPragma().(IPragmaIdentifierNode)
		assert.True(t, ok1)

		assert.Equal(t, "A", idPragmaNode.Identifier())
		nodeValue := idPragmaNode.Value()
		assert.NotNil(t, nodeValue)
		assert.Equal(t, "RMI", nodeValue.VersionType())

		rmiVersion, ok := nodeValue.(yaccpragmaids.IRmiVersion)
		assert.True(t, ok)
		assert.Equal(t, "foo/bar;", rmiVersion.ClassName())
		assert.Equal(t, uint64(0x1234567812345678), rmiVersion.Hash())
	})
	t.Run("RMI(4)", func(t *testing.T) {
		// CORBA - Part 1: Interfaces, v3.3
		// 14.7.2 RMI Hashed Format
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RwPragma, data: nil},
				{token: RwId, data: nil},
				{token: Identifier, data: "A"},
				{token: StringLiteral, data: "RMI:foo/bar;:1234567812345678:ABCD123456781234"},
			})

		parser := YaccPragmaNewParser()
		v := parser.Parse(lexer)
		if !assert.Equal(t, 0, v, lexer.errorMessage) {
			return
		}
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetPragma())
		assert.Equal(t, "Id", lexer.GetPragma().PragmaType())
		if !assert.Implements(t, (*IPragmaIdentifierNode)(nil), lexer.GetPragma()) {
			return
		}
		idPragmaNode, ok1 := lexer.GetPragma().(IPragmaIdentifierNode)
		assert.True(t, ok1)

		assert.Equal(t, "A", idPragmaNode.Identifier())
		nodeValue := idPragmaNode.Value()
		assert.NotNil(t, nodeValue)
		assert.Equal(t, "RMI", nodeValue.VersionType())

		rmiVersion, ok := nodeValue.(yaccpragmaids.IRmiVersion)
		assert.True(t, ok)
		assert.Equal(t, "foo/bar;", rmiVersion.ClassName())
		assert.Equal(t, uint64(0x1234567812345678), rmiVersion.Hash())
	})
	t.Run("DCE(1)", func(t *testing.T) {
		// CORBA - Part 1: Interfaces, v3.3
		// 14.7.3 DCE UUID Format
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RwPragma, data: nil},
				{token: RwId, data: nil},
				{token: Identifier, data: "A"},
				{token: StringLiteral, data: "DCE:700dc518-0110-11ce-ac8f-0800090b5d3e:1.9"},
			})

		parser := YaccPragmaNewParser()
		v := parser.Parse(lexer)
		if !assert.Equal(t, 0, v, lexer.errorMessage) {
			return
		}
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetPragma())
		assert.Equal(t, "Id", lexer.GetPragma().PragmaType())
		if !assert.Implements(t, (*IPragmaIdentifierNode)(nil), lexer.GetPragma()) {
			return
		}
		idPragmaNode, ok1 := lexer.GetPragma().(IPragmaIdentifierNode)
		assert.True(t, ok1)

		assert.Equal(t, "A", idPragmaNode.Identifier())
		nodeValue := idPragmaNode.Value()
		assert.NotNil(t, nodeValue)
		assert.Equal(t, "DCE", nodeValue.VersionType())

		dceVersion, ok := nodeValue.(yaccpragmaids.IDceVersion)
		assert.True(t, ok)
		assert.Equal(t, "700dc518-0110-11ce-ac8f-0800090b5d3e", dceVersion.Identifier())
		assert.Equal(t, uint64(1), dceVersion.Value().Major())
		assert.Equal(t, uint64(9), dceVersion.Value().Minor())
	})
	t.Run("DCE(2)", func(t *testing.T) {
		// CORBA - Part 1: Interfaces, v3.3
		// 14.7.3 DCE UUID Format
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RwPragma, data: nil},
				{token: RwId, data: nil},
				{token: Identifier, data: "A"},
				{token: StringLiteral, data: "DCE:700dc518-0110-11ce-ac8f-0800090b5d3e:1"},
			})

		parser := YaccPragmaNewParser()
		v := parser.Parse(lexer)
		if !assert.Equal(t, 0, v, lexer.errorMessage) {
			return
		}
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetPragma())
		assert.Equal(t, "Id", lexer.GetPragma().PragmaType())
		if !assert.Implements(t, (*IPragmaIdentifierNode)(nil), lexer.GetPragma()) {
			return
		}
		idPragmaNode, ok1 := lexer.GetPragma().(IPragmaIdentifierNode)
		assert.True(t, ok1)

		assert.Equal(t, "A", idPragmaNode.Identifier())
		nodeValue := idPragmaNode.Value()
		assert.NotNil(t, nodeValue)
		assert.Equal(t, "DCE", nodeValue.VersionType())

		dceVersion, ok := nodeValue.(yaccpragmaids.IDceVersion)
		assert.True(t, ok)
		assert.Equal(t, "700dc518-0110-11ce-ac8f-0800090b5d3e", dceVersion.Identifier())
		assert.Equal(t, uint64(1), dceVersion.Value().Major())
		assert.Equal(t, uint64(0), dceVersion.Value().Minor())
	})
	t.Run("DCE(3)", func(t *testing.T) {
		// CORBA - Part 1: Interfaces, v3.3
		// 14.7.3 DCE UUID Format
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RwPragma, data: nil},
				{token: RwId, data: nil},
				{token: Identifier, data: "A"},
				{token: StringLiteral, data: "DCE:700dc518-0110-11ce-ac8f-0800090b5d3e:1."},
			})

		parser := YaccPragmaNewParser()
		v := parser.Parse(lexer)
		if !assert.Equal(t, 0, v, lexer.errorMessage) {
			return
		}
		if !assert.False(t, lexer.error, lexer.errorMessage) {
			return
		}
		assert.NotNil(t, lexer.GetPragma())
		assert.Equal(t, "Id", lexer.GetPragma().PragmaType())
		if !assert.Implements(t, (*IPragmaIdentifierNode)(nil), lexer.GetPragma()) {
			return
		}
		idPragmaNode, ok1 := lexer.GetPragma().(IPragmaIdentifierNode)
		assert.True(t, ok1)

		assert.Equal(t, "A", idPragmaNode.Identifier())
		nodeValue := idPragmaNode.Value()
		assert.NotNil(t, nodeValue)
		assert.Equal(t, "DCE", nodeValue.VersionType())

		dceVersion, ok := nodeValue.(yaccpragmaids.IDceVersion)
		assert.True(t, ok)
		assert.Equal(t, "700dc518-0110-11ce-ac8f-0800090b5d3e", dceVersion.Identifier())
		assert.Equal(t, uint64(1), dceVersion.Value().Major())
		assert.Equal(t, uint64(0), dceVersion.Value().Minor())
	})
	t.Run("DCE(4)", func(t *testing.T) {
		// CORBA - Part 1: Interfaces, v3.3
		// 14.7.3 DCE UUID Format
		lexer := NewYaccIdlLexerImpl(
			0,
			[]lexem{
				{token: RwPragma, data: nil},
				{token: RwId, data: nil},
				{token: Identifier, data: "A"},
				{token: StringLiteral, data: "DCE:error:1."},
			})

		parser := YaccPragmaNewParser()
		v := parser.Parse(lexer)
		assert.Equal(t, 1, v, lexer.errorMessage)

	})
}
