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

	EQ
	NE
	GT
	GTE
	LT
	LTE

	// Delimiters
	COMMA
	SCOLON

	LPAREN
	RPAREN
	LBRACE
	RBRACE

	// Keywords
	LET
	RET
	FN
	IF
)

type Token struct {
	Type  TokenType
	Value string
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
	case EQ:
		return "EQ"
	case NE:
		return "NE"
	case GT:
		return "GT"
	case GTE:
		return "GTE"
	case LT:
		return "LT"
	case LTE:
		return "LTE"
	case COMMA:
		return "COMMA"
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
	case IF:
		return "IF"
	}
	return "UNRECOGNIZED"
}

var keywordsMap = map[string]TokenType{
	"fn":     FN,
	"return": RET,
	"let":    LET,
	"if":     IF,
}

func LookupType(word string) (TokenType, bool) {
	if keywordType, ok := keywordsMap[word]; ok {
		return keywordType, true
	}
	return IDENT, false
}
