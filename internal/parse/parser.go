package parse

// Parser is used to parse function arguments
type Parser struct {
	tokenizer Tokenizer
	token     Token
	literal   string
}
