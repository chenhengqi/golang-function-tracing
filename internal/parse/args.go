package parse

import (
	"bytes"
	"strconv"
)

// Arg represents an function argument
type Arg struct {
	typeSpec *TypeSpec
}

// Args represents a function argument list
type Args struct {
	args []*Arg
}

// TypeSpec defines a type
type TypeSpec struct {
	isPrimitive bool
	isPointer   bool
	isArray     bool
	isSlice     bool
	isChan      bool
	isMap       bool
	isError     bool
	isStruct    bool
	isInterface bool
	inner       *TypeSpec   // for chan, array, slice
	key         *TypeSpec   // for map
	value       *TypeSpec   // for map
	dimension   int         // for array
	fields      []*TypeSpec // for struct
	token       Token
}

// Append appends an argument
func (a *Args) Append(arg *Arg) {
	a.args = append(a.args, arg)
}

// String implements the Stringer interface
func (a *Args) String() string {
	if len(a.args) == 0 {
		return "()"
	}
	buf := bytes.NewBufferString("(\n")
	for _, arg := range a.args {
		buf.WriteString("    ")
		buf.WriteString(arg.typeSpec.String())
		buf.WriteByte('\n')
	}
	buf.WriteString(")\n")
	return buf.String()
}

// String implements the Stringer interface
func (t *TypeSpec) String() string {
	if t.isPrimitive {
		return t.token.String()
	}
	if t.isError {
		return Error.String()
	}
	if t.isInterface {
		return Interface.String()
	}
	if t.isArray {
		return "[" + strconv.FormatInt(int64(t.dimension), 10) + "]" + t.inner.String()
	}
	if t.isSlice {
		return "[]" + t.inner.String()
	}
	if t.isMap {
		return "map[" + t.key.String() + "]" + t.value.String()
	}
	if t.isChan {
		return "chan " + t.inner.String()
	}
	if t.isPointer {
		return "*" + t.inner.String()
	}
	if t.isStruct {
		if len(t.fields) == 0 {
			return "struct{}"
		}
		buf := bytes.NewBufferString("struct {\n")
		for _, field := range t.fields {
			buf.WriteString("        ")
			buf.WriteString(field.String())
			buf.WriteByte('\n')
		}
		buf.WriteString("}\n")
		return buf.String()
	}
	return ""
}

// Alignment returns the alignment of the type in bytes
func (t *TypeSpec) Alignment() int {
	if t.isPrimitive {
		switch t.token {
		case Bool:
			return 1
		case String:
			return 8
		case Int:
			return 8
		case Int8:
			return 1
		case Int16:
			return 2
		case Int32:
			return 4
		case Int64:
			return 8
		case Uint:
			return 8
		case Uint8:
			return 1
		case Uint16:
			return 2
		case Uint32:
			return 4
		case Uint64:
			return 8
		case Uintptr:
			return 8
		case Byte:
			return 1
		case Rune:
			return 4
		case Float32:
			return 4
		case Float64:
			return 8
		case Complex64:
			return 4
		case Complex128:
			return 8
		default:
			panic("NOT primitive type")
		}
	}
	if t.isError {
		return 8
	}
	if t.isInterface {
		return 8
	}
	if t.isArray {
		return t.inner.Alignment()
	}
	if t.isSlice {
		return 8
	}
	if t.isMap {
		return 8
	}
	if t.isChan {
		return 8
	}
	if t.isPointer {
		return 8
	}
	if t.isStruct {
		if len(t.fields) == 0 {
			return 1
		}
		alignment := 1
		for _, field := range t.fields {
			if alignment < field.Alignment() {
				alignment = field.Alignment()
			}
		}
		return alignment
	}
	panic("not implemented")
}

// Size returns the size of the type in bytes
func (t *TypeSpec) Size() int {
	if t.isPrimitive {
		switch t.token {
		case Bool:
			return 1
		case String:
			return 16
		case Int:
			return 8
		case Int8:
			return 1
		case Int16:
			return 2
		case Int32:
			return 4
		case Int64:
			return 8
		case Uint:
			return 8
		case Uint8:
			return 1
		case Uint16:
			return 2
		case Uint32:
			return 4
		case Uint64:
			return 8
		case Uintptr:
			return 8
		case Byte:
			return 1
		case Rune:
			return 4
		case Float32:
			return 4
		case Float64:
			return 8
		case Complex64:
			return 8
		case Complex128:
			return 16
		default:
			panic("NOT primitive type")
		}
	}
	if t.isError {
		return 8
	}
	if t.isInterface {
		return 8
	}
	if t.isArray {
		return t.inner.Size() * t.dimension
	}
	if t.isSlice {
		return 24
	}
	if t.isMap {
		return 8
	}
	if t.isChan {
		return 8
	}
	if t.isPointer {
		return 8
	}
	if t.isStruct {
		if len(t.fields) == 0 {
			return 1
		}
		size := 0
		for _, field := range t.fields {
			size += field.Size()
		}
		return size
	}
	panic("not implemented")
}
