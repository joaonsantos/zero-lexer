package lexer_test

import (
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
			fn add(x, y) {
				return x + y;
			};
				`,
			exp: []token.Token{
				{
					Type: token.FN,
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
					Type: token.COMMA,
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
			let result = add2(x, y);
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
					Value: "add2",
				},
				{
					Type: token.LPAREN,
				},
				{
					Type:  token.IDENT,
					Value: "x",
				},
				{
					Type: token.COMMA,
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
		{
			name: "ex4",
			in: `
			// some comment
				`,
			exp: []token.Token{
				{
					Type: token.EOF,
				},
			},
		},
		{
			name: "ex5",
			in: `
			let a = 1;
			let b = 2;
			if (a == b) { return 0; }
			if (a != b) { return 1; }
			if (a >= b) { return 0; }
			if (a <= b) { return 1; }
				`,
			exp: []token.Token{
				{
					Type: token.LET,
				},
				{
					Type:  token.IDENT,
					Value: "a",
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
					Value: "b",
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
					Type: token.IF,
				},
				{
					Type: token.LPAREN,
				},
				{
					Type:  token.IDENT,
					Value: "a",
				},
				{
					Type: token.EQ,
				},
				{
					Type:  token.IDENT,
					Value: "b",
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
					Type:  token.INTEGER,
					Value: "0",
				},
				{
					Type: token.SCOLON,
				},
				{
					Type: token.RBRACE,
				},
				{
					Type: token.IF,
				},
				{
					Type: token.LPAREN,
				},
				{
					Type:  token.IDENT,
					Value: "a",
				},
				{
					Type: token.NE,
				},
				{
					Type:  token.IDENT,
					Value: "b",
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
					Type:  token.INTEGER,
					Value: "1",
				},
				{
					Type: token.SCOLON,
				},
				{
					Type: token.RBRACE,
				},
				{
					Type: token.IF,
				},
				{
					Type: token.LPAREN,
				},
				{
					Type:  token.IDENT,
					Value: "a",
				},
				{
					Type: token.GTE,
				},
				{
					Type:  token.IDENT,
					Value: "b",
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
					Type:  token.INTEGER,
					Value: "0",
				},
				{
					Type: token.SCOLON,
				},
				{
					Type: token.RBRACE,
				},

				{
					Type: token.IF,
				},
				{
					Type: token.LPAREN,
				},
				{
					Type:  token.IDENT,
					Value: "a",
				},
				{
					Type: token.LTE,
				},
				{
					Type:  token.IDENT,
					Value: "b",
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
					Type:  token.INTEGER,
					Value: "1",
				},
				{
					Type: token.SCOLON,
				},
				{
					Type: token.RBRACE,
				},
				{
					Type: token.EOF,
				},
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			l := lexer.New(tc.in)

			for i := 0; i < len(tc.exp); i++ {
				token := l.NextToken()

				if !cmp.Equal(token, tc.exp[i]) {
					if diff := cmp.Diff(token, tc.exp[i]); diff != "" {
						t.Errorf("NextToken() mismatch (+want -got): \n%s", diff)
					}
				}
			}
		})
	}
}
