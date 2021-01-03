package parse

import (
	"testing"
)

func TestTokenizer(t *testing.T) {
	tokenizer := NewTokenizer([]byte(`(int, string, *struct{int}, []byte)`))
	for token, raw, err := tokenizer.Next(); err == nil; {
		t.Logf("%v: %s\n", token, raw)
		token, raw, err = tokenizer.Next()
	}
}
