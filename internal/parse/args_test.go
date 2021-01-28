package parse

import (
	"testing"
	"unsafe"
)

func TestPrimitiveAlignment(t *testing.T) {
	src := "(bool, string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, byte, rune, float32, float64, complex64, complex128)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	alignment := args.args[0].typeSpec.Alignment()
	var b bool
	if uintptr(alignment) != unsafe.Alignof(b) {
		t.Fatalf("%T alignment expect %d, got %d", b, alignment, unsafe.Alignof(b))
	}

	alignment = args.args[1].typeSpec.Alignment()
	var s string
	if uintptr(alignment) != unsafe.Alignof(s) {
		t.Fatalf("%T alignment expect %d, got %d", s, alignment, unsafe.Alignof(s))
	}

	alignment = args.args[2].typeSpec.Alignment()
	var i int
	if uintptr(alignment) != unsafe.Alignof(i) {
		t.Fatalf("%T alignment expect %d, got %d", i, alignment, unsafe.Alignof(i))
	}

	alignment = args.args[3].typeSpec.Alignment()
	var i8 int8
	if uintptr(alignment) != unsafe.Alignof(i8) {
		t.Fatalf("%T alignment expect %d, got %d", i8, alignment, unsafe.Alignof(i8))
	}

	alignment = args.args[4].typeSpec.Alignment()
	var i16 int16
	if uintptr(alignment) != unsafe.Alignof(i16) {
		t.Fatalf("%T alignment expect %d, got %d", i16, alignment, unsafe.Alignof(i16))
	}

	alignment = args.args[5].typeSpec.Alignment()
	var i32 int32
	if uintptr(alignment) != unsafe.Alignof(i32) {
		t.Fatalf("%T alignment expect %d, got %d", i32, alignment, unsafe.Alignof(i32))
	}

	alignment = args.args[6].typeSpec.Alignment()
	var i64 int64
	if uintptr(alignment) != unsafe.Alignof(i64) {
		t.Fatalf("%T alignment expect %d, got %d", i64, alignment, unsafe.Alignof(i64))
	}

	alignment = args.args[7].typeSpec.Alignment()
	var u uint
	if uintptr(alignment) != unsafe.Alignof(u) {
		t.Fatalf("%T alignment expect %d, got %d", u, alignment, unsafe.Alignof(u))
	}

	alignment = args.args[8].typeSpec.Alignment()
	var u8 uint8
	if uintptr(alignment) != unsafe.Alignof(u8) {
		t.Fatalf("%T alignment expect %d, got %d", u8, alignment, unsafe.Alignof(u8))
	}

	alignment = args.args[9].typeSpec.Alignment()
	var u16 uint16
	if uintptr(alignment) != unsafe.Alignof(u16) {
		t.Fatalf("%T alignment expect %d, got %d", u16, alignment, unsafe.Alignof(u16))
	}

	alignment = args.args[10].typeSpec.Alignment()
	var u32 uint32
	if uintptr(alignment) != unsafe.Alignof(u32) {
		t.Fatalf("%T alignment expect %d, got %d", u32, alignment, unsafe.Alignof(u32))
	}

	alignment = args.args[11].typeSpec.Alignment()
	var u64 uint64
	if uintptr(alignment) != unsafe.Alignof(u64) {
		t.Fatalf("%T alignment expect %d, got %d", u64, alignment, unsafe.Alignof(u64))
	}

	alignment = args.args[12].typeSpec.Alignment()
	var p uintptr
	if uintptr(alignment) != unsafe.Alignof(p) {
		t.Fatalf("%T alignment expect %d, got %d", p, alignment, unsafe.Alignof(p))
	}

	alignment = args.args[13].typeSpec.Alignment()
	var octet byte
	if uintptr(alignment) != unsafe.Alignof(octet) {
		t.Fatalf("%T alignment expect %d, got %d", octet, alignment, unsafe.Alignof(octet))
	}

	alignment = args.args[14].typeSpec.Alignment()
	var r rune
	if uintptr(alignment) != unsafe.Alignof(r) {
		t.Fatalf("%T alignment expect %d, got %d", r, alignment, unsafe.Alignof(r))
	}

	alignment = args.args[15].typeSpec.Alignment()
	var f32 float32
	if uintptr(alignment) != unsafe.Alignof(f32) {
		t.Fatalf("%T alignment expect %d, got %d", f32, alignment, unsafe.Alignof(f32))
	}

	alignment = args.args[16].typeSpec.Alignment()
	var f64 float64
	if uintptr(alignment) != unsafe.Alignof(f64) {
		t.Fatalf("%T alignment expect %d, got %d", f64, alignment, unsafe.Alignof(f64))
	}

	alignment = args.args[17].typeSpec.Alignment()
	var c64 complex64
	if uintptr(alignment) != unsafe.Alignof(c64) {
		t.Fatalf("%T alignment expect %d, got %d", c64, alignment, unsafe.Alignof(c64))
	}

	alignment = args.args[18].typeSpec.Alignment()
	var c128 complex128
	if uintptr(alignment) != unsafe.Alignof(c128) {
		t.Fatalf("%T alignment expect %d, got %d", c128, alignment, unsafe.Alignof(c128))
	}
}

func TestErrorAlignment(t *testing.T) {
	src := "(error)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	alignment := args.args[0].typeSpec.Alignment()
	var err error
	if uintptr(alignment) != unsafe.Alignof(err) {
		t.Fatalf("%T alignment expect %d, got %d", err, alignment, unsafe.Alignof(err))
	}
}

func TestInterfaceAlignment(t *testing.T) {
	src := "(interface)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	alignment := args.args[0].typeSpec.Alignment()
	var iface interface{}
	if uintptr(alignment) != unsafe.Alignof(iface) {
		t.Fatalf("%T alignment expect %d, got %d", iface, alignment, unsafe.Alignof(iface))
	}
}

func TestArrayAlignment(t *testing.T) {
	src := "([8]int, [16]struct{string})"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	alignment := args.args[0].typeSpec.Alignment()
	var arr [18]int
	if uintptr(alignment) != unsafe.Alignof(arr) {
		t.Fatalf("%T alignment expect %d, got %d", arr, alignment, unsafe.Alignof(arr))
	}

	alignment = args.args[1].typeSpec.Alignment()
	var arr2 [18]struct{ s string }
	if uintptr(alignment) != unsafe.Alignof(arr2) {
		t.Fatalf("%T alignment expect %d, got %d", arr2, alignment, unsafe.Alignof(arr2))
	}
}

func TestSliceAlignment(t *testing.T) {
	src := "(interface)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	alignment := args.args[0].typeSpec.Alignment()
	var iface interface{}
	if uintptr(alignment) != unsafe.Alignof(iface) {
		t.Fatalf("%T alignment expect %d, got %d", iface, alignment, unsafe.Alignof(iface))
	}
}

func TestMapAlignment(t *testing.T) {
	src := "(interface)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	alignment := args.args[0].typeSpec.Alignment()
	var iface interface{}
	if uintptr(alignment) != unsafe.Alignof(iface) {
		t.Fatalf("%T alignment expect %d, got %d", iface, alignment, unsafe.Alignof(iface))
	}
}

func TestChanAlignment(t *testing.T) {
	src := "(interface)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	alignment := args.args[0].typeSpec.Alignment()
	var iface interface{}
	if uintptr(alignment) != unsafe.Alignof(iface) {
		t.Fatalf("%T alignment expect %d, got %d", iface, alignment, unsafe.Alignof(iface))
	}
}
