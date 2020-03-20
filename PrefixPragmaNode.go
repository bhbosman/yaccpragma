package yaccpragma

type PrefixPragmaNode struct {
	prefix string
}

func (self PrefixPragmaNode) IsIPragmaPrefixNode() bool {
	return true
}

func (self PrefixPragmaNode) PragmaType() string {
	return "Prefix"
}

func (self PrefixPragmaNode) Prefix() string {
	return self.prefix
}

func NewPrefixPragmaNode(prefix string) IPragmaPrefixNode {
	return &PrefixPragmaNode{prefix: prefix}
}
