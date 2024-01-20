package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	//special token
	EOF     = "EOF"     //end of file
	ILLEGAL = "ILLEGAL" //unknown tokens

	//identifier and literal
	IDENT = "IDENT" //Add, foo, bar
	INT   = "INT"   //234434

	//operators
	ASSIGN = "="
	PLUS   = "+"

	//delimeters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	//keywords

	FUNCTION = "FUNCTION"
	LET      = "LET"
)
