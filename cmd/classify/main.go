package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/djthorpe/go-tensorflow/pkg/classify"
)

func main() {
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "Missing model argument")
		os.Exit(-1)
	}

	if model, err := classify.New(flag.Arg(0)); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	} else {
		fmt.Println(model)
	}
}
