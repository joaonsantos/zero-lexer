package lexer_test

import (
	"log"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/joaonsantos/zero-lexer/lexer"
	"github.com/joaonsantos/zero-lexer/token"
)

func TestNextToken(t *testing.T) {
	tt := []struct {
		name string
		in   string
		exp  []token.Token
	}{
		{
			name: "ex1",
			in:   "let var = 5 + 5;",
			exp: []token.Token{
				{
					Type: token.LET,
				},
				{
					Type:  token.IDENT,
					Value: "var",
				},
				{
					Type: token.ASSIGN,
				},
				{
					Type:  token.INTEGER,
					Value: "5",
				},
				{
					Type: token.PLUS,
				},
				{
					Type:  token.INTEGER,
					Value: "5",
				},
				{
					Type: token.SCOLON,
				},
				{
					Type: token.EOF,
				},
			},
		},
		{
			name: "ex2",
			in: `
			let add = fn(x,y) {
				return x + y;
			};
				`,
			exp: []token.Token{
				{
					Type: token.LET,
				},
				{
					Type:  token.IDENT,
					Value: "add",
				},
				{
					Type: token.ASSIGN,
				},
				{
					Type: token.FN,
				},
				{
					Type: token.LPAREN,
				},
				{
					Type:  token.IDENT,
					Value: "x",
				},
				{
					Type: token.COLON,
				},
				{
					Type:  token.IDENT,
					Value: "y",
				},
				{
					Type: token.RPAREN,
				},
				{
					Type: token.LBRACE,
				},
				{
					Type: token.RET,
				},
				{
					Type:  token.IDENT,
					Value: "x",
				},
				{
					Type: token.PLUS,
				},
				{
					Type:  token.IDENT,
					Value: "y",
				},
				{
					Type: token.SCOLON,
				},
				{
					Type: token.RBRACE,
				},
				{
					Type: token.SCOLON,
				},
				{
					Type: token.EOF,
				},
			},
		},
		{
			name: "ex3",
			in: `
			let x = 1;
			let y = 2;
			let result = add(x, y);
				`,
			exp: []token.Token{
				{
					Type: token.LET,
				},
				{
					Type:  token.IDENT,
					Value: "x",
				},
				{
					Type: token.ASSIGN,
				},
				{
					Type:  token.INTEGER,
					Value: "1",
				},
				{
					Type: token.SCOLON,
				},
				{
					Type: token.LET,
				},
				{
					Type:  token.IDENT,
					Value: "y",
				},
				{
					Type: token.ASSIGN,
				},
				{
					Type:  token.INTEGER,
					Value: "2",
				},
				{
					Type: token.SCOLON,
				},
				{
					Type: token.LET,
				},
				{
					Type:  token.IDENT,
					Value: "result",
				},
				{
					Type: token.ASSIGN,
				},
				{
					Type:  token.IDENT,
					Value: "add",
				},
				{
					Type: token.LPAREN,
				},
				{
					Type:  token.IDENT,
					Value: "x",
				},
				{
					Type: token.COLON,
				},
				{
					Type:  token.IDENT,
					Value: "y",
				},
				{
					Type: token.RPAREN,
				},
				{
					Type: token.SCOLON,
				},
				{
					Type: token.EOF,
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tokens := []token.Token{}
			l := lexer.New(tc.in)
			for i := 0; i < len(tc.exp); i++ {
				token := l.NextToken()
				tokens = append(tokens, token)

				if !cmp.Equal(token, tc.exp[i]) {
					if diff := cmp.Diff(token, tc.exp[i]); diff != "" {
						t.Errorf("test[%d] - NextToken() mismatch (+want -got): \n%s", i, diff)
					}
				}
			}
			log.Printf("%+v", tokens)
		})
	}
}
