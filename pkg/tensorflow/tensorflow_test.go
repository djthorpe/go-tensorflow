package tensorflow_test

import (
	"os"
	"path/filepath"
	"testing"

	// Modules
	tensorflow "github.com/djthorpe/go-tensorflow/pkg/tensorflow"
)

const (
	MODEL_PATH = "../../etc/nasnet"
	IMAGE_PATH = "../../etc/image/beach_wood.jpeg"
)

///////////////////////////////////////////////////////////////////////////////
// TESTS

func Test_Tensorflow_001(t *testing.T) {
	version := tensorflow.Version()
	if version == "" {
		t.Error("Version string should not be empty")
	} else {
		t.Log("Version=", version)
	}
}

func Test_Tensorflow_002(t *testing.T) {
	root := tensorflow.NewRoot()
	if root == nil {
		t.Fatal("Nil root scope")
	}
	scope1 := tensorflow.NewScope(root)
	if scope1 == nil {
		t.Fatal("Nil scope1")
	}
	scope2 := tensorflow.NewScope(root)
	if scope2 == nil {
		t.Fatal("Nil scope2")
	}
}

func Test_Tensorflow_003(t *testing.T) {
	const input = float64(2)

	root := tensorflow.NewRoot()
	c := root.Const(input)
	if graph, err := root.Finalize(); err != nil {
		t.Fatal(err)
	} else if result, err := graph.Run(nil, []tensorflow.Output{c}, nil); err != nil {
		t.Fatal(err)
	} else if result[0].Value().(float64) != input {
		t.Error("Unexpected output:", result)
	}
}

func Test_Tensorflow_004(t *testing.T) {
	const input = float64(2)

	root := tensorflow.NewRoot()
	op := root.Mul(root.Const(input), root.Const(input))
	if graph, err := root.Finalize(); err != nil {
		t.Fatal(err)
	} else if result, err := graph.Run(nil, []tensorflow.Output{op}, nil); err != nil {
		t.Fatal(err)
	} else if result[0].Value().(float64) != input*input {
		t.Error("Unexpected output:", result)
	}
}

func Test_Tensorflow_005(t *testing.T) {
	root := tensorflow.NewRoot()
	op := root.ReadFile(IMAGE_PATH)
	if graph, err := root.Finalize(); err != nil {
		t.Fatal(err)
	} else if result, err := graph.Run(nil, []tensorflow.Output{op}, nil); err != nil {
		t.Fatal(err)
	} else if result[0].DataType() != tensorflow.String {
		t.Fatal("Unexpected datatype")
	} else {
		t.Log(result[0])
	}
}

func Test_Tensorflow_006(t *testing.T) {
	root := tensorflow.NewRoot()
	op := root.DecodeImage(root.ReadFile(IMAGE_PATH))

	if graph, err := root.Finalize(); err != nil {
		t.Fatal(err)
	} else if result, err := graph.Run(nil, []tensorflow.Output{op}, nil); err != nil {
		t.Fatal(err)
	} else if result[0].DataType() != tensorflow.Uint8 {
		t.Fatal("Unexpected datatype", result[0].DataType())
	} else {
		t.Log(result[0])
	}
}

func Test_Tensorflow_007(t *testing.T) {
	root := tensorflow.NewRoot()
	tmpfile := filepath.Join(os.TempDir(), filepath.Base(IMAGE_PATH)+".png")
	op := root.WriteFile(tmpfile, root.EncodePng(root.DecodeImage(root.ReadFile(IMAGE_PATH))))
	if graph, err := root.Finalize(); err != nil {
		t.Fatal(err)
	} else if _, err := graph.Run(nil, []tensorflow.Output{}, []*tensorflow.Op{op}); err != nil {
		t.Fatal(err)
	} else {
		t.Log("Written to", tmpfile)
	}
}

func Test_Tensorflow_008(t *testing.T) {
	root := tensorflow.NewRoot()
	tmpfile := filepath.Join(os.TempDir(), filepath.Base(IMAGE_PATH)+".png")
	image := root.DecodeImage(root.ReadFile(IMAGE_PATH))
	op := root.WriteFile(tmpfile, root.EncodePng(root.AdjustSaturation(image, 0.5)))
	if graph, err := root.Finalize(); err != nil {
		t.Fatal(err)
	} else if _, err := graph.Run(nil, []tensorflow.Output{}, []*tensorflow.Op{op}); err != nil {
		t.Fatal(err)
	} else {
		t.Log("Written to", tmpfile)
	}
}
