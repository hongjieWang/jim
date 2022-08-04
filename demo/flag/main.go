package main

import (
	"flag"
	"fmt"
	"os"
)

func init() {
	flag.CommandLine = flag.NewFlagSet("", flag.ExitOnError)
	flag.CommandLine.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", "question")
		flag.PrintDefaults()
	}
}

func main() {
	var name = flag.String("name", "杰子", "输入姓名")
	flag.Parse()
	fmt.Println("Hello: " + *name)
}
