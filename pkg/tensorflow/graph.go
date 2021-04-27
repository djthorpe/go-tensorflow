package tensorflow

import (
	// Modules
	tf "github.com/djthorpe/go-tensorflow/tensorflow/github.com/tensorflow/tensorflow/tensorflow/go"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Graph struct {
	*tf.Graph
}

type Session struct {
	*tf.Session
}

///////////////////////////////////////////////////////////////////////////////
// LIFECYCLE

func (graph *Graph) NewSession() (*Session, error) {
	// Create session
	if session, err := tf.NewSession(graph.Graph, &tf.SessionOptions{}); err != nil {
		return nil, err
	} else {
		return &Session{session}, nil
	}
}

func (this *Session) Close() error {
	return this.Session.Close()
}

///////////////////////////////////////////////////////////////////////////////
// METHODS

func (this *Session) Run(in map[Output]*Tensor, fetches []Output, targets []*Op) ([]*Tensor, error) {
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
	results, err := this.Session.Run(in_, fetches_, targets_)
	if err != nil {
		return nil, err
	}
	for i, tensor := range results {
		results_[i] = (*Tensor)(tensor)
	}
	return results_, nil
}

func (this *Graph) Run(in map[Output]*Tensor, fetches []Output, targets []*Op) ([]*Tensor, error) {
	// Create session
	if session, err := this.NewSession(); err != nil {
		return nil, err
	} else {
		defer session.Close()
		return session.Run(in, fetches, targets)
	}
}

func (this *Graph) Ops() []*Op {
	ops := this.Graph.Operations()
	ops_ := make([]*Op, len(ops))
	for i, op := range ops {
		op2 := op // Make a copy of operation before turning into pointer
		ops_[i] = (*Op)(&op2)
	}
	return ops_
}
