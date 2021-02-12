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
		for i := 0; i < n; i++ {
			if matches[i] != testcase.parts[i] {
				t.Fatalf("%s and %s mismatched", matches[i], testcase.parts[i])
			}
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
			funcName: "main.(*Foo).Bar",
			parts:    []string{"main.(*Foo).Bar", "main", "(*Foo)", "Bar"},
		},
		{
			funcName: "example.com/foo/bar.(*Foo).Bar",
			parts:    []string{"example.com/foo/bar.(*Foo).Bar", "example.com/foo/bar", "(*Foo)", "Bar"},
		},
	}
	for _, testcase := range testcases {
		matches := funcNameWithReceiverPattern.FindStringSubmatch(testcase.funcName)
		n := len(matches)
		if n != len(testcase.parts) {
			t.Fatalf("function name NOT match, expect %+v, got %+v", testcase.parts, matches)
		}
		for i := 0; i < n; i++ {
			if matches[i] != testcase.parts[i] {
				t.Fatalf("%s and %s mismatched", matches[i], testcase.parts[i])
			}
		}
	}
}
