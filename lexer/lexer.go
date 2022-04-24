package lexer

import (
	"calculator/buffer"
	"calculator/token"
	"strings"
)

var tokens = map[string]token.Token{
	token.PLUS:     token.Token{Type: token.PLUS, Value: token.PLUS},
	token.MINUS:    token.Token{Type: token.MINUS, Value: token.MINUS},
	token.ASTERISK: token.Token{Type: token.ASTERISK, Value: token.ASTERISK},
	token.SLASH:    token.Token{Type: token.SLASH, Value: token.SLASH},
	token.LPAREN:   token.Token{Type: token.LPAREN, Value: token.LPAREN},
	token.RPAREN:   token.Token{Type: token.RPAREN, Value: token.RPAREN},
	token.EOF:      token.Token{Type: token.EOF, Value: token.EOF},
}

type Lexer struct {
	buf *buffer.Buffer
}

func NewLexer(input string) *Lexer {
	buf := buffer.NewBuffer(input)
	t := &Lexer{buf: buf}
	return t
}

// GetNextToken 次のトークンを取得する
func (l *Lexer) GetNextToken() token.Token {
	// 空白を飛ばす
	for isSpace(l.buf.Ch) {
		l.buf.Read()
	}
	// 数字のトークン
	if isNumber(l.buf.Ch) {
		sb := strings.Builder{}
		for isNumber(l.buf.Ch) {
			sb.WriteByte(l.buf.Ch)
			l.buf.Read()
		}
		// 整数型
		if l.buf.Ch != '.' {
			return token.Token{Type: token.NUMBER, Value: sb.String()}
		}
		sb.WriteByte(l.buf.Ch)
		l.buf.Read()
		for isNumber(l.buf.Ch) {
			sb.WriteByte(l.buf.Ch)
			l.buf.Read()
		}
		// 小数点型
		return token.Token{Type: token.NUMBER, Value: sb.String()}
	}
	// 記号のトークン
	if tok, ok := tokens[string(l.buf.Ch)]; ok {
		l.buf.Read()
		return tok
	}
	// 不明なトークン
	tok := token.Token{Type: token.UNKNOWN, Value: string(l.buf.Ch)}
	l.buf.Read()
	return tok
}

// isSpace 空白か判定する
func isSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n'
}

// isNumber 数字か判定する
func isNumber(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
