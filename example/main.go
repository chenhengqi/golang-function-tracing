package main

import (
	"time"

	"example.com/foobar/foo"
)

func main() {
	for {
		bar := foo.Bar{
			I: 20210101,
			S: "hello world",
		}
		bar.Fuzzbuzz()
		foo.Baz("!")

		time.Sleep(time.Second)
	}
}
