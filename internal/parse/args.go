package parse

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
