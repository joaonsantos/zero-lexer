package lexer

import (
	"github.com/joaonsantos/zero-lexer/token"
)

type Lexer struct {
	source    string
	sourceLen int  // the length of the source input
	currPos   int  // current position in source
	readPos   int  // next position that will be read
	char      byte // char currPos points to
	row       int  // current row
	column    int  // current column

}

func New(input string) *Lexer {
	l := &Lexer{
		source:    input,
		sourceLen: len(input),
	}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.eatWhitespace()

	switch l.char {
	case '=':
		tok = newToken(token.ASSIGN)
	case '+':
		tok = newToken(token.PLUS)
	case '-':
		tok = newToken(token.MINUS)
	case '/':
		if l.peekChar() == '/' {
			// skip over comments
			l.readChar()
			l.eatComment()

			// must check EOF, in the edge case the program ends with a comment
			if isAtEOF(l.readPos, l.sourceLen) {
				tok = newToken(token.EOF)
			}
		} else {
			tok = newToken(token.DIV)
		}
	case '>':
		tok = newToken(token.GT)
	case '<':
		tok = newToken(token.MULT)
	case '*':
		tok = newToken(token.MULT)
	case ',':
		tok = newToken(token.COMMA)
	case ';':
		tok = newToken(token.SCOLON)
	case '(':
		tok = newToken(token.LPAREN)
	case ')':
		tok = newToken(token.RPAREN)
	case '{':
		tok = newToken(token.LBRACE)
	case '}':
		tok = newToken(token.RBRACE)
	case 0:
		tok = newToken(token.EOF)
	default:
		if isLetter(l.char) {
			identifier := l.readIdentifier()
			tokenType, ok := token.LookupType(identifier)
			switch ok {
			case true:
				tok = newToken(tokenType)
			default:
				tok = newTokenWithValue(tokenType, identifier)
			}
			return tok
		} else if isNumber(l.char) {
			number := l.readNumber()
			tokenType := token.INTEGER
			tok = newTokenWithValue(tokenType, number)
			return tok
		} else {
			tok = newToken(token.ILLEGAL)
		}
	}
	l.readChar()
	return tok
}

func newToken(typ token.TokenType) token.Token {
	return token.Token{
		Type: typ,
	}
}
func newTokenWithValue(typ token.TokenType, val string) token.Token {
	return token.Token{
		Type:  typ,
		Value: val,
	}
}

func (l *Lexer) readNumber() string {
	number := ""
	for isNumber(l.char) {
		number += string(l.char)
		l.readChar()
	}
	return number
}

func (l *Lexer) readIdentifier() string {
	identifier := ""
	for isAlphaNumeric(l.char) {
		identifier += string(l.char)
		l.readChar()
	}
	return identifier
}

func (l *Lexer) readChar() {
	if isAtEOF(l.readPos, l.sourceLen) {
		l.char = 0
		return
	}
	if isNewline(l.char) {
		l.column++
	}
	l.char = l.source[l.readPos]
	l.currPos = l.readPos
	l.readPos++
	l.row++
}

func (l *Lexer) peekChar() byte {
	if l.readPos > len(l.source) {
		return 0
	}
	return l.source[l.readPos]
}

func (l *Lexer) eatWhitespace() {
	for isWhitespace(l.char) {
		l.readChar()
	}
}

func (l *Lexer) eatComment() {
	for !isNewline(l.char) {
		l.readChar()
	}
	// read any whitespace on newline
	l.eatWhitespace()
}

func isNumber(char byte) bool {
	return char >= '0' && char <= '9'
}

func isLetter(char byte) bool {
	return char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z' || char == '_'
}

func isAlphaNumeric(char byte) bool {
	return isLetter(char) || isNumber(char)
}

func isAtEOF(readPos int, sourceLength int) bool {
	return readPos >= sourceLength
}

func isWhitespace(char byte) bool {
	return char == ' ' || char == '\t' || char == '\r' || isNewline(char)
}

func isNewline(char byte) bool {
	return char == '\n'
}
