package parse

type goType int

// tokens for Go types
const (
	Bool goType = iota
	String
	Int
	Int8
	Int16
	Int32
	Int64
	Uint
	Uint8
	Uint16
	Uint32
	Uint64
	Uintptr
	Byte
	Rune
	Float32
	Float64
	Complex64
	Complex128
	Ptr
	Chan
	Array
	Slice
	Struct
	Interface
)
