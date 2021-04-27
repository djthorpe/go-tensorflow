package tensorflow

import (
	"fmt"
	"sync/atomic"

	// Modules
	tf "github.com/djthorpe/go-tensorflow/tensorflow/github.com/tensorflow/tensorflow/tensorflow/go"
	op "github.com/djthorpe/go-tensorflow/tensorflow/github.com/tensorflow/tensorflow/tensorflow/go/op"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Scope op.Scope

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

var (
	ctr uint64
)

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

// NewRoot creates a new *op.Scope, empty
func NewRoot() *Scope {
	return (*Scope)(op.NewScope())
}

// NewScope returns a new scope
func NewScope(root *Scope) *Scope {
	var num = atomic.AddUint64(&ctr, 1)
	return (*Scope)((*op.Scope)(root).SubScope(fmt.Sprint("input_", num)))
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (this *Scope) Err() error {
	return (*op.Scope)(this).Err()
}

func (this *Scope) Finalize() (*Graph, error) {
	if graph, err := (*op.Scope)(this).Finalize(); err != nil {
		return nil, err
	} else {
<<<<<<< HEAD
		return &Graph{graph}, nil
=======
		return (*Graph)(graph), nil
>>>>>>> 15feaee8721bcc390fb6422c018f8fc9cd60c9aa
	}
}

func (this *Scope) Const(value interface{}) Output {
	return Output(op.Const((*op.Scope)(this).SubScope("const"), value))
}

///////////////////////////////////////////////////////////////////////////////
// MATH
// https://www.tensorflow.org/api_docs/python/tf/math

func (this *Scope) Abs(a Output) Output {
	return Output(op.Abs((*op.Scope)(this), tf.Output(a)))
}

func (this *Scope) Add(a, b Output) Output {
	return Output(op.Add((*op.Scope)(this), tf.Output(a), tf.Output(b)))
}

func (this *Scope) Mul(a, b Output) Output {
	return Output(op.Mul((*op.Scope)(this), tf.Output(a), tf.Output(b)))
}

func (this *Scope) MatMul(a, b Output) Output {
	return Output(op.MatMul((*op.Scope)(this), tf.Output(a), tf.Output(b)))
}

///////////////////////////////////////////////////////////////////////////////
// IO
// https://www.tensorflow.org/api_docs/python/tf/io

// -> string
func (this *Scope) ReadFile(path string) Output {
	return Output(op.ReadFile((*op.Scope)(this), tf.Output(this.Const(path))))
}

// string -> [][][]uint8
func (this *Scope) DecodeImage(data Output) Output {
	return Output(op.DecodeImage((*op.Scope)(this), tf.Output(data)))
}

// [][][]uint8 -> string
func (this *Scope) EncodePng(data Output) Output {
	return Output(op.EncodePng((*op.Scope)(this), tf.Output(data)))
}

// string -> operation
func (this *Scope) WriteFile(path string, data Output) *Op {
	return (*Op)(op.WriteFile((*op.Scope)(this), tf.Output(this.Const(path)), tf.Output(data)))
}

// [][][]float -> [][][]int8
func (this *Scope) AdjustSaturation(data Output, delta float32) Output {
	return Output(op.AdjustSaturation((*op.Scope)(this), tf.Output(data), tf.Output(this.Const(delta))))
}

///////////////////////////////////////////////////////////////////////////////
// IMAGE
// https://www.tensorflow.org/api_docs/python/tf/image
