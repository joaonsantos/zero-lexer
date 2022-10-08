package token

type TokenType int

const (
	// Special Tokens
	ILLEGAL TokenType = iota
	EOF

	// Identifiers + Literals
	IDENT
	INTEGER

	// Operators
	ASSIGN
	PLUS
	MINUS
	DIV
	MULT
	GT
	LT

	// Delimiters
	COLON
	SCOLON

	LPAREN
	RPAREN
	LBRACE
	RBRACE

	// Keywords
	LET
	RET
	FN
)

var keywords = map[string]TokenType{
	"fn":     FN,
	"return": RET,
	"let":    LET,
}

func LookupType(word string) (TokenType, bool) {
	if keywordType, ok := keywords[word]; ok {
		return keywordType, true
	}
	return IDENT, false
}

func (t TokenType) String() string {
	switch t {
	case ILLEGAL:
		return "ILLEGAL"
	case EOF:
		return "EOF"
	case IDENT:
		return "IDENT"
	case INTEGER:
		return "INTEGER"
	case ASSIGN:
		return "ASSIGN"
	case PLUS:
		return "PLUS"
	case MINUS:
		return "MINUS"
	case DIV:
		return "DIV"
	case MULT:
		return "MULT"
	case GT:
		return "GT"
	case LT:
		return "LT"
	case COLON:
		return "COLON"
	case SCOLON:
		return "SCOLON"
	case LPAREN:
		return "LPAREN"
	case RPAREN:
		return "RPAREN"
	case LBRACE:
		return "LBRACE"
	case RBRACE:
		return "RBRACE"
	case LET:
		return "LET"
	case RET:
		return "RET"
	case FN:
		return "FN"
	}
	return "UNRECOGNIZED"
}

type Token struct {
	Type  TokenType
	Value string
}
