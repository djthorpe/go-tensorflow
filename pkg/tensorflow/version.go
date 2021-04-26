package tensorflow

import (
	// Modules
	tf "github.com/djthorpe/go-tensorflow/tensorflow/github.com/tensorflow/tensorflow/tensorflow/go"
)

func Version() string {
	return tf.Version()
}
