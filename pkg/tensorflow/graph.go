package tensorflow

import (
	// Modules
	tf "github.com/djthorpe/go-tensorflow/tensorflow/github.com/tensorflow/tensorflow/tensorflow/go"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Graph tf.Graph

///////////////////////////////////////////////////////////////////////////////
// METHODS

// Executes graph with input variables
func (this *Graph) Run(in map[Output]*Tensor, fetches []Output, targets []*Op) ([]*Tensor, error) {
	// Create session
	session, err := tf.NewSession((*tf.Graph)(this), &tf.SessionOptions{})
	if err != nil {
		return nil, err
	}
	defer session.Close()

	// Run session
	in_ := make(map[tf.Output]*tf.Tensor, len(in))
	for k, v := range in {
		in_[tf.Output(k)] = (*tf.Tensor)(v)
	}
	targets_ := make([]*tf.Operation, len(targets))
	for i, op := range targets {
		targets_[i] = (*tf.Operation)(op)
	}
	fetches_ := make([]tf.Output, len(fetches))
	results_ := make([]*Tensor, len(fetches))
	for i, output := range fetches {
		fetches_[i] = tf.Output(output)
	}
	results, err := session.Run(in_, fetches_, targets_)
	if err != nil {
		return nil, err
	}
	for i, tensor := range results {
		results_[i] = (*Tensor)(tensor)
	}
	return results_, nil
}
