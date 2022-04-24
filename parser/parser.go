package parser

import (
	"calculator/lexer"
	"calculator/token"
	"fmt"
	"strconv"
)

type Parser struct {
	lexer    *lexer.Lexer
	curToken token.Token
}

func NewParser(input string) *Parser {
	l := lexer.NewLexer(input)
	p := &Parser{lexer: l}
	p.forwardToken()
	return p
}

func (p *Parser) Parse() (float64, error) {
	value, err := p.parseExpression()
	if err != nil {
		return value, err
	}
	// トークンがEOFで終わっていなければエラーを返す
	if p.curToken.Type != token.EOF {
		return 0, fmt.Errorf("無効なトークン%+v\n", p.curToken)
	}
	return value, nil
}

func (p *Parser) parseExpression() (float64, error) {
	leftValue, err := p.parseTerm()
	if err != nil {
		return leftValue, err
	}
	// 加算と減算を解析する
	for p.curToken.Type == token.PLUS || p.curToken.Type == token.MINUS {
		ope := p.curToken.Type
		p.forwardToken()
		rightValue, err := p.parseTerm()
		if err != nil {
			return rightValue, err
		}
		if ope == token.PLUS {
			leftValue += rightValue
		} else {
			leftValue -= rightValue
		}
	}
	return leftValue, nil
}

func (p *Parser) parseTerm() (float64, error) {
	leftValue, err := p.parseFactor()
	if err != nil {
		return leftValue, err
	}
	// 乗算と除算を解析する
	for p.curToken.Type == token.ASTERISK || p.curToken.Type == token.SLASH {
		ope := p.curToken.Type
		p.forwardToken()
		rightValue, err := p.parseFactor()
		if err != nil {
			return rightValue, err
		}
		if ope == token.ASTERISK {
			leftValue *= rightValue
		} else {
			leftValue /= rightValue
		}
	}
	return leftValue, nil
}

func (p *Parser) parseFactor() (float64, error) {
	// 括弧内の式を解析する
	if p.curToken.Type == token.LPAREN {
		p.forwardToken()
		value, err := p.parseExpression()
		if err != nil {
			return value, err
		}
		p.forwardToken()
		return value, nil
	}
	// 数値を解析する
	value, err := p.parseNumber()
	return value, err
}

func (p *Parser) parseNumber() (float64, error) {
	//　数値でなければエラーを返す
	if p.curToken.Type != token.NUMBER {
		return 0, fmt.Errorf("無効なトークン%+v\n", p.curToken)
	}
	value, _ := strconv.ParseFloat(p.curToken.Value, 64)
	p.forwardToken()
	return value, nil
}

func (p *Parser) forwardToken() {
	p.curToken = p.lexer.GetNextToken()
}
