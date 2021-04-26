package tensorflow

import (
	"fmt"

	tf "github.com/djthorpe/go-tensorflow/tensorflow/github.com/tensorflow/tensorflow/tensorflow/go"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Tensor tf.Tensor

/*
type Tensor struct {
	Root *op.Scope
	*Scope
	tf.Output
}
*/
///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func NewTensor(v interface{}) (*Tensor, error) {
	if tensor, err := tf.NewTensor(v); err != nil {
		return nil, err
	} else {
		return (*Tensor)(tensor), nil
	}
}

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (this *Tensor) DataType() DataType {
	return DataType((*tf.Tensor)(this).DataType())
}

func (this *Tensor) Value() interface{} {
	return (*tf.Tensor)(this).Value()
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (this *Tensor) String() string {
	str := "<tensor"
	str += fmt.Sprint(" type=", this.DataType())
	if this.DataType() == String {
		str += fmt.Sprintf(" value=%q", this.Value())
	} else {
		str += fmt.Sprint(" value=", this.Value())
	}
	return str + ">"
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS
