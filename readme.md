# zero Lexer

This repository contains the code that implements a lexer
and a read-lex-print-loop for zero, a toy language.

## What is zero?

zero is simple and looks a lot like rust:
```rust
// examples
let x = (5 + 5) * 5 / 10;

fn sum(x, y) { return x + y;}

let sum_result = return sum(x,y);

let is_two_bigger_than_one = return 2 > 1;

let positive_y = if (y > 0) { return y; } else { return 0;}
```

## Try out the rlpl

The lexer supports a wide variety of tokens that you can check interactively.

The easiest way is to use the rlpl.

An example:
```
$ go run ./main.go

Welcome to the zero rlpl!
>> let x = (5 + 5) * 5 / 10;
{Type:LET Value:}
{Type:IDENT Value:x}
{Type:ASSIGN Value:}
{Type:LPAREN Value:}
{Type:INTEGER Value:5}
{Type:PLUS Value:}
{Type:INTEGER Value:5}
{Type:RPAREN Value:}
{Type:MULT Value:}
{Type:INTEGER Value:5}
{Type:DIV Value:}
{Type:INTEGER Value:10}
{Type:SCOLON Value:}

>> fn sum(x, y) { return x + y;}
{Type:FN Value:}
{Type:IDENT Value:sum}
{Type:LPAREN Value:}
{Type:IDENT Value:x}
{Type:COLON Value:}
{Type:IDENT Value:y}
{Type:RPAREN Value:}
{Type:LBRACE Value:}
{Type:RET Value:}
{Type:IDENT Value:x}
{Type:PLUS Value:}
{Type:IDENT Value:y}
{Type:SCOLON Value:}
{Type:RBRACE Value:}

>> let sum_result = return sum(x,y);
{Type:LET Value:}
{Type:IDENT Value:sum_result}
{Type:ASSIGN Value:}
{Type:RET Value:}
{Type:IDENT Value:sum}
{Type:LPAREN Value:}
{Type:IDENT Value:x}
{Type:COLON Value:}
{Type:IDENT Value:y}
{Type:RPAREN Value:}
{Type:SCOLON Value:}

>> let is_two_bigger_than_one = return 2 > 1;
{Type:LET Value:}
{Type:IDENT Value:is_two_bigger_than_one}
{Type:ASSIGN Value:}
{Type:RET Value:}
{Type:INTEGER Value:2}
{Type:GT Value:}
{Type:INTEGER Value:1}
{Type:SCOLON Value:}

>> let positive_y = if (y > 0) { return y; } else { return 0;}
{Type:LET Value:}
{Type:IDENT Value:positive_y}
{Type:ASSIGN Value:}
{Type:IDENT Value:if}
{Type:LPAREN Value:}
{Type:IDENT Value:y}
{Type:GT Value:}
{Type:INTEGER Value:0}
{Type:RPAREN Value:}
{Type:LBRACE Value:}
{Type:RET Value:}
{Type:IDENT Value:y}
{Type:SCOLON Value:}
{Type:RBRACE Value:}
{Type:IDENT Value:else}
{Type:LBRACE Value:}
{Type:RET Value:}
{Type:INTEGER Value:0}
{Type:SCOLON Value:}
{Type:RBRACE Value:}
```

>> let num = 42; // the answer to the ultimate question of life
{Type:LET Value:}
{Type:IDENT Value:num}
{Type:ASSIGN Value:}
{Type:INTEGER Value:42}
{Type:SCOLON Value:}
```

