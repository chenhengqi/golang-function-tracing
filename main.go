package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"os"

	"github.com/chenhengqi/golang-function-tracing/internal/codegen"
	"github.com/iovisor/gobpf/bcc"
)

var processID int
var functionName string
var functionArguments string

func init() {
	flag.IntVar(&processID, "pid", 0, "specify process ID")
	flag.StringVar(&functionName, "func", "", "specify the fully qualified function name, like `example.com/foobar/foo.(*Bar).Fuzzbuzz` or `example.com/foobar/foo.Baz`")
	flag.StringVar(&functionArguments, "args", "", "specify function arguments, like `(string, int)`")
}

func main() {
	flag.Parse()
	if processID == 0 || functionName == "" || functionArguments == "" {
		flag.Usage()
		os.Exit(0)
	}

	bpfProgram := codegen.Args(functionArguments)
	bpfModule := bcc.NewModule(bpfProgram, []string{})
	defer bpfModule.Close()

	uprobeFD, err := bpfModule.LoadUprobe(golangFuncProbe)
	if err != nil {
		panic(err)
	}

	binaryPath, err := resolveBinaryFullPath(processID)
	if err != nil {
		panic(err)
	}

	err = bpfModule.AttachUprobe(binaryPath, functionName, uprobeFD, processID)
	if err != nil {
		panic(err)
	}

	tableID := bpfModule.TableId(bpfTableName)
	table := bcc.NewTable(tableID, bpfModule)
	ch := make(chan []byte)
	perfMap, err := bcc.InitPerfMap(table, ch, nil)
	if err != nil {
		panic(err)
	}

	perfMap.Start()
	defer perfMap.Stop()

	for entry := range ch {
		length := binary.LittleEndian.Uint64(entry)
		fmt.Printf("Length = %v\n", length)
		value := string(entry[8 : 8+int(length)])
		fmt.Printf("Value = %v\n", value)
	}
}
