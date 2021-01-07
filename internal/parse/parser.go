package parse

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

}

func (p *Parser) parseArgs() {

}

func (p *Parser) parseArg() {

}
