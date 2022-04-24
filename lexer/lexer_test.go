package lexer

import (
	"calculator/token"
	"testing"
)

func TestGetNextToken(t *testing.T) {
	input := "1 + 2.3 -  456 *78.901 / (6 - 5)"
	tokenizer := NewLexer(input)
	tests := []struct {
		expected token.Token
	}{
		{token.Token{Type: token.NUMBER, Value: "1"}},
		{token.Token{Type: token.PLUS, Value: token.PLUS}},
		{token.Token{Type: token.NUMBER, Value: "2.3"}},
		{token.Token{Type: token.MINUS, Value: token.MINUS}},
		{token.Token{Type: token.NUMBER, Value: "456"}},
		{token.Token{Type: token.ASTERISK, Value: token.ASTERISK}},
		{token.Token{Type: token.NUMBER, Value: "78.901"}},
		{token.Token{Type: token.SLASH, Value: token.SLASH}},
		{token.Token{Type: token.LPAREN, Value: token.LPAREN}},
		{token.Token{Type: token.NUMBER, Value: "6"}},
		{token.Token{Type: token.MINUS, Value: token.MINUS}},
		{token.Token{Type: token.NUMBER, Value: "5"}},
		{token.Token{Type: token.RPAREN, Value: token.RPAREN}},
	}
	for _, tt := range tests {
		actual := tokenizer.GetNextToken()
		if actual.Type != tt.expected.Type || actual.Value != tt.expected.Value {
			t.Fatalf("not = %q, got = %q", tt.expected, actual)
		}
	}
}
