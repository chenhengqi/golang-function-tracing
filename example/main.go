package main

import (
	"example.com/foobar/foo"
)

func main() {
	bar := foo.Bar{
		I: 20210101,
		S: "hello world",
	}
	bar.Fuzzbuzz()

	foo.Baz("!")
}
