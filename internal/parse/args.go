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
