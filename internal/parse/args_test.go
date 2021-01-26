package parse

import (
	"fmt"
	"testing"
	"unsafe"
)

func TestPrimitiveAlignment(t *testing.T) {
	src := "(bool)"
	parser := NewParser([]byte(src))
	args := parser.Parse()
	alignment := args.args[0].typeSpec.Alignment()

	var b bool
	if uintptr(alignment) != unsafe.Alignof(b) {
		panic(fmt.Sprintf("%T alignment expect %d, got %d", b, alignment, unsafe.Alignof(b)))
	}
}
