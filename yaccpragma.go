package yaccpragma

import "github.com/bhbosman/yaccpragmaids"

type ITypeVersionReader interface {
	Read(s string) (yaccpragmaids.ITypeVersion, error)
}

type ISetPragma interface {
	SetPragma(node IPragmaNode)
}

type IGetPragma interface {
	GetPragma() IPragmaNode
}

type IPragmaNode interface {
	PragmaType() string
}

type IPragmaPrefixNode interface {
	IPragmaNode
	Prefix() string
	IsIPragmaPrefixNode() bool
}

type IPragmaIdentifierNode interface {
	IPragmaNode
	Identifier() string
	Value() yaccpragmaids.ITypeVersion
	IsIPragmaIdentifierNode() bool
}

type IPragmaVersionNode interface {
	IPragmaNode
	IsIPragmaVersionNode() bool
	Identifier() string
	Major() int64
	Minor() int64
}
