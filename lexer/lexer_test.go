package lexer

import (
	"testing"

	"github.com/golu360/fox/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
				let ten = 10;
				let add = fn(x, y) {
				x + y;
				};
				let result = add(five, ten);
				!-/*5;
				5 < 10 > 5;
				if (5 < 10) {
				return true;
				} else {
				return false;
				}
				10 == 10;
				10 != 9;`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
	}
	lexer := NewLexer(input)
	for i, tt := range tests {
		tok := lexer.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
