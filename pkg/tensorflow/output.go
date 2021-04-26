package tensorflow

import (
	"fmt"

	tf "github.com/djthorpe/go-tensorflow/tensorflow/github.com/tensorflow/tensorflow/tensorflow/go"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Output tf.Output

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (o Output) Dimensions() int {
	return tf.Output(o).Shape().NumDimensions()
}

func (o Output) DataType() DataType {
	return DataType(tf.Output(o).DataType())
}

///////////////////////////////////////////////////////////////////////////////
// STRINGIFY

func (o Output) String() string {
	str := "<output"
	str += fmt.Sprintf(" name=%q", tf.Output(o).Op.Name())
	str += fmt.Sprintf(" datatype=%q", tf.Output(o).DataType())
	str += fmt.Sprint(" shape=", tf.Output(o).Shape())
	return str + ">"
}
