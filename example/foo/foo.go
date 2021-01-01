package foo

// Bar is a simple struct used to demonstrate tracing
type Bar struct {
	I int
	S string
}

// Fuzzbuzz is a public API of struct Bar
func (b *Bar) Fuzzbuzz() {
	println(b.I)
	print(b.S)
}

// Baz is a public API of package foo
func Baz(s string) {
	println(s)
}
