package token

//TODO: Possible improvement, use byte or int
type TokenType byte

type Token struct {
	Type     TokenType
	Literal  string
	FileName string
	Line     uint32
}

const (
	ILLEGAL = 0
	EOF     = 1

	//Identifiers + literals
	IDENT = 10
	INT   = 11

	//Operators
	ASSIGN   = 20 // =
	PLUS     = 21 // +
	QUESTION = 22 // ?

	// Delimiters
	COMMA     = 30 //,
	SEMICOLON = 31 //;

	LPAREN = 40 //(
	RPAREN = 41
	LBRACE = 42 //{
	RBRACE = 43 //}

	//Keywords
	DATA     = 50
	ENDPOINT = 51
	GROUP    = 52
	TYPE     = 53

	GET     = 64
	POST    = 65
	PUT     = 66
	DELTE   = 67
	HEAD    = 68
	OPTIONS = 69

	METHOD      = 70
	RESPONSE    = 71
	PATH        = 72
	DESCRIPTION = 73
)

//List of our keywords, may also be seen above
var keywords = map[string]TokenType{
	"data":     DATA,
	"endpoint": ENDPOINT,
	"group":    GROUP,
	"type":     TYPE,
}

//LookupIdent look up string to see if it's a keyword, if not it's an IDENT
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
