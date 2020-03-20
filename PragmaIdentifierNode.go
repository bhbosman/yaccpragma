package yaccpragma

import "github.com/bhbosman/yaccpragmaids"

type PragmaIdentifierNode struct {
	identifier string
	value      yaccpragmaids.ITypeVersion
}

func (self PragmaIdentifierNode) Identifier() string {
	return self.identifier
}

func (self PragmaIdentifierNode) PragmaType() string {
	return "Id"
}

func (self PragmaIdentifierNode) Value() yaccpragmaids.ITypeVersion {
	return self.value
}

func (self PragmaIdentifierNode) IsIPragmaIdentifierNode() bool {
	return true
}

func NewPragmaIdentifierNode(identifier string, value yaccpragmaids.ITypeVersion) IPragmaIdentifierNode {
	return &PragmaIdentifierNode{
		identifier: identifier,
		value:      value,
	}
}
