package codegen

import "testing"

func TestParseFunctionName(t *testing.T) {
	testcases := []struct {
		funcName string
		parts    []string
	}{
		{
			funcName: "main.main",
			parts:    []string{"main.main", "main", "main"},
		},
		{
			funcName: "foo.bar",
			parts:    []string{"foo.bar", "foo", "bar"},
		},
	}
	for _, testcase := range testcases {
		matches := funcNamePattern.FindStringSubmatch(testcase.funcName)
		n := len(matches)
		if n != len(testcase.parts) {
			t.Fatalf("function name NOT match, expect %+v, got %+v", testcase.parts, matches)
		}
	}
}

func TestParseFunctionNameWithReceiver(t *testing.T) {
	testcases := []struct {
		funcName string
		parts    []string
	}{
		{
			funcName: "main.Foo.Bar",
			parts:    []string{"main.Foo.Bar", "main", "Foo", "Bar"},
		},
		{
			funcName: "main.(*Foo).bar",
			parts:    []string{"main.(*Foo).bar", "main", "(*Foo)", "bar"},
		},
	}
	for _, testcase := range testcases {
		matches := funcNameWithReceiverPattern.FindStringSubmatch(testcase.funcName)
		n := len(matches)
		if n != len(testcase.parts) {
			t.Fatalf("function name NOT match, expect %+v, got %+v", testcase.parts, matches)
		}
	}
}