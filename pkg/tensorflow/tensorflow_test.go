package tensorflow_test

import (
	"testing"

	// Modules
	tensorflow "github.com/djthorpe/go-tensorflow/pkg/tensorflow"
)

const (
	MODEL_PATH = "../../etc/nasnet"
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
	model, err := tensorflow.NewModelFromFile(MODEL_PATH, []string{"photoprism"})
	if err != nil {
		t.Fatal("Error loading model:", err)
	}
	t.Log(model)
}
