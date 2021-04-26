package tensorflow

import (
	"fmt"
	"sync/atomic"

	// Modules
	op "github.com/djthorpe/go-tensorflow/tensorflow/github.com/tensorflow/tensorflow/tensorflow/go/op"
)

///////////////////////////////////////////////////////////////////////////////
// TYPES

type Scope op.Scope

///////////////////////////////////////////////////////////////////////////////
// GLOBALS

var (
	ctr uint64
)

///////////////////////////////////////////////////////////////////////////////
// PUBLIC METHODS

// NewRoot creates a new *op.Scope, empty
func NewRoot() *Scope {
	return (*Scope)(op.NewScope())
}

// NewScope returns a new scope
func NewScope(root *Scope) *Scope {
	var num = atomic.AddUint64(&ctr, 1)
	return (*Scope)(root.subScope(fmt.Sprint("input_", num)))
}

///////////////////////////////////////////////////////////////////////////////
// PRIVATE METHODS

func (this *Scope) subScope(namespace string) *op.Scope {
	return (*op.Scope)(this).SubScope(namespace)
}
