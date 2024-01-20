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
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT = "<"
	GT = ">"

	EQ     = "=="
	NOT_EQ = "!="

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
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	IFNT     = "IFNT" //I DONT WANT TO USE ELSE SO THIS IS THE BEST I CAN COME UP WITH
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"IFNT":   IFNT,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
