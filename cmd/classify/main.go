package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	// Modules
	tf "github.com/djthorpe/go-tensorflow/pkg/tensorflow"
)

var (
	flagTag   = flag.String("tag", "", "Model tag")
	flagModel = flag.String("model", "", "Model path")
)

func main() {
	flag.Parse()

	// Require -model argument
	if *flagModel == "" {
		fmt.Fprintln(os.Stderr, "Missing -model argument")
		os.Exit(-1)
	}

	// If no -tag flag then do some introspection into what tags are available
	if *flagModel != "" && *flagTag == "" {
		if tags, err := tf.ModelTags(*flagModel); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		} else {
			fmt.Println("-tag", strings.Join(tags, ","))
			os.Exit(0)
		}
	}

	// Image argument is required
	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "Missing image argument")
		os.Exit(-1)
	}

	// Load model with tag
	model, err := tf.NewModelFromFile(*flagModel, []string{*flagTag})
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	} else {
		fmt.Println(model)
	}

	/*
		if tensor, err := LoadImage(flag.Arg(0)); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		} else {
			input := model.Operation("input_1").Output(0)
			output := model.Operation("predictions/Softmax").Output(0)
			if out, err := model.Run(map[tf.Output]*tf.Tensor{
				input: tensor,
			}, []tf.Output{output}, nil); err != nil {
				fmt.Fprintln(os.Stderr, err)
				os.Exit(-1)
			} else {
				fmt.Println("out=", out)
			}
			fmt.Println(tensor)
		}*/

}
