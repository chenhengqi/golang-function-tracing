package parse

import (
	"fmt"
	"io"
	"strconv"
)

// Parser is used to parse function arguments
type Parser struct {
	tokenizer *Tokenizer
	token     Token
	literal   string
	args      Args
}

// NewParser creates a new parser
func NewParser(src []byte) *Parser {
	return &Parser{
		tokenizer: NewTokenizer(src),
	}
}

// Parse parses function arguments
func (p *Parser) Parse() Args {
	for {
		p.next()
		switch p.token {
		case None:
			return p.args
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
			typeSpec := p.parseArg()
			p.args.Append(&Arg{typeSpec})
		default:
			typeSpec := p.parseArg()
			p.args.Append(&Arg{typeSpec})
		}
	}
}

func (p *Parser) parseArg() *TypeSpec {
	typeSpec := &TypeSpec{}
	switch p.token {
	case Bool, String, Int, Int8, Int16, Int32, Int64, Uint, Uint8, Uint16, Uint32, Uint64, Uintptr, Byte, Rune, Float32, Float64, Complex64, Complex128:
		typeSpec.isPrimitive = true
		typeSpec.token = p.token
	case Asterisk:
		typeSpec.isPointer = true
		p.next()
		typeSpec.inner = p.parseArg()
	case Chan:
		typeSpec.isChan = true
		p.next()
		typeSpec.inner = p.parseArg()
	case Error:
		typeSpec.isError = true
	case Interface:
		typeSpec.isInterface = true
	case Map:
		// map[string]int
		typeSpec.isMap = true
		p.expect(LeftSquareBracket)
		p.next()
		typeSpec.key = p.parseArg()
		p.expect(RightSquareBracket)
		p.next()
		typeSpec.value = p.parseArg()
	case LeftSquareBracket: // Array or Slice
		// [16]int, []int
		p.next()
		switch p.token {
		case RightSquareBracket:
			typeSpec.isSlice = true
			p.next()
			typeSpec.inner = p.parseArg()
		case Number:
			d, err := strconv.ParseInt(p.literal, 10, 64)
			if err != nil {
				panic(err)
			}
			typeSpec.isArray = true
			typeSpec.dimension = int(d)
			p.expect(RightSquareBracket)
			p.next()
			typeSpec.inner = p.parseArg()
		default:
			panic("unexpected token after `[`")
		}
	case Struct:
		// struct {string, int, int}
		typeSpec.isStruct = true
		p.expect(LeftBrace)
	parseStructFields:
		for {
			p.next()
			switch p.token {
			case RightBrace:
				break parseStructFields
			case Comma:
				p.next()
				field := p.parseArg()
				typeSpec.fields = append(typeSpec.fields, field)
			default:
				field := p.parseArg()
				typeSpec.fields = append(typeSpec.fields, field)
			}
		}
	default:
		panic("not implemented")
	}
	return typeSpec
}

func (p *Parser) expect(token Token) {
	p.next()
	if p.token != token {
		panic(fmt.Sprintf("expect %+v, got %+v", token, p.token))
	}
}
