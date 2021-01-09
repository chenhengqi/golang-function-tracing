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
	args      []Arg
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
			p.literal = ""
			return
		}
		panic(err)
	}
	p.token = token
	p.literal = literal
}

func (p *Parser) parseArgs() {
	for {
		p.next()
		switch p.token {
		case RightParenthesis:
			return
		case Comma:
			p.next()
			p.parseArg()
		default:
			p.parseArg()
		}
	}
}

func (p *Parser) parseArg() {

}

func (p *Parser) expect(token Token) {
	p.next()
	if p.token != token {
		panic(fmt.Sprintf("expect %+v, got %+v", token, p.token))
	}
}
