%{
package yaccpragma

%}
%token RWPragma
%token RWPrefix
%token StringLiteral

%union{
	StringLiteral string
}


%token <StringLiteral> StringLiteral

%start pragma_spec
%%

//(1)
pragma_spec:
	RWPragma RWPrefix StringLiteral
%%



