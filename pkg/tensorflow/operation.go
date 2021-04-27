package tensorflow

import (
	"fmt"

	// Modules
	tf "github.com/djthorpe/go-tensorflow/tensorflow/github.com/tensorflow/tensorflow/tensorflow/go"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Op tf.Operation

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (o *Op) Name() string {
	return (*tf.Operation)(o).Name()
}

func (o *Op) Type() string {
	return (*tf.Operation)(o).Type()
}

func (o *Op) NumInputs() int {
	return (*tf.Operation)(o).NumInputs()
}

func (o *Op) NumOutputs() int {
	return (*tf.Operation)(o).NumOutputs()
}

func (o *Op) Output(i int) Output {
	return Output((*tf.Operation)(o).Output(i))
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (o *Op) String() string {
	str := "<op"
	str += fmt.Sprintf(" name=%q", o.Name())
	if t := o.Type(); t != o.Name() {
		str += fmt.Sprintf(" type=%q", t)
	}
	str += fmt.Sprint(" num_inputs=", o.NumInputs())
	str += fmt.Sprint(" num_outputs=", o.NumOutputs())
	return str + ">"
}
