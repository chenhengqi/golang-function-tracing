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
		t.Fatalf("%T alignment expect %d, got %d", b, unsafe.Alignof(b), alignment)
	}

	alignment = args.args[1].typeSpec.Alignment()
	var s string
	if uintptr(alignment) != unsafe.Alignof(s) {
		t.Fatalf("%T alignment expect %d, got %d", s, unsafe.Alignof(s), alignment)
	}

	alignment = args.args[2].typeSpec.Alignment()
	var i int
	if uintptr(alignment) != unsafe.Alignof(i) {
		t.Fatalf("%T alignment expect %d, got %d", i, unsafe.Alignof(i), alignment)
	}

	alignment = args.args[3].typeSpec.Alignment()
	var i8 int8
	if uintptr(alignment) != unsafe.Alignof(i8) {
		t.Fatalf("%T alignment expect %d, got %d", i8, unsafe.Alignof(i8), alignment)
	}

	alignment = args.args[4].typeSpec.Alignment()
	var i16 int16
	if uintptr(alignment) != unsafe.Alignof(i16) {
		t.Fatalf("%T alignment expect %d, got %d", i16, unsafe.Alignof(i16), alignment)
	}

	alignment = args.args[5].typeSpec.Alignment()
	var i32 int32
	if uintptr(alignment) != unsafe.Alignof(i32) {
		t.Fatalf("%T alignment expect %d, got %d", i32, unsafe.Alignof(i32), alignment)
	}

	alignment = args.args[6].typeSpec.Alignment()
	var i64 int64
	if uintptr(alignment) != unsafe.Alignof(i64) {
		t.Fatalf("%T alignment expect %d, got %d", i64, unsafe.Alignof(i64), alignment)
	}

	alignment = args.args[7].typeSpec.Alignment()
	var u uint
	if uintptr(alignment) != unsafe.Alignof(u) {
		t.Fatalf("%T alignment expect %d, got %d", u, unsafe.Alignof(u), alignment)
	}

	alignment = args.args[8].typeSpec.Alignment()
	var u8 uint8
	if uintptr(alignment) != unsafe.Alignof(u8) {
		t.Fatalf("%T alignment expect %d, got %d", u8, unsafe.Alignof(u8), alignment)
	}

	alignment = args.args[9].typeSpec.Alignment()
	var u16 uint16
	if uintptr(alignment) != unsafe.Alignof(u16) {
		t.Fatalf("%T alignment expect %d, got %d", u16, unsafe.Alignof(u16), alignment)
	}

	alignment = args.args[10].typeSpec.Alignment()
	var u32 uint32
	if uintptr(alignment) != unsafe.Alignof(u32) {
		t.Fatalf("%T alignment expect %d, got %d", u32, unsafe.Alignof(u32), alignment)
	}

	alignment = args.args[11].typeSpec.Alignment()
	var u64 uint64
	if uintptr(alignment) != unsafe.Alignof(u64) {
		t.Fatalf("%T alignment expect %d, got %d", u64, unsafe.Alignof(u64), alignment)
	}

	alignment = args.args[12].typeSpec.Alignment()
	var p uintptr
	if uintptr(alignment) != unsafe.Alignof(p) {
		t.Fatalf("%T alignment expect %d, got %d", p, unsafe.Alignof(p), alignment)
	}

	alignment = args.args[13].typeSpec.Alignment()
	var octet byte
	if uintptr(alignment) != unsafe.Alignof(octet) {
		t.Fatalf("%T alignment expect %d, got %d", octet, unsafe.Alignof(octet), alignment)
	}

	alignment = args.args[14].typeSpec.Alignment()
	var r rune
	if uintptr(alignment) != unsafe.Alignof(r) {
		t.Fatalf("%T alignment expect %d, got %d", r, unsafe.Alignof(r), alignment)
	}

	alignment = args.args[15].typeSpec.Alignment()
	var f32 float32
	if uintptr(alignment) != unsafe.Alignof(f32) {
		t.Fatalf("%T alignment expect %d, got %d", f32, unsafe.Alignof(f32), alignment)
	}

	alignment = args.args[16].typeSpec.Alignment()
	var f64 float64
	if uintptr(alignment) != unsafe.Alignof(f64) {
		t.Fatalf("%T alignment expect %d, got %d", f64, unsafe.Alignof(f64), alignment)
	}

	alignment = args.args[17].typeSpec.Alignment()
	var c64 complex64
	if uintptr(alignment) != unsafe.Alignof(c64) {
		t.Fatalf("%T alignment expect %d, got %d", c64, unsafe.Alignof(c64), alignment)
	}

	alignment = args.args[18].typeSpec.Alignment()
	var c128 complex128
	if uintptr(alignment) != unsafe.Alignof(c128) {
		t.Fatalf("%T alignment expect %d, got %d", c128, unsafe.Alignof(c128), alignment)
	}
}

func TestErrorAlignment(t *testing.T) {
	src := "(error)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	alignment := args.args[0].typeSpec.Alignment()
	var err error
	if uintptr(alignment) != unsafe.Alignof(err) {
		t.Fatalf("%T alignment expect %d, got %d", err, unsafe.Alignof(err), alignment)
	}
}

func TestInterfaceAlignment(t *testing.T) {
	src := "(interface)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	alignment := args.args[0].typeSpec.Alignment()
	var iface interface{}
	if uintptr(alignment) != unsafe.Alignof(iface) {
		t.Fatalf("%T alignment expect %d, got %d", iface, unsafe.Alignof(iface), alignment)
	}
}

func TestArrayAlignment(t *testing.T) {
	src := "([8]int, [16]struct{string}, [4]byte, [8]byte)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	alignment := args.args[0].typeSpec.Alignment()
	var arr0 [18]int
	if uintptr(alignment) != unsafe.Alignof(arr0) {
		t.Fatalf("%T alignment expect %d, got %d", arr0, unsafe.Alignof(arr0), alignment)
	}

	alignment = args.args[1].typeSpec.Alignment()
	var arr1 [18]struct{ s string }
	if uintptr(alignment) != unsafe.Alignof(arr1) {
		t.Fatalf("%T alignment expect %d, got %d", arr1, unsafe.Alignof(arr1), alignment)
	}

	alignment = args.args[2].typeSpec.Alignment()
	var arr2 [4]byte
	if uintptr(alignment) != unsafe.Alignof(arr2) {
		t.Fatalf("%T alignment expect %d, got %d", arr2, unsafe.Alignof(arr2), alignment)
	}

	alignment = args.args[3].typeSpec.Alignment()
	var arr3 [8]byte
	if uintptr(alignment) != unsafe.Alignof(arr3) {
		t.Fatalf("%T alignment expect %d, got %d", arr3, unsafe.Alignof(arr3), alignment)
	}
}

func TestSliceAlignment(t *testing.T) {
	src := "([]int, []string, []struct{int, string}, []byte)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	alignment := args.args[0].typeSpec.Alignment()
	var s0 []int
	if uintptr(alignment) != unsafe.Alignof(s0) {
		t.Fatalf("%T alignment expect %d, got %d", s0, unsafe.Alignof(s0), alignment)
	}

	alignment = args.args[1].typeSpec.Alignment()
	var s1 []string
	if uintptr(alignment) != unsafe.Alignof(s1) {
		t.Fatalf("%T alignment expect %d, got %d", s1, unsafe.Alignof(s1), alignment)
	}

	alignment = args.args[2].typeSpec.Alignment()
	var s2 []struct {
		i int
		s string
	}
	if uintptr(alignment) != unsafe.Alignof(s2) {
		t.Fatalf("%T alignment expect %d, got %d", s2, unsafe.Alignof(s2), alignment)
	}

	alignment = args.args[3].typeSpec.Alignment()
	var s3 []byte
	if uintptr(alignment) != unsafe.Alignof(s3) {
		t.Fatalf("%T alignment expect %d, got %d", s3, unsafe.Alignof(s3), alignment)
	}
}

func TestMapAlignment(t *testing.T) {
	src := "(map[int]int, map[string]string, map[string]int, map[string]*struct{string, int})"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	alignment := args.args[0].typeSpec.Alignment()
	var m0 map[int]int
	if uintptr(alignment) != unsafe.Alignof(m0) {
		t.Fatalf("%T alignment expect %d, got %d", m0, unsafe.Alignof(m0), alignment)
	}

	alignment = args.args[1].typeSpec.Alignment()
	var m1 map[string]string
	if uintptr(alignment) != unsafe.Alignof(m1) {
		t.Fatalf("%T alignment expect %d, got %d", m1, unsafe.Alignof(m1), alignment)
	}

	alignment = args.args[2].typeSpec.Alignment()
	var m2 map[string]int
	if uintptr(alignment) != unsafe.Alignof(m2) {
		t.Fatalf("%T alignment expect %d, got %d", m2, unsafe.Alignof(m2), alignment)
	}

	alignment = args.args[3].typeSpec.Alignment()
	var m3 map[string]*struct {
		s string
		i int
	}
	if uintptr(alignment) != unsafe.Alignof(m3) {
		t.Fatalf("%T alignment expect %d, got %d", m3, unsafe.Alignof(m3), alignment)
	}
}

func TestChanAlignment(t *testing.T) {
	src := "(chan int, chan string, chan struct{string}, chan *struct{string, string})"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	alignment := args.args[0].typeSpec.Alignment()
	var ch0 chan int
	if uintptr(alignment) != unsafe.Alignof(ch0) {
		t.Fatalf("%T alignment expect %d, got %d", ch0, unsafe.Alignof(ch0), alignment)
	}

	alignment = args.args[1].typeSpec.Alignment()
	var ch1 chan string
	if uintptr(alignment) != unsafe.Alignof(ch1) {
		t.Fatalf("%T alignment expect %d, got %d", ch1, unsafe.Alignof(ch1), alignment)
	}

	alignment = args.args[2].typeSpec.Alignment()
	var ch2 chan struct{ s string }
	if uintptr(alignment) != unsafe.Alignof(ch2) {
		t.Fatalf("%T alignment expect %d, got %d", ch2, unsafe.Alignof(ch2), alignment)
	}

	alignment = args.args[3].typeSpec.Alignment()
	var ch3 chan *struct {
		s1 string
		s2 string
	}
	if uintptr(alignment) != unsafe.Alignof(ch3) {
		t.Fatalf("%T alignment expect %d, got %d", ch3, unsafe.Alignof(ch3), alignment)
	}
}

func TestPointerAlignment(t *testing.T) {
	src := "(*int, *string, *struct{})"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	alignment := args.args[0].typeSpec.Alignment()
	var p0 *int
	if uintptr(alignment) != unsafe.Alignof(p0) {
		t.Fatalf("%T alignment expect %d, got %d", p0, unsafe.Alignof(p0), alignment)
	}

	alignment = args.args[1].typeSpec.Alignment()
	var p1 *string
	if uintptr(alignment) != unsafe.Alignof(p1) {
		t.Fatalf("%T alignment expect %d, got %d", p1, unsafe.Alignof(p1), alignment)
	}

	alignment = args.args[2].typeSpec.Alignment()
	var p2 *struct{}
	if uintptr(alignment) != unsafe.Alignof(p2) {
		t.Fatalf("%T alignment expect %d, got %d", p2, unsafe.Alignof(p2), alignment)
	}
}

func TestStructAlignment(t *testing.T) {
	src := "(struct{}, struct{int, string}, struct{chan int})"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	alignment := args.args[0].typeSpec.Alignment()
	var s0 struct{}
	if uintptr(alignment) != unsafe.Alignof(s0) {
		t.Fatalf("%T alignment expect %d, got %d", s0, unsafe.Alignof(s0), alignment)
	}

	alignment = args.args[1].typeSpec.Alignment()
	var s1 struct {
		i int
		s string
	}
	if uintptr(alignment) != unsafe.Alignof(s1) {
		t.Fatalf("%T alignment expect %d, got %d", s1, unsafe.Alignof(s1), alignment)
	}

	alignment = args.args[2].typeSpec.Alignment()
	var s2 struct {
		ch chan int
	}
	if uintptr(alignment) != unsafe.Alignof(s2) {
		t.Fatalf("%T alignment expect %d, got %d", s2, unsafe.Alignof(s2), alignment)
	}
}

func TestPrimitiveSize(t *testing.T) {
	src := "(bool, string, int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr, byte, rune, float32, float64, complex64, complex128)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	size := args.args[0].typeSpec.Size()
	var b bool
	if uintptr(size) != unsafe.Sizeof(b) {
		t.Fatalf("%T size expect %d, got %d", b, unsafe.Sizeof(b), size)
	}

	size = args.args[1].typeSpec.Size()
	var s string
	if uintptr(size) != unsafe.Sizeof(s) {
		t.Fatalf("%T size expect %d, got %d", s, unsafe.Sizeof(s), size)
	}

	size = args.args[2].typeSpec.Size()
	var i int
	if uintptr(size) != unsafe.Sizeof(i) {
		t.Fatalf("%T size expect %d, got %d", i, unsafe.Sizeof(i), size)
	}

	size = args.args[3].typeSpec.Size()
	var i8 int8
	if uintptr(size) != unsafe.Sizeof(i8) {
		t.Fatalf("%T size expect %d, got %d", i8, unsafe.Sizeof(i8), size)
	}

	size = args.args[4].typeSpec.Size()
	var i16 int16
	if uintptr(size) != unsafe.Sizeof(i16) {
		t.Fatalf("%T size expect %d, got %d", i16, unsafe.Sizeof(i16), size)
	}

	size = args.args[5].typeSpec.Size()
	var i32 int32
	if uintptr(size) != unsafe.Sizeof(i32) {
		t.Fatalf("%T size expect %d, got %d", i32, unsafe.Sizeof(i32), size)
	}

	size = args.args[6].typeSpec.Size()
	var i64 int64
	if uintptr(size) != unsafe.Sizeof(i64) {
		t.Fatalf("%T size expect %d, got %d", i64, unsafe.Sizeof(i64), size)
	}

	size = args.args[7].typeSpec.Size()
	var u uint
	if uintptr(size) != unsafe.Sizeof(u) {
		t.Fatalf("%T size expect %d, got %d", u, unsafe.Sizeof(u), size)
	}

	size = args.args[8].typeSpec.Size()
	var u8 uint8
	if uintptr(size) != unsafe.Sizeof(u8) {
		t.Fatalf("%T size expect %d, got %d", u8, unsafe.Sizeof(u8), size)
	}

	size = args.args[9].typeSpec.Size()
	var u16 uint16
	if uintptr(size) != unsafe.Sizeof(u16) {
		t.Fatalf("%T size expect %d, got %d", u16, unsafe.Sizeof(u16), size)
	}

	size = args.args[10].typeSpec.Size()
	var u32 uint32
	if uintptr(size) != unsafe.Sizeof(u32) {
		t.Fatalf("%T size expect %d, got %d", u32, unsafe.Sizeof(u32), size)
	}

	size = args.args[11].typeSpec.Size()
	var u64 uint64
	if uintptr(size) != unsafe.Sizeof(u64) {
		t.Fatalf("%T size expect %d, got %d", u64, unsafe.Sizeof(u64), size)
	}

	size = args.args[12].typeSpec.Size()
	var p uintptr
	if uintptr(size) != unsafe.Sizeof(p) {
		t.Fatalf("%T size expect %d, got %d", p, unsafe.Sizeof(p), size)
	}

	size = args.args[13].typeSpec.Size()
	var octet byte
	if uintptr(size) != unsafe.Sizeof(octet) {
		t.Fatalf("%T size expect %d, got %d", octet, unsafe.Sizeof(octet), size)
	}

	size = args.args[14].typeSpec.Size()
	var r rune
	if uintptr(size) != unsafe.Sizeof(r) {
		t.Fatalf("%T size expect %d, got %d", r, unsafe.Sizeof(r), size)
	}

	size = args.args[15].typeSpec.Size()
	var f32 float32
	if uintptr(size) != unsafe.Sizeof(f32) {
		t.Fatalf("%T size expect %d, got %d", f32, unsafe.Sizeof(f32), size)
	}

	size = args.args[16].typeSpec.Size()
	var f64 float64
	if uintptr(size) != unsafe.Sizeof(f64) {
		t.Fatalf("%T size expect %d, got %d", f64, unsafe.Sizeof(f64), size)
	}

	size = args.args[17].typeSpec.Size()
	var c64 complex64
	if uintptr(size) != unsafe.Sizeof(c64) {
		t.Fatalf("%T size expect %d, got %d", c64, unsafe.Sizeof(c64), size)
	}

	size = args.args[18].typeSpec.Size()
	var c128 complex128
	if uintptr(size) != unsafe.Sizeof(c128) {
		t.Fatalf("%T size expect %d, got %d", c128, unsafe.Sizeof(c128), size)
	}
}

func TestErrorSize(t *testing.T) {
	src := "(error)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	size := args.args[0].typeSpec.Size()
	var err error
	if uintptr(size) != unsafe.Sizeof(err) {
		t.Fatalf("%T size expect %d, got %d", err, unsafe.Sizeof(err), size)
	}
}

func TestInterfaceSize(t *testing.T) {
	src := "(interface)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	size := args.args[0].typeSpec.Size()
	var iface interface{}
	if uintptr(size) != unsafe.Sizeof(iface) {
		t.Fatalf("%T size expect %d, got %d", iface, unsafe.Sizeof(iface), size)
	}
}

func TestArraySize(t *testing.T) {
	src := "([5]byte, [8]int, [16]string)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	size := args.args[0].typeSpec.Size()
	var arr0 [5]byte
	if uintptr(size) != unsafe.Sizeof(arr0) {
		t.Fatalf("%T size expect %d, got %d", arr0, unsafe.Sizeof(arr0), size)
	}

	size = args.args[1].typeSpec.Size()
	var arr1 [8]int
	if uintptr(size) != unsafe.Sizeof(arr1) {
		t.Fatalf("%T size expect %d, got %d", arr1, unsafe.Sizeof(arr1), size)
	}

	size = args.args[2].typeSpec.Size()
	var arr2 [16]string
	if uintptr(size) != unsafe.Sizeof(arr2) {
		t.Fatalf("%T size expect %d, got %d", arr2, unsafe.Sizeof(arr2), size)
	}
}

func TestSliceSize(t *testing.T) {
	src := "([]byte, []int, []string)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	size := args.args[0].typeSpec.Size()
	var s0 []byte
	if uintptr(size) != unsafe.Sizeof(s0) {
		t.Fatalf("%T size expect %d, got %d", s0, unsafe.Sizeof(s0), size)
	}

	size = args.args[1].typeSpec.Size()
	var s1 []int
	if uintptr(size) != unsafe.Sizeof(s1) {
		t.Fatalf("%T size expect %d, got %d", s1, unsafe.Sizeof(s1), size)
	}

	size = args.args[2].typeSpec.Size()
	var s2 []string
	if uintptr(size) != unsafe.Sizeof(s2) {
		t.Fatalf("%T size expect %d, got %d", s2, unsafe.Sizeof(s2), size)
	}
}

func TestMapSize(t *testing.T) {
	src := "(map[int]int, map[string]string, map[string]int)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	size := args.args[0].typeSpec.Size()
	var m0 map[int]int
	if uintptr(size) != unsafe.Sizeof(m0) {
		t.Fatalf("%T size expect %d, got %d", m0, unsafe.Sizeof(m0), size)
	}

	size = args.args[1].typeSpec.Size()
	var m1 map[string]string
	if uintptr(size) != unsafe.Sizeof(m1) {
		t.Fatalf("%T size expect %d, got %d", m1, unsafe.Sizeof(m1), size)
	}

	size = args.args[2].typeSpec.Size()
	var m2 map[string]int
	if uintptr(size) != unsafe.Sizeof(m2) {
		t.Fatalf("%T size expect %d, got %d", m2, unsafe.Sizeof(m2), size)
	}
}

func TestChanSize(t *testing.T) {
	src := "(chan int, chan string)"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	size := args.args[0].typeSpec.Size()
	var c0 chan int
	if uintptr(size) != unsafe.Sizeof(c0) {
		t.Fatalf("%T size expect %d, got %d", c0, unsafe.Sizeof(c0), size)
	}

	size = args.args[1].typeSpec.Size()
	var c1 chan string
	if uintptr(size) != unsafe.Sizeof(c1) {
		t.Fatalf("%T size expect %d, got %d", c1, unsafe.Sizeof(c1), size)
	}
}

func TestPointerSize(t *testing.T) {
	src := "(*int, *string, *struct{})"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	size := args.args[0].typeSpec.Size()
	var p0 *int
	if uintptr(size) != unsafe.Sizeof(p0) {
		t.Fatalf("%T size expect %d, got %d", p0, unsafe.Sizeof(p0), size)
	}

	size = args.args[1].typeSpec.Size()
	var p1 *string
	if uintptr(size) != unsafe.Sizeof(p1) {
		t.Fatalf("%T size expect %d, got %d", p1, unsafe.Sizeof(p1), size)
	}

	size = args.args[2].typeSpec.Size()
	var p2 *struct{}
	if uintptr(size) != unsafe.Sizeof(p2) {
		t.Fatalf("%T size expect %d, got %d", p2, unsafe.Sizeof(p2), size)
	}
}

func TestStructSize(t *testing.T) {
	src := "(struct{}, struct{int, string}, struct{byte, int}, struct{int, byte})"
	parser := NewParser([]byte(src))
	args := parser.Parse()

	size := args.args[0].typeSpec.Size()
	var s0 struct{}
	if uintptr(size) != unsafe.Sizeof(s0) {
		t.Fatalf("%T size expect %d, got %d", s0, unsafe.Sizeof(s0), size)
	}

	size = args.args[1].typeSpec.Size()
	var s1 struct {
		i int
		s string
	}
	if uintptr(size) != unsafe.Sizeof(s1) {
		t.Fatalf("%T size expect %d, got %d", s1, unsafe.Sizeof(s1), size)
	}

	size = args.args[2].typeSpec.Size()
	var s2 struct {
		b byte
		i int
	}
	if uintptr(size) != unsafe.Sizeof(s2) {
		t.Fatalf("%T size expect %d, got %d", s2, unsafe.Sizeof(s2), size)
	}

	size = args.args[3].typeSpec.Size()
	var s3 struct {
		i int
		b byte
	}
	if uintptr(size) != unsafe.Sizeof(s3) {
		t.Fatalf("%T size expect %d, got %d", s3, unsafe.Sizeof(s3), size)
	}
}
