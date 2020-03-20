%{
package yaccpragma
import yaccpragmaids "github.com/bhbosman/yaccpragmaids"
import "fmt"

%}
%token RwPragma
%token RwPrefix
%token RwId
%token RwVersion
%token StringLiteral
%token WhiteSpace
%token Identifier
%token IntLiteral

%union {
	StringLiteral string
	IntLiteral int64
}
%token <StringLiteral> StringLiteral Identifier
%token <IntLiteral> IntLiteral
%start pragma_spec
%%

//(1)
pragma_spec:
	// Part01.pdf
	// 14.7.5.1 The ID Pragma
	RwPragma RwId Identifier StringLiteral
	{
		var version yaccpragmaids.ITypeVersion = nil
		if typeVersion, ok := YaccPragmalex.(ITypeVersionReader); ok {
			var err error = nil
			version, err = typeVersion.Read($4)
			if err != nil {
				YaccPragmalex.Error(fmt.Sprintf("ITypeVersionReader read failure. Error: %v", err.Error()))
				return 1
			}
		} else  {
			YaccPragmalex.Error("no ITypeVersionReader")
			return 1
		}

		if setPragma, ok := YaccPragmalex.(ISetPragma); ok {
			v := NewPragmaIdentifierNode($3, version)
			setPragma.SetPragma(v)
		}
	}
	// Part01.pdf
	// 14.7.5.2 The Prefix Pragma
	|RwPragma RwPrefix StringLiteral
	{
		if setPragma, ok := YaccPragmalex.(ISetPragma); ok {
			v := NewPrefixPragmaNode($3)
			setPragma.SetPragma(v)
		}
	}
	// Part01.pdf
	// 14.7.5.3 The Version Pragma
	|RwPragma RwVersion Identifier IntLiteral '.' IntLiteral
	{
		if setPragma, ok := YaccPragmalex.(ISetPragma); ok {
			v := NewPragmaVersionNode($3, $4, $6)
			setPragma.SetPragma(v)
		}
	}

%%



