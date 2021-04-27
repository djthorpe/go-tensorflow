package tensorflow

import (
	"fmt"

	tf "github.com/djthorpe/go-tensorflow/tensorflow/github.com/tensorflow/tensorflow/tensorflow/go"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Output tf.Output
type Consumer tf.Consumer

///////////////////////////////////////////////////////////////////////////////
// PROPERTIES

func (o Output) Dimensions() int {
	return tf.Output(o).Shape().NumDimensions()
}

func (o Output) DataType() DataType {
	return DataType(tf.Output(o).DataType())
}

func (o Output) Consumers() []Consumer {
	consumers := tf.Output(o).Consumers()
	consumers_ := make([]Consumer, len(consumers))
	for i, consumer := range consumers {
		consumers_[i] = Consumer(consumer)
	}
	return consumers_
}

func (c Consumer) Operation() *Op {
	return (*Op)(c.Op)
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

func (c Consumer) String() string {
	str := "<consumer"
	str += fmt.Sprint(" op=", c.Operation())
	str += fmt.Sprint(" index=", c.Index)
	return str + ">"
}
