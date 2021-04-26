package tensorflow

import (
	tf "github.com/djthorpe/go-tensorflow/tensorflow/github.com/tensorflow/tensorflow/tensorflow/go"
	op "github.com/djthorpe/go-tensorflow/tensorflow/github.com/tensorflow/tensorflow/tensorflow/go/op"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Tensor struct {
	Root *op.Scope
	*Scope
	tf.Output
}

type Output tf.Output
type DataType tf.DataType

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewTensor(scope *Scope, output Output) *Tensor {
	this := new(Tensor)
	if scope == nil {
		return nil
	}
	this.Root = (*op.Scope)(scope)
	this.Scope = NewScope(scope)
	this.Output = (tf.Output)(output)
	return this
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *Tensor) DataType() DataType {
	return DataType(this.Output.DataType())
}
