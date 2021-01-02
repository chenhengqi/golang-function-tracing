package main

import (
	"fmt"
	"os"
)

func resolveBinaryFullPath(pid int) (string, error) {
	path := fmt.Sprintf("/proc/%d/exe", pid)
	return os.Readlink(path)
}
