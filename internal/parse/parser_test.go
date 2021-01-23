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

func TestParsePrimitive(t *testing.T) {
	src := "(bool, string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, byte, rune, float32, float64, complex64, complex128)"
	parser := NewParser([]byte(src))
	args := parser.Parse()
	if len(args.args) != 19 {
		t.Fatalf("expect 19 arguments, got %d", len(args.args))
	}
}

func TestParseComplex(t *testing.T) {
	src := "(*struct{int, string}, map[string]int, chan int, [8]int8, []int16, error, interface, struct{uint, uint8, uint16, uint32, uint64, uintptr, byte, rune, float32, float64, complex64, complex128})"
	parser := NewParser([]byte(src))
	args := parser.Parse()
	if len(args.args) != 8 {
		t.Fatalf("expect 8 arguments, got %d", len(args.args))
	}
}
