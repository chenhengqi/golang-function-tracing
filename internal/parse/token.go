package parse

// Token is used to distinct different tokens
type Token int

// tokens for Go types and symbols
const (
	None Token = iota
	Bool
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

	// (int, string, *struct{int}, []byte)
	LeftParenthesis
	RightParenthesis
	LeftSquareBracket
	RightSquareBracket
	LeftBrace
	RightBrace
	Comma
	Asterisk // for pointer
	Number   // for array
)

var stringToToken = map[string]Token{
	"bool":       Bool,
	"string":     String,
	"int":        Int,
	"int8":       Int8,
	"int16":      Int16,
	"int32":      Int32,
	"int64":      Int64,
	"uint":       Uint,
	"uint8":      Uint8,
	"uint16":     Uint16,
	"uint32":     Uint32,
	"uint64":     Uint64,
	"uintptr":    Uintptr,
	"byte":       Byte,
	"rune":       Rune,
	"float32":    Float32,
	"float64":    Float64,
	"complex64":  Complex64,
	"complex128": Complex128,
	"chan":       Chan,
	"struct":     Struct,
	"interface":  Interface,
}
