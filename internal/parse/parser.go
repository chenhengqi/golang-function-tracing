package parse

import (
	"fmt"
	"io"
)

// Parser is used to parse function arguments
type Parser struct {
	tokenizer *Tokenizer
	token     Token
	literal   string
}

// NewParser creates a new parser
func NewParser(src []byte) *Parser {
	return &Parser{
		tokenizer: NewTokenizer(src),
	}
}

// Parse parses function arguments
func (p *Parser) Parse() {
	for {
		p.next()
		switch p.token {
		case None:
			return
		case LeftParenthesis:
			p.parseArgs()
		default:
			panic(fmt.Sprintf("unexpected token: %+v, literal: %s", p.token, p.literal))
		}
	}
}

func (p *Parser) next() {
	token, literal, err := p.tokenizer.Next()
	if err != nil {
		if err == io.EOF {
			p.token = None
			return
		}
		panic(err)
	}
	p.token = token
	p.literal = literal
}

func (p *Parser) parseArgs() {

}

func (p *Parser) parseArg() {

}
