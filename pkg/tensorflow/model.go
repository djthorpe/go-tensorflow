package tensorflow

import (
	// Modules
	tf "github.com/djthorpe/go-tensorflow/tensorflow/github.com/tensorflow/tensorflow/tensorflow/go"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Model struct {
	*tf.SavedModel
}

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

func NewModelFromFile(path string, tags []string) (*Model, error) {
	this := new(Model)

	if model, err := tf.LoadSavedModel(path, tags, &tf.SessionOptions{}); err != nil {
		return nil, err
	} else {
		this.SavedModel = model
	}

	// Return success
	return this, nil
}
