package parse

import "testing"

func TestParseReceiver(t *testing.T) {
	src := `*struct{int,string}`
	args := Receiver([]byte(src))
	if len(args) != 1 {
		t.Fatalf("wrong receiver size: %+v", len(args))
	}

	arg := args[0]
	if !arg.typeSpec.isPointer {
		t.Fatalf("wrong receiver type: %+v", arg)
	}
}
