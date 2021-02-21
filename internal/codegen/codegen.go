package codegen

import (
	"fmt"
	"regexp"

	"github.com/chenhengqi/golang-function-tracing/internal/parse"
)

var funcNamePattern = regexp.MustCompile(`^(?P<package_name>.*)\.(?P<func_name>[a-zA-Z][a-zA-Z0-9]*)$`)
var funcNameWithReceiverPattern = regexp.MustCompile(`^(?P<package_name>.*)\.(?P<receiver>\(?\*?[a-zA-Z][a-zA-Z0-9]*\)?)\.(?P<func_name>[a-zA-Z][a-zA-Z0-9]*)$`)

var errWrongFunctionName = fmt.Errorf("wrong full qualified function name")
var errReceiverShouldBePointer = fmt.Errorf("receiver should be pointer")

// Args generates a code snippet for the provided arguments
func Args(fullFuncName, funcArgs string) (parse.Args, string, error) {
	// var stackOffset int

	names := funcNamePattern.FindStringSubmatch(fullFuncName)
	if len(names) != 3 {
		names = funcNameWithReceiverPattern.FindStringSubmatch(fullFuncName)
		if len(names) != 4 {
			return nil, "", errWrongFunctionName
		}
		receiver := names[2]
		if len(receiver) < 2 || receiver[0] != '(' || receiver[1] != '*' {
			return nil, "", errReceiverShouldBePointer
		}
		// stackOffset = 8
	}

	parser := parse.NewParser([]byte(funcArgs))
	args := parser.Parse()
	if len(args) == 0 {
		return args, bpfProgramForNoArgs, nil
	}
	return args, args.String(), nil
}

// ArgsWithReceiver generates a code snippet for the provided arguments
func ArgsWithReceiver(fullFuncName, funcArgs, receiver string) (parse.Args, string, error) {
	return nil, "", nil
}
