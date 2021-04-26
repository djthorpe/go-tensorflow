package tensorflow

import (
	tf "github.com/djthorpe/go-tensorflow/tensorflow/github.com/tensorflow/tensorflow/tensorflow/go"
	op "github.com/djthorpe/go-tensorflow/tensorflow/github.com/tensorflow/tensorflow/tensorflow/go/op"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type File struct {
	*tf.Tensor
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func NewFile(ctx *Scope, path string) (*File, error) {
	this := new(File)
	scope := NewScope(ctx)
	contents := op.ReadFile(scope.subScope("ReadFile"), op.Const(scope.subScope("filename"), path))
	this.Tensor = 
	// Return success
	return this, nil
}
