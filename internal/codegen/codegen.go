package codegen

import (
	"regexp"
)

var funcNamePattern = regexp.MustCompile(`(?P<package_name>.*)\.(?P<func_name>[a-zA-Z][a-zA-Z0-9]*)`)
var funcNameWithReceiverPattern = regexp.MustCompile(`(?P<package_name>.*)\.(?P<receiver>\(?\*?[a-zA-Z][a-zA-Z0-9]*\)?)(?P<func_name>\.[a-zA-Z][a-zA-Z0-9]*)`)

// Args generates a code snippet for the provided arguments
func Args(funcName, funcArgs string) (string, error) {
	return "", nil
}
