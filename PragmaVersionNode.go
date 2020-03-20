package yaccpragma

type PragmaVersionNode struct {
	identifier   string
	major, minor int64
}

func (self PragmaVersionNode) Major() int64 {
	return self.major
}

func (self PragmaVersionNode) Minor() int64 {
	return self.minor
}

func (self PragmaVersionNode) Identifier() string {
	return self.identifier
}

func (self PragmaVersionNode) IsIPragmaVersionNode() bool {
	return true
}

func (self PragmaVersionNode) PragmaType() string {
	return "Version"
}

func NewPragmaVersionNode(identifier string, major, minor int64) IPragmaVersionNode {
	return &PragmaVersionNode{
		identifier: identifier,
		major:      major,
		minor:      minor,
	}
}
