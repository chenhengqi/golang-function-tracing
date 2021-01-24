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
	Map
	Chan
	Array
	Error
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
	"map":        Map,
	"chan":       Chan,
	"error":      Error,
	"struct":     Struct,
	"interface":  Interface,
}

var tokenToString = map[Token]string{
	Bool:       "bool",
	String:     "string",
	Int:        "int",
	Int8:       "int8",
	Int16:      "int16",
	Int32:      "int32",
	Int64:      "int64",
	Uint:       "uint",
	Uint8:      "uint8",
	Uint16:     "uint16",
	Uint32:     "uint32",
	Uint64:     "uint64",
	Uintptr:    "uintptr",
	Byte:       "byte",
	Rune:       "rune",
	Float32:    "float32",
	Float64:    "float64",
	Complex64:  "complex64",
	Complex128: "complex128",
	Map:        "map",
	Chan:       "chan",
	Error:      "error",
	Struct:     "struct",
	Interface:  "interface",
}

// String implements the Stringer interface
func (t Token) String() string {
	return tokenToString[t]
}
