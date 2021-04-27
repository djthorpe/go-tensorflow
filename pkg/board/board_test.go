package board_test

import (
	"os"
	"path/filepath"
	"testing"

	// Modules
	board "github.com/djthorpe/go-tensorflow/pkg/board"
	tensorflow "github.com/djthorpe/go-tensorflow/pkg/tensorflow"
)

const (
	MODEL_PATH = "../../etc/nasnet"
)

///////////////////////////////////////////////////////////////////////////////
// TESTS

func Test_Board_001(t *testing.T) {
	root := tensorflow.NewRoot()

	root.DecodeImage(root.ReadFile(MODEL_PATH))
	tmpfile := filepath.Join(os.TempDir(), "board.png")
	if graph, err := root.Finalize(); err != nil {
		t.Fatal(err)
	} else if board, err := board.NewBoardWithGraph(graph); err != nil {
		t.Fatal(err)
	} else if err := board.WritePNG(tmpfile); err != nil {
		t.Fatal(err)
	} else {
		t.Log("board written to:", tmpfile)
	}
}

func Test_Board_002(t *testing.T) {
	tmpfile := filepath.Join(os.TempDir(), filepath.Base(MODEL_PATH)+".png")
	if model, err := tensorflow.NewModelFromFile(MODEL_PATH, []string{"photoprism"}); err != nil {
		t.Fatal(err)
	} else if board, err := board.NewBoardWithModel(model); err != nil {
		t.Fatal(err)
	} else if err := board.WritePNG(tmpfile); err != nil {
		t.Fatal(err)
	} else {
		t.Log("board written to:", tmpfile)
	}
}
