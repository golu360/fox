package parser

import (
	"github.com/golu360/fox/lexer"
	"github.com/golu360/fox/token"
)

type Parser struct {
	l *lexer.Lexer

	currToken token.Token
	peekToken token.Token
}

func NewParser(l *lexer.Lexer) *Parser {
	return &Parser{
		l: l,
	}
}

func (p *Parser) nextToken() {
	p.currToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
