package parse

import "testing"

func TestParse(t *testing.T) {
	src := "(string, int)"
	parser := NewParser([]byte(src))
	args := parser.Parse()
	if len(args.args) != 2 {
		t.Fatalf("expect 2 arguments, got %d", len(args.args))
	}
}
