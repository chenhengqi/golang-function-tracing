package parse

// Receiver parses the input receiver
func Receiver(raw []byte) Args {
	src := "(" + string(raw) + ")"
	return NewParser([]byte(src)).Parse()
}
