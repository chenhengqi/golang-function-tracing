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
	src := "([8]int, [16]struct{string}, [4]byte, [8]byte)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	alignment := args.args[0].typeSpec.Alignment()
	var arr0 [18]int
	if uintptr(alignment) != unsafe.Alignof(arr0) {
		t.Fatalf("%T alignment expect %d, got %d", arr0, alignment, unsafe.Alignof(arr0))
	}

	alignment = args.args[1].typeSpec.Alignment()
	var arr1 [18]struct{ s string }
	if uintptr(alignment) != unsafe.Alignof(arr1) {
		t.Fatalf("%T alignment expect %d, got %d", arr1, alignment, unsafe.Alignof(arr1))
	}

	alignment = args.args[2].typeSpec.Alignment()
	var arr2 [4]byte
	if uintptr(alignment) != unsafe.Alignof(arr2) {
		t.Fatalf("%T alignment expect %d, got %d", arr2, alignment, unsafe.Alignof(arr2))
	}

	alignment = args.args[3].typeSpec.Alignment()
	var arr3 [8]byte
	if uintptr(alignment) != unsafe.Alignof(arr3) {
		t.Fatalf("%T alignment expect %d, got %d", arr3, alignment, unsafe.Alignof(arr3))
	}
}

func TestSliceAlignment(t *testing.T) {
	src := "([]int, []string, []struct{int, string}, []byte)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	alignment := args.args[0].typeSpec.Alignment()
	var s0 []int
	if uintptr(alignment) != unsafe.Alignof(s0) {
		t.Fatalf("%T alignment expect %d, got %d", s0, alignment, unsafe.Alignof(s0))
	}

	alignment = args.args[1].typeSpec.Alignment()
	var s1 []string
	if uintptr(alignment) != unsafe.Alignof(s1) {
		t.Fatalf("%T alignment expect %d, got %d", s1, alignment, unsafe.Alignof(s1))
	}

	alignment = args.args[2].typeSpec.Alignment()
	var s2 []struct {
		i int
		s string
	}
	if uintptr(alignment) != unsafe.Alignof(s2) {
		t.Fatalf("%T alignment expect %d, got %d", s2, alignment, unsafe.Alignof(s2))
	}

	alignment = args.args[3].typeSpec.Alignment()
	var s3 []byte
	if uintptr(alignment) != unsafe.Alignof(s3) {
		t.Fatalf("%T alignment expect %d, got %d", s3, alignment, unsafe.Alignof(s3))
	}
}

func TestMapAlignment(t *testing.T) {
	src := "(map[int]int, map[string]string, map[string]int, map[string]*struct{string, int})"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	alignment := args.args[0].typeSpec.Alignment()
	var m0 map[int]int
	if uintptr(alignment) != unsafe.Alignof(m0) {
		t.Fatalf("%T alignment expect %d, got %d", m0, alignment, unsafe.Alignof(m0))
	}

	alignment = args.args[1].typeSpec.Alignment()
	var m1 map[string]string
	if uintptr(alignment) != unsafe.Alignof(m1) {
		t.Fatalf("%T alignment expect %d, got %d", m1, alignment, unsafe.Alignof(m1))
	}

	alignment = args.args[2].typeSpec.Alignment()
	var m2 map[string]int
	if uintptr(alignment) != unsafe.Alignof(m2) {
		t.Fatalf("%T alignment expect %d, got %d", m2, alignment, unsafe.Alignof(m2))
	}

	alignment = args.args[3].typeSpec.Alignment()
	var m3 map[string]*struct {
		s string
		i int
	}
	if uintptr(alignment) != unsafe.Alignof(m3) {
		t.Fatalf("%T alignment expect %d, got %d", m3, alignment, unsafe.Alignof(m3))
	}
}

func TestChanAlignment(t *testing.T) {
	src := "(chan int, chan string, chan struct{string}, chan *struct{string, string})"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	alignment := args.args[0].typeSpec.Alignment()
	var ch0 chan int
	if uintptr(alignment) != unsafe.Alignof(ch0) {
		t.Fatalf("%T alignment expect %d, got %d", ch0, alignment, unsafe.Alignof(ch0))
	}

	alignment = args.args[1].typeSpec.Alignment()
	var ch1 chan string
	if uintptr(alignment) != unsafe.Alignof(ch1) {
		t.Fatalf("%T alignment expect %d, got %d", ch1, alignment, unsafe.Alignof(ch1))
	}

	alignment = args.args[2].typeSpec.Alignment()
	var ch2 chan struct{ s string }
	if uintptr(alignment) != unsafe.Alignof(ch2) {
		t.Fatalf("%T alignment expect %d, got %d", ch2, alignment, unsafe.Alignof(ch2))
	}

	alignment = args.args[3].typeSpec.Alignment()
	var ch3 chan *struct {
		s1 string
		s2 string
	}
	if uintptr(alignment) != unsafe.Alignof(ch3) {
		t.Fatalf("%T alignment expect %d, got %d", ch3, alignment, unsafe.Alignof(ch3))
	}
}

func TestPointerAlignment(t *testing.T) {
	src := "(*int, *string, *struct{})"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	alignment := args.args[0].typeSpec.Alignment()
	var p0 *int
	if uintptr(alignment) != unsafe.Alignof(p0) {
		t.Fatalf("%T alignment expect %d, got %d", p0, alignment, unsafe.Alignof(p0))
	}

	alignment = args.args[1].typeSpec.Alignment()
	var p1 *string
	if uintptr(alignment) != unsafe.Alignof(p1) {
		t.Fatalf("%T alignment expect %d, got %d", p1, alignment, unsafe.Alignof(p1))
	}

	alignment = args.args[2].typeSpec.Alignment()
	var p2 *struct{}
	if uintptr(alignment) != unsafe.Alignof(p2) {
		t.Fatalf("%T alignment expect %d, got %d", p2, alignment, unsafe.Alignof(p2))
	}
}

func TestStructAlignment(t *testing.T) {
	src := "(struct{}, struct{int, string}, struct{chan int})"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	alignment := args.args[0].typeSpec.Alignment()
	var s0 struct{}
	if uintptr(alignment) != unsafe.Alignof(s0) {
		t.Fatalf("%T alignment expect %d, got %d", s0, alignment, unsafe.Alignof(s0))
	}

	alignment = args.args[1].typeSpec.Alignment()
	var s1 struct {
		i int
		s string
	}
	if uintptr(alignment) != unsafe.Alignof(s1) {
		t.Fatalf("%T alignment expect %d, got %d", s1, alignment, unsafe.Alignof(s1))
	}

	alignment = args.args[2].typeSpec.Alignment()
	var s2 struct {
		ch chan int
	}
	if uintptr(alignment) != unsafe.Alignof(s2) {
		t.Fatalf("%T alignment expect %d, got %d", s2, alignment, unsafe.Alignof(s2))
	}
}

func TestPrimitiveSize(t *testing.T) {
	src := "(bool, string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, byte, rune, float32, float64, complex64, complex128)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	size := args.args[0].typeSpec.Size()
	var b bool
	if uintptr(size) != unsafe.Sizeof(b) {
		t.Fatalf("%T size expect %d, got %d", b, size, unsafe.Sizeof(b))
	}

	size = args.args[1].typeSpec.Size()
	var s string
	if uintptr(size) != unsafe.Sizeof(s) {
		t.Fatalf("%T size expect %d, got %d", s, size, unsafe.Sizeof(s))
	}

	size = args.args[2].typeSpec.Size()
	var i int
	if uintptr(size) != unsafe.Sizeof(i) {
		t.Fatalf("%T size expect %d, got %d", i, size, unsafe.Sizeof(i))
	}

	size = args.args[3].typeSpec.Size()
	var i8 int8
	if uintptr(size) != unsafe.Sizeof(i8) {
		t.Fatalf("%T size expect %d, got %d", i8, size, unsafe.Sizeof(i8))
	}

	size = args.args[4].typeSpec.Size()
	var i16 int16
	if uintptr(size) != unsafe.Sizeof(i16) {
		t.Fatalf("%T size expect %d, got %d", i16, size, unsafe.Sizeof(i16))
	}

	size = args.args[5].typeSpec.Size()
	var i32 int32
	if uintptr(size) != unsafe.Sizeof(i32) {
		t.Fatalf("%T size expect %d, got %d", i32, size, unsafe.Sizeof(i32))
	}

	size = args.args[6].typeSpec.Size()
	var i64 int64
	if uintptr(size) != unsafe.Sizeof(i64) {
		t.Fatalf("%T size expect %d, got %d", i64, size, unsafe.Sizeof(i64))
	}

	size = args.args[7].typeSpec.Size()
	var u uint
	if uintptr(size) != unsafe.Sizeof(u) {
		t.Fatalf("%T size expect %d, got %d", u, size, unsafe.Sizeof(u))
	}

	size = args.args[8].typeSpec.Size()
	var u8 uint8
	if uintptr(size) != unsafe.Sizeof(u8) {
		t.Fatalf("%T size expect %d, got %d", u8, size, unsafe.Sizeof(u8))
	}

	size = args.args[9].typeSpec.Size()
	var u16 uint16
	if uintptr(size) != unsafe.Sizeof(u16) {
		t.Fatalf("%T size expect %d, got %d", u16, size, unsafe.Sizeof(u16))
	}

	size = args.args[10].typeSpec.Size()
	var u32 uint32
	if uintptr(size) != unsafe.Sizeof(u32) {
		t.Fatalf("%T size expect %d, got %d", u32, size, unsafe.Sizeof(u32))
	}

	size = args.args[11].typeSpec.Size()
	var u64 uint64
	if uintptr(size) != unsafe.Sizeof(u64) {
		t.Fatalf("%T size expect %d, got %d", u64, size, unsafe.Sizeof(u64))
	}

	size = args.args[12].typeSpec.Size()
	var p uintptr
	if uintptr(size) != unsafe.Sizeof(p) {
		t.Fatalf("%T size expect %d, got %d", p, size, unsafe.Sizeof(p))
	}

	size = args.args[13].typeSpec.Size()
	var octet byte
	if uintptr(size) != unsafe.Sizeof(octet) {
		t.Fatalf("%T size expect %d, got %d", octet, size, unsafe.Sizeof(octet))
	}

	size = args.args[14].typeSpec.Size()
	var r rune
	if uintptr(size) != unsafe.Sizeof(r) {
		t.Fatalf("%T size expect %d, got %d", r, size, unsafe.Sizeof(r))
	}

	size = args.args[15].typeSpec.Size()
	var f32 float32
	if uintptr(size) != unsafe.Sizeof(f32) {
		t.Fatalf("%T size expect %d, got %d", f32, size, unsafe.Sizeof(f32))
	}

	size = args.args[16].typeSpec.Size()
	var f64 float64
	if uintptr(size) != unsafe.Sizeof(f64) {
		t.Fatalf("%T size expect %d, got %d", f64, size, unsafe.Sizeof(f64))
	}

	size = args.args[17].typeSpec.Size()
	var c64 complex64
	if uintptr(size) != unsafe.Sizeof(c64) {
		t.Fatalf("%T size expect %d, got %d", c64, size, unsafe.Sizeof(c64))
	}

	size = args.args[18].typeSpec.Size()
	var c128 complex128
	if uintptr(size) != unsafe.Sizeof(c128) {
		t.Fatalf("%T size expect %d, got %d", c128, size, unsafe.Sizeof(c128))
	}
}

func TestErrorSize(t *testing.T) {
	src := "(error)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	size := args.args[0].typeSpec.Size()
	var err error
	if uintptr(size) != unsafe.Sizeof(err) {
		t.Fatalf("%T size expect %d, got %d", err, size, unsafe.Sizeof(err))
	}
}

func TestInterfaceSize(t *testing.T) {
	src := "(interface)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	size := args.args[0].typeSpec.Size()
	var iface interface{}
	if uintptr(size) != unsafe.Sizeof(iface) {
		t.Fatalf("%T size expect %d, got %d", iface, size, unsafe.Sizeof(iface))
	}
}

func TestArraySize(t *testing.T) {
	src := "([5]byte, [8]int, [16]string)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	size := args.args[0].typeSpec.Size()
	var arr0 [5]byte
	if uintptr(size) != unsafe.Sizeof(arr0) {
		t.Fatalf("%T size expect %d, got %d", arr0, size, unsafe.Sizeof(arr0))
	}

	size = args.args[1].typeSpec.Size()
	var arr1 [8]int
	if uintptr(size) != unsafe.Sizeof(arr1) {
		t.Fatalf("%T size expect %d, got %d", arr1, size, unsafe.Sizeof(arr1))
	}

	size = args.args[2].typeSpec.Size()
	var arr2 [16]string
	if uintptr(size) != unsafe.Sizeof(arr2) {
		t.Fatalf("%T size expect %d, got %d", arr2, size, unsafe.Sizeof(arr2))
	}
}

func TestSliceSize(t *testing.T) {
	src := "([]byte, []int, []string)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	size := args.args[0].typeSpec.Size()
	var s0 []byte
	if uintptr(size) != unsafe.Sizeof(s0) {
		t.Fatalf("%T size expect %d, got %d", s0, size, unsafe.Sizeof(s0))
	}

	size = args.args[1].typeSpec.Size()
	var s1 []int
	if uintptr(size) != unsafe.Sizeof(s1) {
		t.Fatalf("%T size expect %d, got %d", s1, size, unsafe.Sizeof(s1))
	}

	size = args.args[2].typeSpec.Size()
	var s2 []string
	if uintptr(size) != unsafe.Sizeof(s2) {
		t.Fatalf("%T size expect %d, got %d", s2, size, unsafe.Sizeof(s2))
	}
}

func TestMapSize(t *testing.T) {
	src := "(map[int]int, map[string]string, map[string]int)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	size := args.args[0].typeSpec.Size()
	var m0 map[int]int
	if uintptr(size) != unsafe.Sizeof(m0) {
		t.Fatalf("%T size expect %d, got %d", m0, size, unsafe.Sizeof(m0))
	}

	size = args.args[1].typeSpec.Size()
	var m1 map[string]string
	if uintptr(size) != unsafe.Sizeof(m1) {
		t.Fatalf("%T size expect %d, got %d", m1, size, unsafe.Sizeof(m1))
	}

	size = args.args[2].typeSpec.Size()
	var m2 map[string]int
	if uintptr(size) != unsafe.Sizeof(m2) {
		t.Fatalf("%T size expect %d, got %d", m2, size, unsafe.Sizeof(m2))
	}
}

func TestChanSize(t *testing.T) {
	src := "(chan int, chan string)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	size := args.args[0].typeSpec.Size()
	var c0 chan int
	if uintptr(size) != unsafe.Sizeof(c0) {
		t.Fatalf("%T size expect %d, got %d", c0, size, unsafe.Sizeof(c0))
	}

	size = args.args[1].typeSpec.Size()
	var c1 chan string
	if uintptr(size) != unsafe.Sizeof(c1) {
		t.Fatalf("%T size expect %d, got %d", c1, size, unsafe.Sizeof(c1))
	}
}

func TestPointerSize(t *testing.T) {
	src := "(*int, *string, *struct{})"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	size := args.args[0].typeSpec.Size()
	var p0 *int
	if uintptr(size) != unsafe.Sizeof(p0) {
		t.Fatalf("%T size expect %d, got %d", p0, size, unsafe.Sizeof(p0))
	}

	size = args.args[1].typeSpec.Size()
	var p1 *string
	if uintptr(size) != unsafe.Sizeof(p1) {
		t.Fatalf("%T size expect %d, got %d", p1, size, unsafe.Sizeof(p1))
	}

	size = args.args[2].typeSpec.Size()
	var p2 *struct{}
	if uintptr(size) != unsafe.Sizeof(p2) {
		t.Fatalf("%T size expect %d, got %d", p2, size, unsafe.Sizeof(p2))
	}
}

func TestStructSize(t *testing.T) {
	src := "(struct{}, struct{int, string}, struct{byte, int})"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	size := args.args[0].typeSpec.Size()
	var s0 struct{}
	if uintptr(size) != unsafe.Sizeof(s0) {
		t.Fatalf("%T size expect %d, got %d", s0, size, unsafe.Sizeof(s0))
	}

	size = args.args[1].typeSpec.Size()
	var s1 struct {
		i int
		s string
	}
	if uintptr(size) != unsafe.Sizeof(s1) {
		t.Fatalf("%T size expect %d, got %d", s1, size, unsafe.Sizeof(s1))
	}

	size = args.args[2].typeSpec.Size()
	var s2 struct {
		b byte
		i int
	}
	if uintptr(size) != unsafe.Sizeof(s2) {
		t.Fatalf("%T size expect %d, got %d", s2, size, unsafe.Sizeof(s2))
	}
}
