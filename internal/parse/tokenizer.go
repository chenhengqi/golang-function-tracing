package parse

import (
	"bytes"
	"io"
)

// Tokenizer is used to tokenize function arguments
type Tokenizer struct {
	src   []byte
	index int
}

// NewTokenizer creates a new Tokenizer
func NewTokenizer(src []byte) *Tokenizer {
	return &Tokenizer{
		src:   src,
		index: 0,
	}
}

func (t *Tokenizer) peek() (byte, error) {
	if t.index >= len(t.src) {
		return byte(0), io.EOF
	}
	return t.src[t.index], nil
}

// Next returns tokens one by one
func (t *Tokenizer) Next() (Token, string, error) {
scan:
	b, err := t.peek()
	if err != nil {
		return None, "", err
	}
	switch b {
	case ' ', '\t', '\n', '\r':
		t.index++
		goto scan
	case '(':
		t.index++
		return LeftParenthesis, "(", nil
	case ')':
		t.index++
		return RightParenthesis, ")", nil
	case '[':
		t.index++
		return LeftSquareBracket, "[", nil
	case ']':
		t.index++
		return RightSquareBracket, "]", nil
	case '{':
		t.index++
		return LeftBrace, "{", nil
	case '}':
		t.index++
		return RightBrace, "}", nil
	case ',':
		t.index++
		return Comma, ",", nil
	case '*':
		t.index++
		return Asterisk, "*", nil
	default:
		if '0' <= b && b <= '9' {
			return t.readNumber()
		}
		if 'a' <= b && b <= 'z' {
			return t.readSymbol()
		}
		if 'A' <= b && b <= 'Z' {
			return t.readSymbol()
		}
		panic("should not reach here")
	}
}

func (t *Tokenizer) readNumber() (Token, string, error) {
	num := bytes.NewBuffer(nil)
	for t.index < len(t.src) {
		if '0' <= t.src[t.index] && t.src[t.index] <= '9' {
			num.WriteByte(t.src[t.index])
			t.index++
		} else {
			break
		}
	}
	return Number, num.String(), nil
}

func (t *Tokenizer) readSymbol() (Token, string, error) {
	sym := bytes.NewBuffer(nil)
	for t.index < len(t.src) {
		if ('a' <= t.src[t.index] && t.src[t.index] <= 'z') ||
			('A' <= t.src[t.index] && t.src[t.index] <= 'Z') ||
			('0' <= t.src[t.index] && t.src[t.index] <= '9') {
			sym.WriteByte(t.src[t.index])
			t.index++
		} else {
			break
		}
	}
	return stringToToken[sym.String()], sym.String(), nil
}
